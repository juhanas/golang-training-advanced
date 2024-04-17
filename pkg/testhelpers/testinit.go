package testhelpers

import (
	"log/slog"
	"os"
)

// init is run before tests in this package and all packages that import this package
// Each separate init-func is run once (in random order), they do not replace each other
// Now we use this function to configure the logger for tests
// This function will be always called before tests in all packages that import this package
// Note1: init() is special in Go, and thus does not need to be exported specifically (no "Init()")
// Note2: In general, side-effects should be avoided. This is to demonstrate the syntax and possibilities in Go.
func init() {
	logLevel := slog.LevelInfo
	opts := &slog.HandlerOptions{
		AddSource: false,
		Level:     logLevel,
	}
	var handler slog.Handler = slog.NewJSONHandler(os.Stdout, opts)
	logger := slog.New(handler)
	slog.SetDefault(logger)
}
