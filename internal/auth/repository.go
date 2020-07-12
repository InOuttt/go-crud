package auth

import (
	"context"
	"github.com/inouttt/go-crud/internal/entity"
	"github.com/inouttt/go-crud/pkg/dbcontext"
	"github.com/inouttt/go-crud/pkg/log"
	dbx "github.com/go-ozzo/ozzo-dbx"
)

type Repository interface {
	Get(ctx context.Context, username string) (entity.User, error)
}

// repository persists users in database
type repository struct {
	db     *dbcontext.DB
	logger log.Logger
}

// NewRepository creates a new user repository
func NewRepository(db *dbcontext.DB, logger log.Logger) Repository {
	return repository{db, logger}
}

// Get reads the user with the specified ID from the database.
func (r repository) Get(ctx context.Context, username string) (entity.User, error) {
	var user entity.User
	err := r.db.With(ctx).Select().Where(dbx.HashExp{"username": username}).One(&user)
	return user, err
}