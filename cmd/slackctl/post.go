// Copyright 2017 The slackctl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"context"

	slack "github.com/lestrrat/go-slack"
	"github.com/pkg/errors"
	cli "github.com/spf13/cobra"
	"github.com/zchee/slackctl"
)

var (
	// postCmd represents the post command
	postCmd = &cli.Command{
		Use:        "post <channel name> <text>",
		Short:      "post text to any slack channel",
		RunE:       runPost,
		ArgAliases: []string{"channel", "text"},
	}
)

func init() {
	RootCmd.AddCommand(postCmd)
}

func runPost(cmd *cli.Command, args []string) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	if len(args) < 2 {
		return errors.Errorf("%s needs a channel name and post text", cmd.Name())
	}
	channel, text := args[0], args[1]

	if err := slackctl.Post(ctx, slack.New(token), channel, text); err != nil {
		return err
	}

	return nil
}
