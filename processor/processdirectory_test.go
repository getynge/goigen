package processor

import "testing"

func TestProcessDirectory(t *testing.T) {
	methods1, _, err := ProcessDirectory("testfiles", "*Example")

	if err != nil {
		t.Fatal(err)
	}

	methods2, _, err := ProcessDirectory("testfiles", "Example")

	if err != nil {
		t.Fatal(err)
	}

	if len(methods1) != 2 {
		t.Fatalf("expected 2 methods, got %d\n", len(methods1))
	}

	if len(methods2) != 1 {
		t.Fatalf("expected 1 method, got %d\n", len(methods2))
	}
}
