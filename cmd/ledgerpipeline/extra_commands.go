package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/joshuataye/ledgerpipeline/internal/anomaly"
	"github.com/joshuataye/ledgerpipeline/internal/budget"
	"github.com/joshuataye/ledgerpipeline/internal/cliutil"
	"github.com/joshuataye/ledgerpipeline/internal/compare"
	"github.com/joshuataye/ledgerpipeline/internal/config"
	"github.com/joshuataye/ledgerpipeline/internal/filter"
	"github.com/joshuataye/ledgerpipeline/internal/forecast"
	"github.com/joshuataye/ledgerpipeline/internal/format"
	"github.com/joshuataye/ledgerpipeline/internal/import/fixedwidth"
	"github.com/joshuataye/ledgerpipeline/internal/importfmt"
	"github.com/joshuataye/ledgerpipeline/internal/insights"
	"github.com/joshuataye/ledgerpipeline/internal/matching"
	"github.com/joshuataye/ledgerpipeline/internal/period"
	"github.com/joshuataye/ledgerpipeline/internal/pipeline"
	"github.com/joshuataye/ledgerpipeline/internal/recurring"
	"github.com/joshuataye/ledgerpipeline/internal/stats"
	"github.com/joshuataye/ledgerpipeline/internal/storage"
	"github.com/joshuataye/ledgerpipeline/internal/tax"
	"github.com/joshuataye/ledgerpipeline/internal/validate"
)

func runInsights(args []string) int {
	input, rest, err := parseInputFlag(args)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 1
	}
	fs := flag.NewFlagSet("insights", flag.ContinueOnError)
	filterFlags := cliutil.ParseFilterFlags(fs)
	_ = fs.Parse(rest)
	txns, err := loadTransactions(input)
	if err != nil {
		fmt.Fprintf(os.Stderr, "parse: %v\n", err)
		return 1
	}
	result, err := pipeline.Run(txns, pipeline.Config{
		Filter:   filterFlags.Options(),
		Validate: validate.Validator{},
	})
	if err != nil {
		fmt.Fprintf(os.Stderr, "pipeline: %v\n", err)
		return 1
	}
	summary := insights.Build(result.Transactions, result.Summaries, result.NetTotal)
	return printErr(insights.WriteReport(os.Stdout, summary))
}

func runMonthly(args []string) int {
	input, rest, err := parseInputFlag(args)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 1
	}
	fs := flag.NewFlagSet("monthly", flag.ContinueOnError)
	filterFlags := cliutil.ParseFilterFlags(fs)
	_ = fs.Parse(rest)
	txns, err := loadTransactions(input)
	if err != nil {
		fmt.Fprintf(os.Stderr, "parse: %v\n", err)
		return 1
	}
	txns = filter.Apply(txns, filterFlags.Options())
	for _, r := range period.ByMonth(txns) {
		fmt.Printf("%s count=%d net=%.2f\n", r.Label(), r.Count, r.Net)
	}
	return 0
}

func runStats(args []string) int {
	input, rest, err := parseInputFlag(args)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 1
	}
	fs := flag.NewFlagSet("stats", flag.ContinueOnError)
	filterFlags := cliutil.ParseFilterFlags(fs)
	_ = fs.Parse(rest)
	txns, err := loadTransactions(input)
	if err != nil {
		fmt.Fprintf(os.Stderr, "parse: %v\n", err)
		return 1
	}
	txns = filter.Apply(txns, filterFlags.Options())
	s := stats.Compute(txns)
	fmt.Printf("count=%d debit_sum=%.2f credit_sum=%.2f largest_debit=%.2f\n",
		s.Count, s.DebitSum, s.CreditSum, s.LargestDebit)
	return 0
}

func runMatchTransfers(args []string) int {
	input, rest, err := parseInputFlag(args)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 1
	}
	fs := flag.NewFlagSet("match-transfers", flag.ContinueOnError)
	maxDays := fs.Int("max-days", 3, "maximum days between transfer legs")
	_ = fs.Parse(rest)
	txns, err := loadTransactions(input)
	if err != nil {
		fmt.Fprintf(os.Stderr, "parse: %v\n", err)
		return 1
	}
	opt := matching.DefaultOptions()
	opt.MaxDays = *maxDays
	pairs := matching.FindTransfers(txns, opt)
	for _, p := range pairs {
		fmt.Printf("%s %.2f -> %s %.2f\n",
			p.From.Date.Format("2006-01-02"), p.From.Amount,
			p.To.Date.Format("2006-01-02"), p.To.Amount)
	}
	unmatched := matching.Unmatched(txns, pairs)
	fmt.Printf("unmatched: %d\n", len(unmatched))
	return 0
}

