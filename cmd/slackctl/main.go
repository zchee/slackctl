// Copyright 2017 The slackctl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"os"

	cli "github.com/spf13/cobra"
)

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cli.Command{
	Use:   "slackctl",
	Short: "A slack management tool",
}

func main() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
