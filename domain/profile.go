package domain

type Profile struct {
	PeerID    string `json:"peerID"`
	Title     string `json:"title"`
	Summary   string `json:"summary"`
	CreatedAt string `json:"created_at"`
}

type ProfileRepository interface {
	GetByPeerID(peerID string) (Profile, error)
	Store(Profile) error
}

type ProfileUsecase interface {
	GetByPeerID(peerID string) (Profile, error)
	Store(Profile) error
}
