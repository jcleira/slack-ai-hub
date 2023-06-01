package slack

import (
	"context"
	"errors"
	"fmt"

	"github.com/slack-go/slack"
	"go.uber.org/zap"
)

var (
	// ErrInvalidCredentials is returned when the slack token is invalid.
	ErrInvalidCredentials = errors.New("Invalid credentials")

	// ErrRTMError is returned when the RTM connection is broken.
	ErrRTMError = errors.New("RTM error")
)

// SlackCallback is a function that is called when a message is received, this
// function should be provided by the caller.
type SlackCallback func(string) (string, error)

// Slack is a slack client that can be used to send and receive messages using
// the slack RTM API and a callback funtion to perform the desired operations
// in the incoming messages to provide the desired results.
type Slack struct {
	channel string
	logger  *zap.Logger

	slack *slack.Client
}

// New creates a new slack client.
func New(slakToken, channel string, logger *zap.Logger) *Slack {
	return &Slack{
		channel: channel,
		logger:  logger,
		slack:   slack.New(slakToken),
	}
}

// Start starts the slack client and waits for incoming messages, when a
// message is received the callback function is called and the result is sent
// back to the slack channel.
func (s *Slack) Start(ctx context.Context, callback SlackCallback) error {
	rtm := s.slack.NewRTM()
	go rtm.ManageConnection()

	for {
		select {
		case <-ctx.Done():
			return nil

		case message := <-rtm.IncomingEvents:
			if err := s.handleIncomingMessage(message, callback); err != nil {
				return fmt.Errorf("s.handleIncomingMessage: %w", err)
			}
		}
	}
}

func (s *Slack) handleIncomingMessage(
	message slack.RTMEvent, callback SlackCallback) error {
	switch ev := message.Data.(type) {
	case *slack.MessageEvent:
		if ev.SubType == "" && ev.BotID == "" {
			result, err := callback(ev.Text)
			if err != nil {
				s.logger.Error("callback", zap.Error(err))

				return nil
			}

			if err := s.postAnswer(result); err != nil {
				return fmt.Errorf("s.postAnswer: %w", err)
			}
		}

	case *slack.RTMError:
		return ErrRTMError

	case *slack.InvalidAuthEvent:
		return ErrInvalidCredentials
	}

	return nil
}

func (s *Slack) postAnswer(answer string) error {
	if _, _, err := s.slack.PostMessage(s.channel,
		slack.MsgOptionText(answer, false),
	); err != nil {
		return fmt.Errorf("s.slack.PostMessage: %w", err)
	}

	return nil
}
