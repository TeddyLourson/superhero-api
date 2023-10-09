package modules

import (
	"log"

	usrv1 "buf.build/gen/go/digibear/digibear/protocolbuffers/go/usr/v1"
	utils "github.com/digibearapp/digibear-api/pkg"
	pkgerr "github.com/digibearapp/digibear-api/pkg/err"
	pkgstore "github.com/digibearapp/digibear-api/pkg/store"
	pkgtypes "github.com/digibearapp/digibear-api/pkg/types"
	usrdb "github.com/digibearapp/digibear-api/usr/db"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

// UsrStorePgx implements UsrStore for the Postgres DB.
type UsrStorePgx struct {
	Store pkgstore.StorePgx
}

// NewUsrStorePgx returns a new UsrStorePgx.
func NewUsrStorePgx(store pkgstore.StorePgx) *UsrStorePgx {
	return &UsrStorePgx{
		Store: store,
	}
}

// Creates a new Profile inside the DB.
func (p *UsrStorePgx) CreateProfile(profile *usrv1.Profile) (uuid.UUID, error) {
	q := usrdb.New(p.Store.DB)

	createProfileParams := usrdb.CreateProfileParams{
		ID:             utils.UUIDFromString(profile.Id),
		Username:       profile.Username,
		Code:           utils.StringToText(profile.Code.Value),
		ImagePath:      utils.StringToText(profile.ImagePath.Value),
		AgeRestriction: utils.UInt32ToInt4(profile.AgeRestriction.Value),
		SyncedThemeID:  utils.StringToText(profile.SyncedThemeId.Value),
		AccountID:      utils.UUIDFromString(profile.AccountId),
	}

	res, err := q.CreateProfile(p.Store.Ctx, createProfileParams)
	if nil != err {
		return uuid.UUID{}, pkgerr.NewInternalError(err)
	}
	return res, nil
}

// Gets the Profile with the given ID.
func (p *UsrStorePgx) GetProfile(id uuid.UUID) (*usrv1.Profile, error) {
	q := usrdb.New(p.Store.DB)

	row, err := q.GetProfile(p.Store.Ctx, id)
	if nil != err {
		return nil, pkgerr.NewInternalError(err)
	}
	profile := &usrv1.Profile{
		Id:             row.ID.String(),
		Username:       row.Username,
		AccountId:      row.AccountID.String(),
		AgeRestriction: utils.Int4ToNullUInt32(row.AgeRestriction),
		Code:           utils.TextToNullString(row.Code),
		ImagePath:      utils.TextToNullString(row.ImagePath),
		CreatedAt:      *utils.TimestamptzToString(row.CreatedAt),
		UpdatedAt:      utils.TimestamptzToString(row.UpdatedAt),
	}
	return profile, nil
}

// Checks if the given email is in use.
func (p *UsrStorePgx) CheckEmailInUse(email pkgtypes.Email) (bool, error) {
	q := usrdb.New(p.Store.DB)
	_, err := q.GetAccountIDByEmail(p.Store.Ctx, email.String)
	if nil != err {
		if err == pgx.ErrNoRows {
			return false, nil
		}
		return false, pkgerr.NewInternalError(err)
	}
	return true, nil
}

// Gets the Account with the given ID and all the necessary information to select a profile.
func (p *UsrStorePgx) GetAccountForProfileSelection(id uuid.UUID) (*usrv1.AccountForProfileSelection, error) {
	q := usrdb.New(p.Store.DB)

	rows, err := q.GetAccountForProfileSelection(p.Store.Ctx, id)
	if nil != err {
		return nil, pkgerr.NewInternalError(err)
	}
	if 0 == len(rows) {
		return nil, pkgerr.NewNotFoundError("AccountForProfileSelection", id.String())
	}
	account := &usrv1.Account{
		Id:                rows[0].ID.String(),
		Email:             rows[0].Email,
		AcquiredWatchTime: uint32(rows[0].AcquiredWatchTime),
		CreatedAt:         *utils.TimestamptzToString(rows[0].CreatedAt),
		UpdatedAt:         utils.TimestamptzToString(rows[0].UpdatedAt),
	}
	log.Printf("account: %+v", account)
	var profiles [](*usrv1.Profile)
	for i := range rows {
		row := rows[i]
		profile := &usrv1.Profile{
			Id:             row.ProfileID.String(),
			Username:       row.ProfileUsername,
			Code:           utils.TextToNullString(row.ProfileCode),
			AgeRestriction: utils.Int4ToNullUInt32(row.ProfileAgeRestriction),
			ImagePath:      utils.TextToNullString(row.ProfileImagePath),
			AccountId:      row.ProfileAccountID.String(),
			CreatedAt:      *utils.TimestamptzToString(row.ProfileCreatedAt),
			UpdatedAt:      utils.TimestamptzToString(row.ProfileUpdatedAt),
		}
		profiles = append(profiles, profile)
	}

	log.Printf("profiles: %+v", profiles)

	return &usrv1.AccountForProfileSelection{
		Account:  account,
		Profiles: profiles,
	}, nil
}

// Updates a Profile based on its ID.
func (p *UsrStorePgx) UpdateProfile(profile *usrv1.Profile) (uuid.UUID, error) {
	q := usrdb.New(p.Store.DB)

	updateProfileParams := usrdb.UpdateProfileParams{
		ID:             utils.UUIDFromString(profile.Id),
		Username:       profile.Username,
		Code:           utils.StringToText(profile.Code.Value),
		ImagePath:      utils.StringToText(profile.ImagePath.Value),
		AgeRestriction: utils.UInt32ToInt4(profile.AgeRestriction.Value),
		SyncedThemeID:  utils.StringToText(profile.SyncedThemeId.Value),
		// AccountID:      accountID,
	}

	res, err := q.UpdateProfile(p.Store.Ctx, updateProfileParams)
	if nil != err {
		return uuid.UUID{}, pkgerr.NewInternalError(err)
	}
	return res, nil
}

// Deletes a Profile with the given ID.
func (p *UsrStorePgx) DeleteProfile(id uuid.UUID) (uuid.UUID, error) {
	q := usrdb.New(p.Store.DB)

	res, err := q.DeleteProfile(p.Store.Ctx, id)
	if err != nil {
		return uuid.UUID{}, pkgerr.NewInternalError(err)
	}
	return res, nil
}
