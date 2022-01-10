package bleveProfileRepository

import (
	"locust/domain"

	"github.com/blevesearch/bleve/v2"
)

type bleveProfileRepository struct {
	Index *bleve.Index
}

func NewBleveProfileRepository(index *bleve.Index) domain.ProfileRepository {
	return bleveProfileRepository{
		Index: index,
	}
}

func (r bleveProfileRepository) GetProfileByPeerID(peerID string) (domain.Profile, error) {
	return domain.Profile{}, nil
}

func (r bleveProfileRepository) Store(peerID string, profile *domain.Profile) error {
	return nil
}
