package badgerRepository

import (
	"encoding/json"
	"locust/domain"
	"time"

	"github.com/dgraph-io/badger"
)

type badgerProfileRepository struct {
	Conn *badger.DB
}

func NewProfileRepository(Conn *badger.DB) domain.ProfileRepository {
	return &badgerProfileRepository{Conn}
}

func (repo *badgerProfileRepository) GetByPeerID(peerID string) (domain.Profile, error) {
	return domain.Profile{}, nil
}

func (repo *badgerProfileRepository) Store(profile domain.Profile) error {
	serializedProfile, err := json.Marshal(profile)
	if err != nil {
		return err
	}

	err = repo.Conn.Update(func(txn *badger.Txn) error {
		// Create a new entry and set a TTL
		e := badger.NewEntry([]byte(profile.PeerID), serializedProfile).WithTTL(time.Hour)
		err := txn.SetEntry(e)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}
