package course

// Service implements service functions.
type Service interface{}

// ServiceImpl represents dependecies of service.
type ServiceImpl struct{}

// NewService is a service constructor.
func NewService(repository Repository) *ServiceImpl {
	return &ServiceImpl{}
}
