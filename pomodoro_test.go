package main

import (
	"testing"
)

func TestValidateNumber(t *testing.T) {
	tests := []struct {
		name string
		input int
		wantErr bool
	}{
		{"small positive integer", 1, false},
		{"large positive integer", 1000000, false},
		{"zero", 0, true},
		{"small negative integer", -1, true},
		{"large negative integer", -1000000, true},
	}

	for _, tt := range tests {
		err := validateNumber(tt.input)
		gotError := err != nil

	 	if gotError != tt.wantErr {
			t.Errorf("test '%s': expected error=%v, got error=%v", tt.name, tt.wantErr, gotError)
		}
	}
}