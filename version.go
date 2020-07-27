package main

import (
	"fmt"
	"runtime"
)

var (
	// Version release version
	Version = "0.0.1"
)

// FullVersion returns the full version and commit hash
func FullVersion() string {
	return fmt.Sprintf("%s@%s", Version, runtime.GOOS)
}
