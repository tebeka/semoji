package main

import (
	"os/exec"
	"strings"
	"testing"
)

func TestEyes(t *testing.T) {
	cmd := exec.Command("go", "run", ".", "eyes")
	data, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("run: %s\n%s", err, string(data))
	}

	out := string(data)
	matches := []string{
		"👀",
		"😁",
		"😙",
		"😚",
		"🙄",
	}
	for _, m := range matches {
		if !strings.Contains(out, m) {
			t.Fatalf("%q missing:\n%s", m, out)
		}
	}
}
