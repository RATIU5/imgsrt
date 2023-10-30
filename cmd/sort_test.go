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
		{"y/M", time.Date(2022, 10, 1, 0, 0, 0, 0, time.UTC), "2022/October"},
		{"M/y", time.Date(2022, 10, 1, 0, 0, 0, 0, time.UTC), "October/2022"},
	}

	for _, test := range tests {
		dirGen := parsePattern(test.pattern)
		got := dirGen(test.time)
		if got != test.want {
			t.Errorf("parsePattern(%q)(%v) = %q, want %q", test.pattern, test.time, got, test.want)
		}
	}
}

func TestIsImageOrVideo(t *testing.T) {
	tests := []struct {
		filename string
		want     bool
	}{
		{"example.jpg", true},
		{"example.jpeg", true},
		{"example.png", true},
		{"example.mp4", true},
		{"example.mov", true},
		{"example.r3d", true},
		{"example.txt", false},
		{"example.doc", false},
		{"example.pdf", false},
	}

	for _, test := range tests {
		t.Run(test.filename, func(t *testing.T) {
			got := isImageOrVideo(test.filename)
			if got != test.want {
				t.Errorf("isImageOrVideo(%q) = %v, want %v", test.filename, got, test.want)
			}
		})
	}
}
