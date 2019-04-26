package course

// Repository implements funcs of repository.
type Repository interface{}

// RepositoryImpl as dependecies of repository.
type RepositoryImpl struct{}

// NewRepository is a repository constructor.
func NewRepository() *RepositoryImpl {
	return &RepositoryImpl{}
}
