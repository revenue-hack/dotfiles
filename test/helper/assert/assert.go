package assert

import (
	"bytes"
	"encoding/json"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Equal(t *testing.T, actual, expected any) {
	t.Helper()
	if diff := cmp.Diff(actual, expected); diff != "" {
		t.Errorf("values must match:\n%s", diff)
	}
}

func NotEqual(t *testing.T, actual, expected any) {
	t.Helper()
	if ok := cmp.Equal(actual, expected); ok {
		t.Errorf("values must not match: %v", expected)
	}
}

func EqualJson(t *testing.T, actual, expected string) {
	t.Helper()
	if diff := cmp.Diff(compactJson(actual), compactJson(expected)); diff != "" {
		t.Errorf("json values must match:\n%s", diff)
	}
}

func compactJson(s string) string {
	buf := bytes.NewBuffer(nil)
	if err := json.Compact(buf, []byte(s)); err != nil {
		return s
	} else {
		return buf.String()
	}
}
