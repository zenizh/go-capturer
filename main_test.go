package capturer

import (
	"fmt"
	"os"
	"testing"
)

func TestCaptureStdout(t *testing.T) {
	out := CaptureStdout(func() {
		fmt.Fprint(os.Stdout, "foo")
	})

	if out != "foo" {
		t.Errorf("Unexpected output: %s", out)
	}
}

func TestCaptureStderr(t *testing.T) {
	out := CaptureStderr(func() {
		fmt.Fprint(os.Stderr, "foo")
	})

	if out != "foo" {
		t.Errorf("Unexpected output: %s", out)
	}
}

func TestCaptureOutput(t *testing.T) {
	out := CaptureOutput(func() {
		fmt.Fprint(os.Stdout, "foo")
		fmt.Fprint(os.Stderr, "bar")
	})

	if out != "foobar" {
		t.Errorf("Unexpected output: %s", out)
	}
}
