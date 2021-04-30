package utils

import (
	"context"
	"errors"
)

var ErrNotAllowed error = errors.New("not allowed")

type TokenContexKey string

func GetUserFromContext(ctx context.Context) string {
	if v := ctx.Value(TokenContexKey("token")); v != nil {
		return v.(string)
	}
	return ""
}
