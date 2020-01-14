// +build !windows

// Package main is a simple wrapper of the real sled entrypoint package.
//
// This package should NOT be extended or modified in any way; to modify the
// sled binary, work in the `gitlab.com/<USER>/sled/cmd` package.
//
package main

import (
	sled "github.com/junland/sled/cmd"
)

func main() {
	sled.Run()
}
