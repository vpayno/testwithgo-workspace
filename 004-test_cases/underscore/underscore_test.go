package underscore

import (
	"fmt"
	"testing"
)

func TestCamel_orig(t *testing.T) {
	// moved test cases to underscore_cases_test.go
	for _, tc := range camelTestCases {
		t.Logf("Testing: %q", tc.arg)
		got := Camel(tc.arg)
		if got != tc.want {
			t.Errorf("Camel(%q) = %q; want %q", tc.arg, got, tc.want)
		}
	}
}

func TestCamel(t *testing.T) {
	// moved test cases to underscore_cases_test.go
	for _, tt := range camelTestCases {
		t.Run(tt.arg, func(t *testing.T) {
			if got := Camel(tt.arg); got != tt.want {
				t.Fatalf("Camel() = %v, want %v", got, tt.want)
			}
			t.Log("this will print if it succeeds or fails...")
			fmt.Println("this won't print if it fails...")
		})
	}
}
