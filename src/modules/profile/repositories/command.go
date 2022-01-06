package repositories

import "locust/src/modules/profile/entities"

type ProfileCommand interface {
	InsertProfile(profile entities.Profile) error
	InsertProfiles(profiles entities.ProfileList) error
}

type ProfileCommandImpl struct {
}

func NewProfileCommand() ProfileCommand {
	return &ProfileCommandImpl{}
}

func (repo *ProfileCommandImpl) InsertProfile(profile entities.Profile) error {
	return nil
}

func (repo *ProfileCommandImpl) InsertProfiles(profile entities.ProfileList) error {
	return nil
}
