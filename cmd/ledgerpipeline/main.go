package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/joshuataye/ledgerpipeline/internal/account"
	"github.com/joshuataye/ledgerpipeline/internal/batchfile"
	"github.com/joshuataye/ledgerpipeline/internal/budget"
	"github.com/joshuataye/ledgerpipeline/internal/categorize/rules"
	"github.com/joshuataye/ledgerpipeline/internal/export"
	"github.com/joshuataye/ledgerpipeline/internal/import/ofx"
	"github.com/joshuataye/ledgerpipeline/internal/import/qif"
	"github.com/joshuataye/ledgerpipeline/internal/orchestration"
	"github.com/joshuataye/ledgerpipeline/internal/parser"
	"github.com/joshuataye/ledgerpipeline/internal/pipeline"
	"github.com/joshuataye/ledgerpipeline/internal/report"
	"github.com/joshuataye/ledgerpipeline/internal/validate"
)

func main() {
	if len(os.Args) < 2 {
		printUsage()
		os.Exit(1)
	}
	switch os.Args[1] {
	case "summarize":
		os.Exit(runSummarize(os.Args[2:]))
	case "reconcile":
		os.Exit(runReconcile(os.Args[2:]))
	case "budget":
		os.Exit(runBudget(os.Args[2:]))
	case "recurring":
		os.Exit(runRecurring(os.Args[2:]))
	case "export":
		os.Exit(runExport(os.Args[2:]))
	case "batch":
		os.Exit(runBatch(os.Args[2:]))
	case "import-ofx":
		os.Exit(runImportOFX(os.Args[2:]))
	case "import-qif":
		os.Exit(runImportQIF(os.Args[2:]))
	case "help", "-h", "--help":
		printUsage()
	default:
		fmt.Fprintf(os.Stderr, "unknown command %q\n", os.Args[1])
		printUsage()
		os.Exit(1)
	}
}

func printUsage() {
	fmt.Fprintf(os.Stderr, `usage: ledgerpipeline <command> [flags]

commands:
  summarize   aggregate categories from a CSV statement
  reconcile   compare opening/closing balances to parsed rows
  budget      compare actual category totals against budget lines
  recurring   detect recurring charges in a statement
  export      write JSON category summary
  batch       summarize every CSV in a directory
  import-ofx  parse an OFX snippet file to CSV-style rows on stdout
  import-qif  parse a QIF file to CSV-style rows on stdout

`)
}

func parseInputFlag(args []string) (string, []string, error) {
	fs := flag.NewFlagSet("cmd", flag.ExitOnError)
	input := fs.String("input", "", "path to CSV bank statement")
	if err := fs.Parse(args); err != nil {
		return "", nil, err
	}
	if *input == "" {
		return "", nil, fmt.Errorf("-input required")
	}
	return *input, fs.Args(), nil
}

func loadTransactions(path string) ([]parser.Transaction, error) {
	return parser.ParseFile(path)
}

func runSummarize(args []string) int {
	fs := flag.NewFlagSet("summarize", flag.ExitOnError)
	input := fs.String("input", "", "path to CSV bank statement")
	output := fs.String("output", "", "optional output file")
	dedupe := fs.Bool("dedupe", false, "remove duplicate rows")
	normalize := fs.Bool("normalize", false, "normalize merchant descriptions")
	rulesPath := fs.String("rules", "", "JSON file of categorization rules")
	accountID := fs.String("account", "default", "account id for orchestrated run")
	opening := fs.Float64("opening", 0, "account opening balance")
	if err := fs.Parse(args); err != nil {
		return 1
	}
	if *input == "" {
		fmt.Fprintln(os.Stderr, "usage: ledgerpipeline summarize -input <csv> [-output file] [-dedupe] [-normalize]")
		return 1
	}
	txns, err := loadTransactions(*input)
	if err != nil {
		fmt.Fprintf(os.Stderr, "parse: %v\n", err)
		return 1
	}
	var catRules []rules.Rule
	if *rulesPath != "" {
		catRules, err = rules.LoadFile(*rulesPath)
		if err != nil {
			fmt.Fprintf(os.Stderr, "rules: %v\n", err)
			return 1
		}
	}
	reg, err := account.NewRegistry(account.Account{
		ID: *accountID, Name: *accountID, Type: account.TypeChecking, Opening: *opening,
	})
	if err != nil {
		fmt.Fprintf(os.Stderr, "account: %v\n", err)
		return 1
	}
	orch := orchestration.NewBatchOrchestrator(reg)
	result, err := orch.Run(txns, orchestration.RunConfig{
		AccountID: *accountID,
		Pipeline: pipeline.Config{
			Dedupe:          *dedupe,
			Normalize:       *normalize,
			CategorizeRules: catRules,
			Validate:        validate.Validator{},
		},
	})
	if err != nil {
		fmt.Fprintf(os.Stderr, "pipeline: %v\n", err)
		return 1
	}
	out := os.Stdout
	if *output != "" {
		f, err := os.Create(*output)
		if err != nil {
			fmt.Fprintf(os.Stderr, "create output: %v\n", err)
			return 1
		}
		defer f.Close()
		out = f
	}
	if err := report.WriteSummary(out, result.Pipeline.Summaries, result.Pipeline.NetTotal); err != nil {
		fmt.Fprintf(os.Stderr, "report: %v\n", err)
		return 1
	}
	return 0
}

