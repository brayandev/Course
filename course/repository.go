package course

// Repository implements funcs of repository.
type Repository interface{}

// RepositoryImpl as dependecies of repository.
type RepositoryImpl struct{}

func newRepository() *RepositoryImpl {
	return &RepositoryImpl{}
}
