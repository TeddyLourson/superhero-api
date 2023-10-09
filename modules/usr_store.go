package modules

import (
	usrv1 "buf.build/gen/go/digibear/digibear/protocolbuffers/go/usr/v1"
	pkgtypes "github.com/digibearapp/digibear-api/pkg/types"
	"github.com/google/uuid"
)

type UsrStore interface {
	// ::::: CREATE ::::: //
	// Creates a new Profile inside the DB.
	CreateProfile(profile *usrv1.Profile) (uuid.UUID, error)

	// ::::: READ ::::: //
	// Gets the Profile with the given ID.
	GetProfile(id uuid.UUID) (*usrv1.Profile, error)
	// Checks if the given email is in use.
	CheckEmailInUse(email pkgtypes.Email) (bool, error)
	// Gets the Account with the given ID and all the necessary information to select a profile.
	GetAccountForProfileSelection(id uuid.UUID) (*usrv1.AccountForProfileSelection, error)

	// ::::: UPDATE ::::: //
	// Updates a Profile based on its ID.
	UpdateProfile(profile *usrv1.Profile) (uuid.UUID, error)

	// ::::: DELETE ::::: //
	// Deletes a Profile with the given ID.
	DeleteProfile(id uuid.UUID) (uuid.UUID, error)
}
