// Copyright 2017 The slackctl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package slackctl

import (
	"os"
)

var (
	token = os.Getenv("SLACKCTL_TOKEN")
)
