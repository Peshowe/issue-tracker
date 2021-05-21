package mailer

import "context"

type UserPreference struct {
	User                 string `bson:"user,omitempty"`
	IsMailNotificationOn bool   `bson:"is_mail_notification_on,omitempty"`
}

type MailRepo interface {
	SetUserPreference(ctx context.Context, userPreference *UserPreference) error
	GetUserPreference(ctx context.Context, user string) (*UserPreference, error)
}
