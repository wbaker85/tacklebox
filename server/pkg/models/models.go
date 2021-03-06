package models

import (
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	// ErrDuplicateEmail is generated if somebody tries to sign up with the same email twice
	ErrDuplicateEmail = errors.New("models: duplicate email")

	// ErrInvalidCredentials is generated if a user enters an invalid username and/or password
	ErrInvalidCredentials = errors.New("models: invalid credentials")

	// ErrInvalidUser is generated when something is done for a non-existent user ID
	ErrInvalidUser = errors.New("models: invalid user")

	// ErrInvalidBin is generated when a URL contains a non-existant bin ID
	ErrInvalidBin = errors.New("models: invalid bin")

	// ErrInvalidHook is generated when a hook can't be found
	ErrInvalidHook = errors.New("models: invalid hook")
)

// HookDocument is a BSON representation of a received webhook for storing in mongodb.
type HookDocument struct {
	ID      *primitive.ObjectID `bson:"_id"`
	Content string              `bson:"content"`
}

// HookRecord represents a row in a SQL database containing information about a stored document hook.
type HookRecord struct {
	ID      int
	BinID   string
	HookID  string
	Created time.Time
}

// HookData is the relevant data for a hook, for presenting as a JSON response to a client request
type HookData struct {
	ID      string
	BinID   string
	Content string
	Created time.Time
}

// User is a registered user of the application
type User struct {
	ID             int
	Email          string
	HashedPassword []byte
	Created        time.Time
}

// Bin is essentially a folder for hooks.  Bins are owned by a user.
type Bin struct {
	ID      string
	UserID  int
	Created time.Time
}
