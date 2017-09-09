package slack

// Auto-generated by internal/cmd/genmethods/genmethods.go. DO NOT EDIT!

import (
	"context"
	"net/url"
	"strconv"

	"github.com/lestrrat/go-slack/objects"
	"github.com/pkg/errors"
)

var _ = strconv.Itoa
var _ = objects.EpochTime(0)

// RTMStartCall is created by RTMService.Start method call
type RTMStartCall struct {
	service *RTMService
}

// Start creates a RTMStartCall object in preparation for accessing the rtm.start endpoint
func (s *RTMService) Start() *RTMStartCall {
	var call RTMStartCall
	call.service = s
	return &call
}

// Values returns the RTMStartCall object as url.Values
func (c *RTMStartCall) Values() (url.Values, error) {
	v := url.Values{}
	v.Set(`token`, c.service.token)
	return v, nil
}

// Do executes the call to access rtm.start endpoint
func (c *RTMStartCall) Do(ctx context.Context) (*RTMResponse, error) {
	const endpoint = "rtm.start"
	v, err := c.Values()
	if err != nil {
		return nil, err
	}
	var res struct {
		SlackResponse
		*RTMResponse
	}
	if err := c.service.client.postForm(ctx, endpoint, v, &res); err != nil {
		return nil, errors.Wrap(err, `failed to post to rtm.start`)
	}
	if !res.OK {
		return nil, errors.New(res.Error.String())
	}

	return res.RTMResponse, nil
}
