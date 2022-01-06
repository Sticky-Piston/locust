package repositories

type Repositories interface {
	ProfileCommand
	ProfileQuery
}

type RepositoriesImpl struct {
	*ProfileCommandImpl
	*ProfileQueryImpl
}

func NewRepository() *RepositoriesImpl {
	return &RepositoriesImpl{}
}
