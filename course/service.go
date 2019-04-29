package course

import (
	"context"
	"time"

	"github.com/gofrs/uuid"
)

// Service implements service functions.
type Service interface {
	CreateCourse(ctx context.Context, course Course) (string, error)
}

// ServiceImpl represents dependecies of service.
type ServiceImpl struct {
	repository Repository
}

// NewService is a service constructor.
func NewService(repository Repository) *ServiceImpl {
	return &ServiceImpl{
		repository: repository,
	}
}

// CreateCourse is responsible for create a new course.
func (s *ServiceImpl) CreateCourse(ctx context.Context, course Course) (string, error) {
	courseID, err := uuid.NewV4()
	if err != nil {
		return "", err
	}

	course.CourseID = courseID.String()
	course.Creation = time.Now()

	cErr := s.repository.createCourse(ctx, course)
	if cErr != nil {
		return "", cErr
	}

	return course.CourseID, nil
}
