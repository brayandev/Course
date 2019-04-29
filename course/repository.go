package course

import (
	"context"
	"time"

	"go.uber.org/zap"
	mgo "gopkg.in/mgo.v2"
)

// Repository implements funcs of repository.
type Repository interface {
	createCourse(ctx context.Context, course Course) error
}

// RepositoryImpl as dependecies of repository.
type RepositoryImpl struct {
	session      *mgo.Session
	dbName       string
	dbCollection string
	logger       *zap.Logger
}

// NewRepository is a repository constructor.
func NewRepository(session *mgo.Session, dbName, dbCollection string) *RepositoryImpl {
	return &RepositoryImpl{
		session:      session,
		dbName:       dbName,
		dbCollection: dbCollection,
	}
}

func (r *RepositoryImpl) createCourse(ctx context.Context, course Course) error {
	start := time.Now()
	c := r.session.DB(r.dbName).C(r.dbCollection)

	err := c.Insert(course)
	r.logger.Info(
		"db insert",
		zap.Duration("duration", time.Since(start)),
		zap.Any("course", course),
		zap.NamedError("error", err),
	)
	return err
}

// NewMongoDB is a mongoDB constructor.
func NewMongoDB(endpoint string) (*mgo.Session, error) {
	var mgoSession *mgo.Session
	if mgoSession == nil {
		var err error
		mgoSession, err = mgo.Dial(endpoint)
		if err != nil {
			return nil, err
		}
	}

	return mgoSession.Clone(), nil
}
