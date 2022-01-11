package domain

type Profile struct {
	Payload string
	Author  string
}

type ProfileRepository interface {
	GetProfile() (Profile, error)
	GetProfileByPeerID(peerID string) (Profile, error)
	Store(peerID string, profile *Profile) error
}

type ProfileUsecase interface {
	GetProfile() (Profile, error)                      // Serves local profile
	GetProfileByPeerID(peerID string) (Profile, error) // Servers profile fetched from repository or from node if it's not available
	Store(peerID string, profile *Profile) error       // Stores a profile
}