func runAnomalies(args []string) int {
	input, rest, err := parseInputFlag(args)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 1
	}
	fs := flag.NewFlagSet("anomalies", flag.ContinueOnError)
	threshold := fs.Float64("threshold", 2.0, "z-score threshold")
	_ = fs.Parse(rest)
	txns, err := loadTransactions(input)
	if err != nil {
		fmt.Fprintf(os.Stderr, "parse: %v\n", err)
		return 1
	}
	for _, hit := range anomaly.Detect(txns, *threshold) {
		fmt.Printf("%s %.2f z=%.2f\n",
			hit.Transaction.Date.Format("2006-01-02"),
			hit.Transaction.Amount, hit.ZScore)
	}
	return 0
}

func runForecast(args []string) int {
	input, rest, err := parseInputFlag(args)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 1
	}
	fs := flag.NewFlagSet("forecast", flag.ContinueOnError)
	months := fs.Int("months", 3, "months of history to use")
	_ = fs.Parse(rest)
	txns, err := loadTransactions(input)
	if err != nil {
		fmt.Fprintf(os.Stderr, "parse: %v\n", err)
		return 1
	}
	for _, p := range forecast.Linear(txns, *months) {
		fmt.Printf("%s net=%.2f\n", p.Label, p.Net)
	}
	return 0
}

func runCompare(args []string) int {
	fs := flag.NewFlagSet("compare", flag.ContinueOnError)
	beforePath := fs.String("before", "", "path to earlier CSV statement")
	afterPath := fs.String("after", "", "path to later CSV statement")
	if err := fs.Parse(args); err != nil {
		return 1
	}
	if *beforePath == "" || *afterPath == "" {
		fmt.Fprintln(os.Stderr, "usage: ledgerpipeline compare -before <csv> -after <csv>")
		return 1
	}
	beforeTxns, err := loadTransactions(*beforePath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "before: %v\n", err)
		return 1
	}
	afterTxns, err := loadTransactions(*afterPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "after: %v\n", err)
		return 1
	}
	deltas := compare.Periods(period.ByMonth(beforeTxns), period.ByMonth(afterTxns))
	for _, d := range deltas {
		fmt.Printf("%s prev=%.2f cur=%.2f change=%.2f\n", d.Label, d.Previous, d.Current, d.Change)
	}
	return 0
}

func runTaxReport(args []string) int {
	input, rest, err := parseInputFlag(args)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 1
	}
	fs := flag.NewFlagSet("tax-report", flag.ContinueOnError)
	_ = fs.Parse(rest)
	txns, err := loadTransactions(input)
	if err != nil {
		fmt.Fprintf(os.Stderr, "parse: %v\n", err)
		return 1
	}
	return printErr(tax.WriteReport(os.Stdout, tax.Report(txns)))
}

func runSnapshot(args []string) int {
	fs := flag.NewFlagSet("snapshot", flag.ContinueOnError)
	input := fs.String("input", "", "path to CSV bank statement")
	output := fs.String("output", "", "snapshot JSON output path")
	loadPath := fs.String("load", "", "load and print snapshot JSON instead of saving")
	if err := fs.Parse(args); err != nil {
		return 1
	}
	if *loadPath != "" {
		txns, err := storage.Load(*loadPath)
		if err != nil {
			fmt.Fprintf(os.Stderr, "load: %v\n", err)
			return 1
		}
		printTransactions(txns)
		return 0
	}
	if *input == "" || *output == "" {
		fmt.Fprintln(os.Stderr, "usage: ledgerpipeline snapshot -input <csv> -output <file.json>")
		return 1
	}
	txns, err := loadTransactions(*input)
	if err != nil {
		fmt.Fprintf(os.Stderr, "parse: %v\n", err)
		return 1
	}
	if err := storage.Save(*output, txns); err != nil {
		fmt.Fprintf(os.Stderr, "save: %v\n", err)
		return 1
	}
	return 0
}

func runImportAny(args []string) int {
	fs := flag.NewFlagSet("import-any", flag.ContinueOnError)
	input := fs.String("input", "", "path to statement file")
	if err := fs.Parse(args); err != nil {
		return 1
	}
	if *input == "" {
		fmt.Fprintln(os.Stderr, "usage: ledgerpipeline import-any -input <file>")
		return 1
	}
	format, err := importfmt.Detect(*input)
	if err != nil {
		fmt.Fprintf(os.Stderr, "detect: %v\n", err)
		return 1
	}
	fmt.Fprintf(os.Stderr, "format: %s\n", format)
	txns, err := importfmt.ParseFile(*input)
	if err != nil {
		fmt.Fprintf(os.Stderr, "parse: %v\n", err)
		return 1
	}
	printTransactions(txns)
	return 0
}

