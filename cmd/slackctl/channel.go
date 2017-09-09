// Copyright 2017 The slackctl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	cli "github.com/spf13/cobra"
	"github.com/zchee/slackctl"
)

// channelCmd represents the channels command
var channelCmd = &cli.Command{
	Use:   "channel",
	Short: "Show your team channels",
	RunE:  runChannel,
}

func init() {
	RootCmd.AddCommand(channelCmd)
}

func runChannel(cmd *cli.Command, args []string) error {
	if err := slackctl.Channels(); err != nil {
		return err
	}
	return nil
}
