package usecase

import "locust/domain"

type profileUsecase struct {
	profileRepository domain.ProfileRepository
}

func NewProfileUsecase(profileRepository domain.ProfileRepository) domain.ProfileUsecase {
	return &profileUsecase{
		profileRepository: profileRepository,
	}
}

func (u *profileUsecase) GetByPeerID(peerID string) (domain.Profile, error) {
	profile, err := u.profileRepository.GetByPeerID(peerID)
	if err != nil {
		return domain.Profile{}, err
	}

	return profile, nil
}

func (u *profileUsecase) Store(profile domain.Profile) error {
	err := u.profileRepository.Store(profile)
	if err != nil {
		return err
	}

	return nil
}
