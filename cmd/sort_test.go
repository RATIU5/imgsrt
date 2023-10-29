package cmd

import (
	"testing"
	"time"
)

func TestParsePattern(t *testing.T) {
	tests := []struct {
		pattern string
		time    time.Time
		want    string
	}{
		{"y/m", time.Date(2022, 10, 1, 0, 0, 0, 0, time.UTC), "2022/10"},
		{"y/m/d", time.Date(2022, 10, 1, 0, 0, 0, 0, time.UTC), "2022/10/01"},
		{"y", time.Date(2022, 10, 1, 0, 0, 0, 0, time.UTC), "2022"},
		{"m", time.Date(2022, 10, 1, 0, 0, 0, 0, time.UTC), "10"},
		{"m/y", time.Date(2022, 10, 1, 0, 0, 0, 0, time.UTC), "10/2022"},
	}

	for _, test := range tests {
		dirGen := parsePattern(test.pattern)
		got := dirGen(test.time)
		if got != test.want {
			t.Errorf("parsePattern(%q)(%v) = %q, want %q", test.pattern, test.time, got, test.want)
		}
	}
}