func runReconcile(args []string) int {
	fs := flag.NewFlagSet("reconcile", flag.ExitOnError)
	input := fs.String("input", "", "path to CSV bank statement")
	opening := fs.Float64("opening", 0, "opening balance")
	closing := fs.Float64("closing", 0, "closing balance")
	dedupe := fs.Bool("dedupe", false, "remove duplicate rows")
	if err := fs.Parse(args); err != nil {
		return 1
	}
	if *input == "" {
		fmt.Fprintln(os.Stderr, "usage: ledgerpipeline reconcile -input <csv> -opening N -closing N")
		return 1
	}
	txns, err := loadTransactions(*input)
	if err != nil {
		fmt.Fprintf(os.Stderr, "parse: %v\n", err)
		return 1
	}
	result, err := pipeline.Run(txns, pipeline.Config{
		Dedupe:   *dedupe,
		Validate: validate.Validator{},
		Reconcile: &pipeline.ReconcileConfig{
			Opening: *opening,
			Closing: *closing,
		},
	})
	if err != nil {
		fmt.Fprintf(os.Stderr, "pipeline: %v\n", err)
		return 1
	}
	r := result.Reconcile
	fmt.Printf("opening: %.2f\nclosing: %.2f\ncomputed: %.2f\ndelta: %.2f\n",
		r.OpeningBalance, r.ClosingBalance, r.ComputedClose, r.Delta)
	return 0
}

func runBudget(args []string) int {
	fs := flag.NewFlagSet("budget", flag.ExitOnError)
	input := fs.String("input", "", "path to CSV bank statement")
	limits := fs.String("limits", "", "comma-separated category:limit pairs")
	if err := fs.Parse(args); err != nil {
		return 1
	}
	if *input == "" || *limits == "" {
		fmt.Fprintln(os.Stderr, "usage: ledgerpipeline budget -input <csv> -limits Food:200,Transport:50")
		return 1
	}
	txns, err := loadTransactions(*input)
	if err != nil {
		fmt.Fprintf(os.Stderr, "parse: %v\n", err)
		return 1
	}
	result, err := pipeline.Run(txns, pipeline.Config{Validate: validate.Validator{}})
	if err != nil {
		fmt.Fprintf(os.Stderr, "pipeline: %v\n", err)
		return 1
	}
	actual := map[string]float64{}
	for _, s := range result.Summaries {
		actual[s.Category] = s.TotalAmount
	}
	lines, err := parseBudgetLines(*limits)
	if err != nil {
		fmt.Fprintf(os.Stderr, "limits: %v\n", err)
		return 1
	}
	for _, v := range budget.Compare(actual, lines) {
		fmt.Println(budget.FormatVariance(v))
	}
	return 0
}

func parseBudgetLines(spec string) ([]budget.Line, error) {
	var lines []budget.Line
	for _, part := range strings.Split(spec, ",") {
		part = strings.TrimSpace(part)
		if part == "" {
			continue
		}
		kv := strings.SplitN(part, ":", 2)
		if len(kv) != 2 {
			return nil, fmt.Errorf("invalid limit %q", part)
		}
		limit, err := strconv.ParseFloat(strings.TrimSpace(kv[1]), 64)
		if err != nil {
			return nil, err
		}
		lines = append(lines, budget.Line{Category: strings.TrimSpace(kv[0]), Limit: limit})
	}
	return lines, nil
}

