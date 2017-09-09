// Copyright 2017 The slackctl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package slackctl

import (
	"context"

	slack "github.com/lestrrat/go-slack"
	"github.com/pkg/errors"
)

func Auth(ctx context.Context, client *slack.Client) (*slack.AuthTestResponse, error) {
	auth, err := client.Auth().Test().Do(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get auth information")
	}
	return auth, nil
}
