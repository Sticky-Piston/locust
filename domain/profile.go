package domain

type Profile struct {
	Payload string
	Author  string
}

type ProfileRepository interface {
	GetProfileByPeerID(peerID string) (Profile, error)
	Store(peerID string, profile *Profile) error
}

type ProfileUsecase interface {
	GetProfileByPeerID(peerID string) (Profile, error)
	Store(peerID string, profile *Profile) error
}
