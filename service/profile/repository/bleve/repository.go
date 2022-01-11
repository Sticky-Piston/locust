package bleveProfileRepository

import (
	"locust/domain"
	"log"

	"github.com/blevesearch/bleve/v2"
)

type bleveProfileRepository struct {
	Index bleve.Index
}

func NewBleveProfileRepository(index bleve.Index) domain.ProfileRepository {
	return bleveProfileRepository{
		Index: index,
	}
}

func (r bleveProfileRepository) GetProfile() (domain.Profile, error) {
	return domain.Profile{
		Author:  "blaat",
		Payload: "blaat",
	}, nil
}

func (r bleveProfileRepository) GetProfileByPeerID(peerID string) (domain.Profile, error) {
	document, err := r.Index.Document(peerID)
	if err != nil {
		return domain.Profile{}, nil
	}

	log.Println(document)

	// TODO unmarshal bleve stuff
	return domain.Profile{}, nil
}

func (r bleveProfileRepository) Store(peerID string, profile *domain.Profile) error {
	err := r.Index.Index(peerID, profile)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