func runRecurring(args []string) int {
	fs := flag.NewFlagSet("recurring", flag.ExitOnError)
	input := fs.String("input", "", "path to CSV bank statement")
	minCount := fs.Int("min-count", 2, "minimum occurrences")
	if err := fs.Parse(args); err != nil {
		return 1
	}
	if *input == "" {
		fmt.Fprintln(os.Stderr, "usage: ledgerpipeline recurring -input <csv> [-min-count N]")
		return 1
	}
	txns, err := loadTransactions(*input)
	if err != nil {
		fmt.Fprintf(os.Stderr, "parse: %v\n", err)
		return 1
	}
	result, err := pipeline.Run(txns, pipeline.Config{
		Validate:     validate.Validator{},
		MinRecurring: *minCount,
	})
	if err != nil {
		fmt.Fprintf(os.Stderr, "pipeline: %v\n", err)
		return 1
	}
	for _, c := range result.Recurring {
		fmt.Printf("%s (%s): count=%d avg=%.2f\n", c.Description, c.Category, c.Count, c.AvgAmount)
	}
	return 0
}

func runExport(args []string) int {
	fs := flag.NewFlagSet("export", flag.ExitOnError)
	input := fs.String("input", "", "path to CSV bank statement")
	output := fs.String("output", "", "JSON output path")
	dedupe := fs.Bool("dedupe", false, "remove duplicate rows")
	if err := fs.Parse(args); err != nil {
		return 1
	}
	if *input == "" {
		fmt.Fprintln(os.Stderr, "usage: ledgerpipeline export -input <csv> [-output file.json]")
		return 1
	}
	txns, err := loadTransactions(*input)
	if err != nil {
		fmt.Fprintf(os.Stderr, "parse: %v\n", err)
		return 1
	}
	result, err := pipeline.Run(txns, pipeline.Config{Dedupe: *dedupe, Validate: validate.Validator{}})
	if err != nil {
		fmt.Fprintf(os.Stderr, "pipeline: %v\n", err)
		return 1
	}
	out := os.Stdout
	if *output != "" {
		f, err := os.Create(*output)
		if err != nil {
			fmt.Fprintf(os.Stderr, "create output: %v\n", err)
			return 1
		}
		defer f.Close()
		out = f
	}
	if err := export.WriteJSON(out, result.Summaries, result.NetTotal); err != nil {
		fmt.Fprintf(os.Stderr, "export: %v\n", err)
		return 1
	}
	return 0
}

func runBatch(args []string) int {
	fs := flag.NewFlagSet("batch", flag.ExitOnError)
	dir := fs.String("dir", "", "directory containing CSV statements")
	dedupe := fs.Bool("dedupe", false, "remove duplicate rows")
	if err := fs.Parse(args); err != nil {
		return 1
	}
	if *dir == "" {
		fmt.Fprintln(os.Stderr, "usage: ledgerpipeline batch -dir <path> [-dedupe]")
		return 1
	}
	txns, err := batchfile.ReadDir(*dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "batch: %v\n", err)
		return 1
	}
	result, err := pipeline.Run(txns, pipeline.Config{Dedupe: *dedupe, Validate: validate.Validator{}})
	if err != nil {
		fmt.Fprintf(os.Stderr, "pipeline: %v\n", err)
		return 1
	}
	if err := report.WriteSummary(os.Stdout, result.Summaries, result.NetTotal); err != nil {
		fmt.Fprintf(os.Stderr, "report: %v\n", err)
		return 1
	}
	return 0
}

func runImportOFX(args []string) int {
	fs := flag.NewFlagSet("import-ofx", flag.ExitOnError)
	input := fs.String("input", "", "path to OFX snippet file")
	if err := fs.Parse(args); err != nil {
		return 1
	}
	if *input == "" {
		fmt.Fprintln(os.Stderr, "usage: ledgerpipeline import-ofx -input <file.ofx>")
		return 1
	}
	txns, err := ofx.ParseFile(*input)
	if err != nil {
		fmt.Fprintf(os.Stderr, "ofx: %v\n", err)
		return 1
	}
	printTransactions(txns)
	return 0
}

func runImportQIF(args []string) int {
	fs := flag.NewFlagSet("import-qif", flag.ExitOnError)
	input := fs.String("input", "", "path to QIF file")
	if err := fs.Parse(args); err != nil {
		return 1
	}
	if *input == "" {
		fmt.Fprintln(os.Stderr, "usage: ledgerpipeline import-qif -input <file.qif>")
		return 1
	}
	txns, err := qif.ParseFile(*input)
	if err != nil {
		fmt.Fprintf(os.Stderr, "qif: %v\n", err)
		return 1
	}
	printTransactions(txns)
	return 0
}

func printTransactions(txns []parser.Transaction) {
	fmt.Println("date,description,category,amount")
	for _, tx := range txns {
		fmt.Printf("%s,%s,%s,%.2f\n",
			tx.Date.Format("2006-01-02"), tx.Description, tx.Category, tx.Amount)
	}
}
