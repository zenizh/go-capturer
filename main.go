package capturer

import (
	"bytes"
	"io"
	"os"
)

var captureStdout, captureStderr bool

// CaptureStdout captures stdout.
func CaptureStdout(f func()) string {
	captureStdout = true
	defer func() {
		captureStdout = false
	}()
	return capture(f)
}

// CaptureStderr captures stderr.
func CaptureStderr(f func()) string {
	captureStderr = true
	defer func() {
		captureStderr = false
	}()
	return capture(f)
}

// CaptureOutput captures stdout and stderr.
func CaptureOutput(f func()) string {
	captureStdout = true
	captureStderr = true
	defer func() {
		captureStdout = false
		captureStderr = false
	}()
	return capture(f)
}

func capture(f func()) string {
	r, w, err := os.Pipe()
	if err != nil {
		panic(err)
	}

	if captureStdout {
		stdout := os.Stdout
		os.Stdout = w
		defer func() {
			os.Stdout = stdout
		}()
	}

	if captureStderr {
		stderr := os.Stderr
		os.Stderr = w
		defer func() {
			os.Stderr = stderr
		}()
	}

	f()
	w.Close()

	var buf bytes.Buffer
	io.Copy(&buf, r)

	return buf.String()
}
