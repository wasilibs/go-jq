package jq

import (
	"bytes"
	"strings"
	"testing"

	"github.com/wasilibs/go-jq/internal/runner"
	"github.com/wasilibs/go-jq/internal/wasm"
)

func TestRuns(t *testing.T) {
	stdin := strings.NewReader(`{"foo": 0}`)
	stdout := bytes.Buffer{}
	stderr := bytes.Buffer{}

	ret := runner.Run("jq", []string{}, wasm.Jq, stdin, &stdout, &stderr, ".")
	if ret != 0 {
		t.Fatalf("unexpected return code: %d", ret)
	}

	if have, want := stdout.String(), "{\n  \"foo\": 0\n}\n"; have != want {
		t.Fatalf("unexpected stdout: have=%q want=%q", have, want)
	}
}
