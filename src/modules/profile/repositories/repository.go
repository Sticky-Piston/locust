package repositories

type Repositories interface {
	ProfileCommand
}

type RepositoriesImpl struct {
}

func NewRepository() *RepositoriesImpl {
	return &RepositoriesImpl{}
}
