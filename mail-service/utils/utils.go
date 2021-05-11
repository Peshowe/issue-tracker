package utils

import (
	"context"
	"errors"
)

var ErrNotAllowed error = errors.New("not allowed")

//TokenContexKey is the custom type for the key that retrieves the token stored in the context
type TokenContexKey string

//GetUserFromContext retrieves the currently authenticated user from the context
func GetUserFromContext(ctx context.Context) string {
	if v := ctx.Value(TokenContexKey("token")); v != nil {
		return v.(string)
	}
	return ""
}
