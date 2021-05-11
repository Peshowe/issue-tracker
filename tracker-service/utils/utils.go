package utils

import (
	"context"
	"errors"
	"log"

	"go.mongodb.org/mongo-driver/bson"
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

//EncodeToBytes encodes anything passed to a BSON byte array
func EncodeToBytes(p interface{}) []byte {

	res, err := bson.Marshal(p)
	if err != nil {
		log.Fatal(err)
	}

	return res
}
