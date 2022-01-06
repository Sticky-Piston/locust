package entities

type Profile struct {
	Title   string
	Summary string
	Skills  []string
}

type ProfileList []*Profile
