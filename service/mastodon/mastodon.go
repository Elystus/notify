package mastodon

import (
    "context"
    "github.com/mattn/go-mastodon"
    "github.com/pkg/errors"
)

// twilioClient abstracts twilio-go MessageService for writing unit tests
//
//go:generate mockery --name=mastodonClient --output=. --case=underscore --inpackage
type mastodonClient interface {
    PostStatus(ctx context.Context, toot *mastodon.Toot) (*mastodon.Status, error)
}

// Compile-time check that mastodon.MessageService satisfies twilioClient interface.
var _ mastodonClient = new(mastodon.Client)

// Service encapsulates the Twilio Message Service client along with internal state for storing recipient phone numbers.
type Service struct {
    client     *mastodon.Client
    recipients []string
}

// New returns a new instance of Twilio notification service.
func New(serverUrl string, clientID string, clientSecret string, accessToken string) (*Service, error) {
    mastodonConf := mastodon.Config{
        Server:       serverUrl,
        ClientID:     clientID,
        ClientSecret: clientSecret,
        AccessToken:  accessToken,
    }
    client := mastodon.NewClient(&mastodonConf)

    svc := &Service{
        client: client,
    }
    return svc, nil
}

// AddReceivers takes strings of recipient account names and appends them to the internal recipient users slice.
// The Send method will send a given message to all those phone numbers.
func (s *Service) AddReceivers(recipients ...string) {
    s.recipients = append(s.recipients, recipients...)
}

// Send takes a message subject and a message body and sends them to all previously set usernames
func (s *Service) Send(ctx context.Context, subject, message string) error {
    for _, recipient := range s.recipients {
        select {
        case <-ctx.Done():
            return ctx.Err()
        default:
            tootBody := recipient + " - " + subject + "\n" + message
            newToot := mastodon.Toot{
                Status:     tootBody,
                Visibility: mastodon.VisibilityDirectMessage,
            }

            _, err := s.client.PostStatus(ctx, &newToot)
            if err != nil {
                return errors.Wrapf(err, "failed to send message to account '%s' using Mastodon", recipient)
            }
        }
    }

    return nil
}
