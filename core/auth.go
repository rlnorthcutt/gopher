package core

import (
	"errors"

	"github.com/pocketbase/pocketbase/models"
	"github.com/pocketbase/pocketbase/tools/auth"
)

// AuthService wraps PocketBase's auth functionality.
type AuthService struct {
	pb *pocketbase.PocketBase
}

// NewAuthService returns a new AuthService using the given PocketBase instance.
func NewAuthService(pb *pocketbase.PocketBase) *AuthService {
	return &AuthService{pb: pb}
}

// Login authenticates a user via email and password.
func (a *AuthService) Login(email, password string) (*models.Record, error) {
	authData, err := a.pb.Dao().FindAuthRecordByEmail("users", email)
	if err != nil {
		return nil, errors.New("invalid email or password")
	}

	if !auth.CheckPassword(password, authData.Get("password").(string)) {
		return nil, errors.New("invalid email or password")
	}

	return authData, nil
}

// LoginWithPassword uses email + password to authenticate a user.
func (a *AuthService) LoginWithPassword(email, password string) (*models.Record, error) {
	return a.pb.Dao().AuthRecordLogin("users", email, password)
}

// Signup registers a new user in the users collection.
func (a *AuthService) Signup(email, password string) (*models.Record, error) {
	record := models.NewRecord(a.pb.Dao().ModelCollection("users"))

	record.Set("email", email)
	record.Set("password", password)
	record.Set("passwordConfirm", password)

	if err := a.pb.Dao().SaveRecord(record); err != nil {
		return nil, err
	}

	return record, nil
}

// GetUserByID fetches a user by ID.
func (a *AuthService) GetUserByID(id string) (*models.Record, error) {
	return a.pb.Dao().FindRecordById("users", id)
}

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