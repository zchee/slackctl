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

// EmojiListCall is created by EmojiService.List method call
type EmojiListCall struct {
	service *EmojiService
}

// List creates a EmojiListCall object in preparation for accessing the emoji.list endpoint
func (s *EmojiService) List() *EmojiListCall {
	var call EmojiListCall
	call.service = s
	return &call
}

// Values returns the EmojiListCall object as url.Values
func (c *EmojiListCall) Values() (url.Values, error) {
	v := url.Values{}
	v.Set(`token`, c.service.token)
	return v, nil
}

// Do executes the call to access emoji.list endpoint
func (c *EmojiListCall) Do(ctx context.Context) (*EmojiListResponse, error) {
	const endpoint = "emoji.list"
	v, err := c.Values()
	if err != nil {
		return nil, err
	}
	var res struct {
		SlackResponse
		*EmojiListResponse
	}
	if err := c.service.client.postForm(ctx, endpoint, v, &res); err != nil {
		return nil, errors.Wrap(err, `failed to post to emoji.list`)
	}
	if !res.OK {
		return nil, errors.New(res.Error.String())
	}

	return res.EmojiListResponse, nil
}
