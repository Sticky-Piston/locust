package profile

import "locust/domain"

type profileUsecase struct {
	Repository domain.ProfileRepository
}

func NewProfileUsecase(repository domain.ProfileRepository) domain.ProfileUsecase {
	return &profileUsecase{
		Repository: repository,
	}
}

func (u *profileUsecase) GetProfile() (domain.Profile, error) {
	return u.Repository.GetProfile()
}

func (u *profileUsecase) GetProfileByPeerID(peerID string) (domain.Profile, error) {
	// Try to fetch document from repository
	profile, err := u.Repository.GetProfileByPeerID(peerID)
	if err != nil {
		return domain.Profile{}, nil
	}

	// TODO: Fetch from peer if not in repository

	// TODO: Store fetched profile in repository

	return profile, err
}

func (u *profileUsecase) Store(peerID string, profile *domain.Profile) error {
	return u.Repository.Store(peerID, profile)
}
