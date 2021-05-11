package grpc

import (
	"context"

	"github.com/Peshowe/issue-tracker/mail-service/grpc-contract/mail-service/v1/mailer"
	mlr "github.com/Peshowe/issue-tracker/mail-service/mailer"
)

//mailServer implements the gRPC interface and calls the business logic for the mail service
type mailServer struct {
	mailService mlr.MailService
	mailer.UnimplementedMailServiceServer
}

func encodeUserPreference(userPreference *mlr.UserPreference) *mailer.UserPreference {
	return &mailer.UserPreference{
		User:                 userPreference.User,
		IsMailNotificationOn: userPreference.IsMailNotificationOn,
	}
}

func decodeUserPreference(userPreference *mailer.UserPreference) *mlr.UserPreference {
	return &mlr.UserPreference{
		User:                 userPreference.User,
		IsMailNotificationOn: userPreference.IsMailNotificationOn,
	}
}

func (m *mailServer) GetUserPreference(ctx context.Context, request *mailer.UserPreferenceRequest) (*mailer.UserPreference, error) {
	user := request.GetUser()
	userPreference, err := m.mailService.GetUserPreference(ctx, user)

	if err != nil {
		return nil, err
	}

	return encodeUserPreference(userPreference), nil
}

func (m *mailServer) SetUserPreference(ctx context.Context, userPreference *mailer.UserPreference) (*mailer.GenericResponse, error) {
	err := m.mailService.SetUserPreference(ctx, decodeUserPreference(userPreference))

	if err != nil {
		return nil, err
	}

	return &mailer.GenericResponse{}, nil
}
