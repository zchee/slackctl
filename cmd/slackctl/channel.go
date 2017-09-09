// Copyright 2017 The slackctl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"context"
	"fmt"
	"os"

	slack "github.com/lestrrat/go-slack"
	"github.com/pkg/errors"
	cli "github.com/spf13/cobra"
	"github.com/zchee/slackctl"
)

// channelCmd represents the channels command
var channelCmd = &cli.Command{
	Use:     "channel",
	Short:   "Show your team channels",
	PreRunE: preChannel,
	RunE:    runChannel,
}

var (
	channelSortby string
)

func init() {
	RootCmd.AddCommand(channelCmd)

	channelCmd.Flags().StringVarP(&channelSortby, "sort", "s", "", "sort header name")
}

func preChannel(cmd *cli.Command, args []string) error {
	switch channelSortby {
	case "", "name", "member":
		// nothing to do
	default:
		return errors.Errorf("slackctl: invalid sort header name: %s", channelSortby)
	}
	return nil
}

func runChannel(cmd *cli.Command, args []string) error {
	client := slack.New(token)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	auth, err := slackctl.Auth(ctx, client)
	if err != nil {
		return err
	}
	fmt.Fprintf(os.Stderr, "Team: %s\n", auth.Team)

	if err := slackctl.Channels(ctx, client, channelSortby); err != nil {
		return err
	}
	return nil
}
