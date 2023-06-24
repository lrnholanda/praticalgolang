package main

import (
	"errors"
	"testing"
)

type testConfig struct {
	args []string
	err  error
	config
}

func TestParseArgs(t *testing.T) {
	for _, tc := range tests {
		c, err := parseArgs(tc.args)
		if tc.result.err != nil && err.Error() != tc.result.err.Error() {
			t.Fatalf("Expected error to be: %v, got: %v\n", tc.result.err, err)
		}
		if tc.result.err == nil && err != nil {
			t.Errorf("Expected nil error, got: %v\n", err)
		}
		if c.printUsage != tc.result.printUsage {
			t.Errorf("Expected printUsage to be: %v, got: %v\n", tc.result.printUsage, c.printUsage)
		}
		if c.numTimes != tc.result.numTimes {
			t.Errorf("Expected numTimes to be: %v, got: %v\n", tc.result.numTimes, c.numTimes)
		}
	}
}
