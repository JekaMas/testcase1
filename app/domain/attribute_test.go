package domain

import "testing"

func TestAttributeValid(t *testing.T) {
	attrs := []Attribute{"A0", "A9", "A10", "F1", "Z21"}
	for i := 0; i < len(attrs); i++ {
		if attrs[i].IsCorrectFormat() != true {
			t.Fatalf("Attribute %q should be valid", attrs[i])
		}
	}
}

func TestAttributeNotValid(t *testing.T) {
	attrs := []Attribute{"А", "A-1", "A00", "AA0", "1A", "_A0", " A0", "", " "}
	for i := 0; i < len(attrs); i++ {
		if attrs[i].IsCorrectFormat() != false {
			t.Fatalf("Attribute %q shouldnt be valid", attrs[i])
		}
	}
}
