package account

import "testing"

func TestRegistryRejectsDuplicateID(t *testing.T) {
	_, err := NewRegistry(
		Account{ID: "a1", Name: "Main", Type: TypeChecking},
		Account{ID: "a1", Name: "Dup", Type: TypeSavings},
	)
	if err == nil {
		t.Fatal("expected duplicate error")
	}
}

func TestRegistryGet(t *testing.T) {
	r, err := NewRegistry(Account{ID: "chk", Name: "Checking", Type: TypeChecking, Opening: 100})
	if err != nil {
		t.Fatal(err)
	}
	a, ok := r.Get("chk")
	if !ok || a.Opening != 100 {
		t.Fatalf("got %+v ok=%v", a, ok)
	}
}
