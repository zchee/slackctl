// Copyright 2017 The slackctl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package slackctl

import (
	"context"

	slack "github.com/lestrrat/go-slack"
	"github.com/pkg/errors"
)

func Post(ctx context.Context, client *slack.Client, channel, text string) error {
	if _, err := client.Chat().PostMessage(channel).AsUser(true).LinkNames(true).Text(text).Do(ctx); err != nil {
		return errors.Wrap(err, "failed to post message")
	}

	return nil
}
