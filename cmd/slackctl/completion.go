// Copyright 2017 The slackctl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"os"

	"github.com/pkg/errors"
	cli "github.com/spf13/cobra"
)

// CompletionCmd represents the channels command
var CompletionCmd = &cli.Command{
	Use:   "completion",
	Short: "Show slackctl shell completion",
	RunE:  runCompletion,
}

func init() {
	RootCmd.AddCommand(CompletionCmd)
}

func runCompletion(cmd *cli.Command, args []string) error {
	if len(args) == 0 {
		return errors.New("slackctl: required shell name [bash, zsh]")
	}
	shell := args[0]

	switch shell {
	case "bash":
		return cmd.GenBashCompletion(os.Stdout)
	case "zsh":
		return cmd.GenZshCompletion(os.Stdout)
	default:
		return errors.Errorf("slackctl: unknown shell name: %s", shell)
	}
}
