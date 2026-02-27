package main

import (
	"math"
	"testing"
)

func TestParseBatteryPower(t *testing.T) {
	tests := []struct {
		name  string
		raw   string
		want  float64
		valid bool
	}{
		{
			name:  "signed discharging value",
			raw:   "-4483",
			want:  4.483,
			valid: true,
		},
		{
			name:  "uint64 wrapped discharging value",
			raw:   "18446744073709547133",
			want:  4.483,
			valid: true,
		},
		{
			name:  "signed charging value",
			raw:   "5326",
			want:  -5.326,
			valid: true,
		},
		{
			name:  "zero value",
			raw:   "0",
			want:  0,
			valid: true,
		},
		{
			name:  "out of range value",
			raw:   "18446744073709544",
			valid: false,
		},
		{
			name:  "non numeric",
			raw:   "bad-value",
			valid: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, ok := parseBatteryPower(tt.raw)
			if ok != tt.valid {
				t.Fatalf("parseBatteryPower(%q) valid = %v, want %v", tt.raw, ok, tt.valid)
			}
			if !tt.valid {
				return
			}
			if math.Abs(got-tt.want) > 0.0001 {
				t.Fatalf("parseBatteryPower(%q) = %v, want %v", tt.raw, got, tt.want)
			}
		})
	}
}