func runImportFixedWidth(args []string) int {
	fs := flag.NewFlagSet("import-fixedwidth", flag.ContinueOnError)
	input := fs.String("input", "", "path to fixed-width export file")
	if err := fs.Parse(args); err != nil {
		return 1
	}
	if *input == "" {
		fmt.Fprintln(os.Stderr, "usage: ledgerpipeline import-fixedwidth -input <file.txt>")
		return 1
	}
	data, err := os.ReadFile(*input)
	if err != nil {
		fmt.Fprintf(os.Stderr, "read: %v\n", err)
		return 1
	}
	lines := strings.Split(string(data), "\n")
	txns, err := fixedwidth.ParseLines(lines, importfmt.DefaultFixedWidthLayout())
	if err != nil {
		fmt.Fprintf(os.Stderr, "parse: %v\n", err)
		return 1
	}
	printTransactions(txns)
	return 0
}

func runIntervals(args []string) int {
	input, rest, err := parseInputFlag(args)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 1
	}
	fs := flag.NewFlagSet("intervals", flag.ContinueOnError)
	minCount := fs.Int("min-count", 2, "minimum occurrences")
	_ = fs.Parse(rest)
	txns, err := loadTransactions(input)
	if err != nil {
		fmt.Fprintf(os.Stderr, "parse: %v\n", err)
		return 1
	}
	for _, iv := range recurring.Intervals(txns, *minCount) {
		fmt.Printf("%s (%s): every %d days, count=%d\n",
			iv.Description, iv.Category, iv.Days, iv.Count)
	}
	return 0
}

func runFormatExport(args []string) int {
	fs := flag.NewFlagSet("format-export", flag.ContinueOnError)
	input := fs.String("input", "", "path to CSV bank statement")
	output := fs.String("output", "", "optional output file")
	formatName := fs.String("format", "csv", "output format: csv, tsv, markdown")
	filterFlags := cliutil.ParseFilterFlags(fs)
	if err := fs.Parse(args); err != nil {
		return 1
	}
	if *input == "" {
		fmt.Fprintln(os.Stderr, "usage: ledgerpipeline format-export -input <csv> [-format csv|tsv|markdown]")
		return 1
	}
	txns, err := loadTransactions(*input)
	if err != nil {
		fmt.Fprintf(os.Stderr, "parse: %v\n", err)
		return 1
	}
	result, err := pipeline.Run(txns, pipeline.Config{
		Filter:   filterFlags.Options(),
		Validate: validate.Validator{},
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
	switch strings.ToLower(*formatName) {
	case "tsv":
		return printErr(format.WriteTSV(out, result.Summaries))
	case "markdown", "md":
		return printErr(format.WriteMarkdown(out, result.Summaries, result.NetTotal))
	default:
		return printErr(format.WriteCSV(out, result.Summaries))
	}
}

func runBudgetAnalysis(args []string) int {
	fs := flag.NewFlagSet("budget-analysis", flag.ContinueOnError)
	input := fs.String("input", "", "path to CSV bank statement")
	limits := fs.String("limits", "", "comma-separated category:limit pairs")
	if err := fs.Parse(args); err != nil {
		return 1
	}
	if *input == "" || *limits == "" {
		fmt.Fprintln(os.Stderr, "usage: ledgerpipeline budget-analysis -input <csv> -limits Food:200")
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
	for _, u := range budget.Analyze(budget.Compare(actual, lines)) {
		fmt.Printf("%s: used %.1f%% (actual %.2f limit %.2f)\n",
			u.Category, u.UsedPct, u.Actual, u.Limit)
	}
	return 0
}

func runFilterPreset(args []string) int {
	fs := flag.NewFlagSet("filter-preset", flag.ContinueOnError)
	configPath := fs.String("config", "", "filter JSON config path")
	days := fs.Int("days", 0, "last N days preset")
	input := fs.String("input", "", "path to CSV bank statement")
	if err := fs.Parse(args); err != nil {
		return 1
	}
	if *input == "" {
		fmt.Fprintln(os.Stderr, "usage: ledgerpipeline filter-preset -input <csv> [-config file] [-days N]")
		return 1
	}
	var opt filter.Options
	if *configPath != "" {
		cfg, err := config.Load(*configPath)
		if err != nil {
			fmt.Fprintf(os.Stderr, "config: %v\n", err)
			return 1
		}
		opt = filter.FromConfig(cfg)
	}
	if *days > 0 {
		opt = filter.PresetLastNDays(opt.To, *days)
	}
	txns, err := loadTransactions(*input)
	if err != nil {
		fmt.Fprintf(os.Stderr, "parse: %v\n", err)
		return 1
	}
	filtered := filter.Apply(txns, opt)
	fmt.Printf("kept %d of %d transactions\n", len(filtered), len(txns))
	return 0
}

func printErr(err error) int {
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		return 1
	}
	return 0
}
