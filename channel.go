// Copyright 2017 The slackctl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package slackctl

import (
	"context"
	"fmt"
	"os"
	"sort"
	"text/tabwriter"

	slack "github.com/lestrrat/go-slack"
	"github.com/pkg/errors"
)

func Channels(sortby string) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	client := slack.New(token)
	auth, err := client.Auth().Test().Do(ctx)
	if err != nil {
		return errors.Wrap(err, "failed to get auth information")
	}
	fmt.Fprintf(os.Stderr, "Team: %s\n", auth.Team)

	channels, err := client.Channels().List().ExclArchived(true).Do(ctx)
	if err != nil {
		return errors.Wrap(err, "failed to get channel list")
	}

	if sortby != "" {
		switch sortby {
		case "name":
			sort.Slice(channels, func(i, j int) bool {
				return channels[i].Name < channels[j].Name
			})
		case "member":
			sort.Slice(channels, func(i, j int) bool {
				return len(channels[j].Members) < len(channels[i].Members)
			})
		}
	}

	tw := tabwriter.NewWriter(os.Stdout, 1, 8, 1, '\t', 0)
	for _, ch := range channels {
		if _, err := tw.Write([]byte(fmt.Sprintf("%s\t%d\n", ch.Name, len(ch.Members)))); err != nil {
			return errors.Wrap(err, "could not write to tabwriter")
		}
	}
	if err := tw.Flush(); err != nil {
		return errors.Wrap(err, "failed to flush tabwriter")
	}

	return nil
}
