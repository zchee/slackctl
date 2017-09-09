// Copyright 2017 The slackctl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package slackctl

import (
	"context"
	"fmt"
	"os"
	"sort"
	"strconv"
	"text/tabwriter"

	slack "github.com/lestrrat/go-slack"
)

func Channels(sortby string) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	client := slack.New(token)
	authresp, err := client.Auth().Test().Do(ctx)
	if err != nil {
		return err
	}
	fmt.Fprintf(os.Stderr, "Team: %s\n", authresp.Team)

	channels, err := client.Channels().List().ExclArchived(true).Do(ctx)
	if err != nil {
		return err
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
		tw.Write([]byte(fmt.Sprintf("%s\t%s\n", ch.Name, strconv.Itoa(len(ch.Members)))))
	}
	if err := tw.Flush(); err != nil {
		return err
	}

	return nil
}
