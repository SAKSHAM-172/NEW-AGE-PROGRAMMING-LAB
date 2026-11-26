package main

import "testing"

func TestDivide(t *testing.T) {
	result, err := Divide(10, 2)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result != 5 {
		t.Errorf("Divide(10, 2) = %v; want 5", result)
	}

	_, err = Divide(10, 0)
	if err == nil {
		t.Error("expected error for division by zero, got nil")
	}
}

func TestSquareRoot(t *testing.T) {
	result, err := SquareRoot(16)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result != 4 {
		t.Errorf("SquareRoot(16) = %v; want 4", result)
	}

	_, err = SquareRoot(-9)
	if err == nil {
		t.Error("expected error for negative input, got nil")
	}
}

func TestGreet(t *testing.T) {
	msg, err := Greet("Asha")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if msg != "Hello, Asha!" {
		t.Errorf("Greet(\"Asha\") = %q; want %q", msg, "Hello, Asha!")
	}

	_, err = Greet("")
	if err == nil {
		t.Error("expected error for empty name, got nil")
	}
}
