package repositories

import (
	"locust/src/modules/profile/entities"

	badger "github.com/dgraph-io/badger"
)

type ProfileCommand interface {
	InsertProfile(profile entities.Profile) error
	InsertProfiles(profiles entities.ProfileList) error
}

type ProfileCommandImpl struct {
	db  *badger.DB
	txn *badger.Txn
}

func NewProfileCommand(db *badger.DB) ProfileCommand {
	return &ProfileCommandImpl{
		db: db,
	}
}

func (repo *ProfileCommandImpl) InsertProfile(profile entities.Profile) error {
	return nil
}

func (repo *ProfileCommandImpl) InsertProfiles(profile entities.ProfileList) error {
	return nil
}
