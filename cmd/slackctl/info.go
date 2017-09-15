// Copyright 2017 The slackctl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"context"
	"os"
	"text/tabwriter"

	slack "github.com/lestrrat/go-slack"
	"github.com/pkg/errors"
	cli "github.com/spf13/cobra"
	"github.com/zchee/slackctl"
)

// infoCmd represents the info command
var infoCmd = &cli.Command{
	Use:   "info",
	Short: "show your logined slack information",
	RunE:  runInfo,
}

func init() {
	RootCmd.AddCommand(infoCmd)
}

func runInfo(cmd *cli.Command, args []string) error {
	if token == "" {
		return errors.New("SLACKCTL_TOKEN is empty")
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	auth, err := slackctl.Auth(ctx, slack.New(token))
	if err != nil {
		return err
	}
	tw := tabwriter.NewWriter(os.Stdout, 2, 8, 1, ' ', tabwriter.AlignRight)
	tw.Write([]byte("URL\t: " + auth.URL + "\n"))
	tw.Write([]byte("Team\t: " + auth.Team + "\n"))
	tw.Write([]byte("User\t: " + auth.User + "\n"))
	if err := tw.Flush(); err != nil {
		return errors.Wrap(err, "could not flush tabwriter")
	}

	return nil
}
