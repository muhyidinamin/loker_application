package main

import (
	"testing"
)

type CommandResult struct {
	command string
}

var commandResults = []CommandResult{
	{"init 5"},
	{"input SIM 12345"},
	{"input KTP 34710"},
	{"input KTP 35770"},
	{"input KTP 24710"},
	{"input SIM 98775"},
	{"status"},
	{"input SIM 87214"},
	{"leave 3"},
	{"input SIM 87214"},
	{"find 34710"},
	{"search SIM"},
	{"find 99999"},
}

func TestResult(t *testing.T) {
	for _, test := range commandResults {
		result := runCommand(test.command)
		if result != nil {
			t.Error(result.Error())
		}
	}
}
