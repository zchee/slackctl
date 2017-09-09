package slack

import (
	"context"

	"github.com/lestrrat/go-slack/objects"
	"golang.org/x/oauth2"
)

// Logger is an interface for logging/tracing the client's
// execution.
//
// In particular, `Debugf` will only be called if `WithDebug`
// is provided to the constructor.
type Logger interface {
	Debugf(context.Context, string, ...interface{})
	Infof(context.Context, string, ...interface{})
}

const (
	ParseFull = "full"
	ParseNone = "none"
)

type ControlSequence interface {
	Data() string
	Surface() string
	String() string
}

type ChannelLink struct {
	ID      string
	Channel string
}

type UserLink struct {
	ID       string
	Username string
}

type ExternalLink struct {
	URL  string
	Text string
}

// DefaultSlackAPIEndpoint contains the prefix used for Slack REST API
const (
	DefaultAPIEndpoint         = "https://slack.com/api/"
	DefaultOAuth2AuthEndpoint  = "https://slack.com/oauth/authorize"
	DefaultOAuth2TokenEndpoint = "https://slack.com/api/oauth.access"
)

// Oauth2Endpoint contains the Slack OAuth2 endpoint configuration
var OAuth2Endpoint = oauth2.Endpoint{
	AuthURL:  DefaultOAuth2AuthEndpoint,
	TokenURL: DefaultOAuth2TokenEndpoint,
}

type Client struct {
	auth         *AuthService
	bots         *BotsService
	channels     *ChannelsService
	chat         *ChatService
	emoji        *EmojiService
	oauth        *OAuthService
	reactions    *ReactionsService
	rtm          *RTMService
	users        *UsersService
	usersProfile *UsersProfileService
	debug        bool
	slackURL     string
	token        string
}

// SlackResponse is the general response part given by all
// slack API response.
type SlackResponse struct {
	OK        bool          `json:"ok"`
	ReplyTo   int           `json:"reply_to,omitempty"`
	Error     ErrorResponse `json:"error,omitempty"`
	Timestamp string        `json:"ts"`
}

// ErrorResponse wraps errors returned by Slack. It's usually a string,
// but it could be a structure.
// https://api.slack.com/rtm#handling_responses
type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"msg"`
}

// AuthTestResponse is the data structure response from auth.test
type AuthTestResponse struct {
	URL    string `json:"url"`
	Team   string `json:"team"`
	User   string `json:"user"`
	TeamID string `json:"team_id"`
	UserID string `json:"user_id"`
}

type ChannelsHistoryResponse struct {
	HasMore  bool                `json:"has_more"`
	Latest   string              `json:"latest"`
	Messages objects.MessageList `json:"messages"`
}

type ChatResponse struct {
	Channel   string      `json:"channel"`
	Timestamp string      `json:"ts"`
	Message   interface{} `json:"message"` // TODO
}

type EmojiListResponse map[string]string

type OAuthAccessResponse struct {
	AccessToken string
	Scope       string
}

// ReactionsGetResponse represents the response obtained from
// reactions.get API (https://api.slack.com/methods/reactions.get)
type ReactionsGetResponse struct {
	Channel string           `json:"channel"`
	Message *objects.Message `json:"message"`
	File    *objects.File    `json:"file"`
	Comment string           `json:"comment"`
}

type ReactionsGetResponseList []ReactionsGetResponse
type ReactionsListResponse struct {
	Items  ReactionsGetResponseList `json:"items"`
	Paging Paging                   `json:"paging"`
}

type RTMResponse struct {
	URL      string               `json:"url"`
	Self     *objects.UserDetails `json:"self"`
	Team     *objects.Team        `json:"team"`
	Users    []*objects.User      `json:"users"`
	Channels []*objects.Channel   `json:"channels"`
	Groups   []*objects.Group     `json:"groups"`
	Bots     []*objects.Bot       `json:"bots"`
	IMs      []*objects.IM        `json:"ims"`
}

type Paging struct {
	Count int `json:"count"`
	Total int `json:"total"`
	Page  int `json:"page"`
	Pages int `json:"pages"`
}

// InteractiveButtonRequest is a request that is sent when a user
// hits a Slack button. Note: this is experimental
type InteractiveButtonRequest struct {
	ActionTimestamp string             `json:"action_ts"`
	Actions         objects.ActionList `json:"actions"`
	AttachmentID    int                `json:"attachment_id,string"`
	CallbackID      string             `json:"callback_id"`
	Channel         struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	} `json:"channel"`
	IsAppUnfurl      bool                    `json:"is_app_unfurl"`
	MessageTimestamp string                  `json:"message_ts"`
	OriginalMessage  *objects.Message        `json:"original_message"`
	Options          objects.OptionList      `json:"options"`
	OptionGroups     objects.OptionGroupList `json:"option_groups"`
	ResponseURL      string                  `json:"response_url"`
	Team             struct {
		Domain string `json:"domain"`
		ID     string `json:"id"`
	} `json:"team"`
	Token string `json:"token"`
	User  struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	} `json:"user"`
}

type StarsListResponse struct {
	Items  StarredItemList `json:"items"`
	Paging Paging          `json:"paging"`
}
type StarredItem interface{}
type StarredItemList []StarredItem
