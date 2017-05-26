// Use of this source code is governed by a BSD-style license that can be found
// in the LICENSE file.

// puffs-gen-c transpiles a Puffs program to a C program.
//
// The command line arguments list the source Puffs files. If no arguments are
// given, it reads from stdin.
//
// The generated program is written to stdout.
package main

import (
	"os"

	"github.com/google/puffs/cmd/internal/cgen"
	"github.com/google/puffs/lang/generate"
)

func main() {
	if err := generate.Main(&cgen.Generator{Extension: 'c'}); err != nil {
		os.Stderr.WriteString(err.Error() + "\n")
		os.Exit(1)
	}
}