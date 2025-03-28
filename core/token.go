package core

import (
	"errors"

	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/models"
	"github.com/pocketbase/pocketbase/tools/auth"
)

// VerifyToken uses PB's authRefresh logic to verify a token.
// It returns the associated user record or an error if invalid.
func VerifyToken(pb *pocketbase.PocketBase, token string) (*models.Record, error) {
	authStore, err := auth.NewRecordAuthStore(pb.Dao(), "users")
	if err != nil {
		return nil, errors.New("failed to create auth store")
	}
	authStore.Token = token

	if err := pb.RefreshToken(authStore, ""); err != nil {
		return nil, errors.New("invalid or expired token")
	}

	return authStore.Record(), nil
}
