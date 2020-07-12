package profile

import (
	"context"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/inouttt/go-crud/internal/entity"
	User "github.com/inouttt/go-crud/internal/user"
	"github.com/inouttt/go-crud/pkg/log"
	"github.com/inouttt/go-crud/pkg/utils"
	// "golang.org/x/crypto/bcrypt"
	"time"
)

// Service encapsulates usecase logic for profiles.
type Service interface {
	Get(ctx context.Context, id string) (Profile, error)
	Query(ctx context.Context, offset, limit int) ([]Profile, error)
	Count(ctx context.Context) (int, error)
	Create(ctx context.Context, input CreateProfileRequest) (Profile, error)
	Update(ctx context.Context, id string, input UpdateProfileRequest) (Profile, error)
	Delete(ctx context.Context, id string) (Profile, error)
}

// Profile represents the data about an profile.
type Profile struct {
	entity.Profile
}

// CreateProfileRequest represents an profile creation request.
type CreateProfileRequest struct {
	UserID string `json:"UserID"`
	Email string `json:"Email"`
	Address string `json:"Address"`
	Password string `json:"Password"`
}

// Validate validates the CreateProfileRequest fields.
func (m CreateProfileRequest) Validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.UserID, validation.Required, validation.Length(0, 128)),
		validation.Field(&m.Email, validation.Required, validation.Length(0, 128), is.Email),
		validation.Field(&m.Address, validation.Required, validation.Length(0, 128)),
		validation.Field(&m.Password, validation.Required, validation.Length(0, 128)),
	)
}

// UpdateProfileRequest represents an profile update request.
type UpdateProfileRequest struct {
	Email string `json:"email"`
	Address string `json:"address"`
	Password string `json:"password"`
}

// Validate validates the CreateProfileRequest fields.
func (m UpdateProfileRequest) Validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.Email, validation.Length(0, 128), is.Email),
		validation.Field(&m.Address, validation.Length(0, 128)),
		validation.Field(&m.Password, validation.Length(0, 128)),
	)
}

type service struct {
	repo   Repository
	logger log.Logger
	user  User.Repository
}

// NewService creates a new profile service.
func NewService(repo Repository, logger log.Logger, user User.Repository) Service {
	return service{repo, logger, user}
}

// Get returns the profile with the specified the profile ID.
func (s service) Get(ctx context.Context, id string) (Profile, error) {
	profile, err := s.repo.Get(ctx, id)
	if err != nil {
		return Profile{}, err
	}
	return Profile{profile}, nil
}

// Create creates a new profile.
func (s service) Create(ctx context.Context, req CreateProfileRequest) (Profile, error) {
	if err := req.Validate(); err != nil {
		return Profile{}, err
	}
	id := entity.GenerateID()
	now := time.Now()
	pass := utils.GetHashedPassword(req.Password)
	err := s.user.Create(ctx, entity.User{
		Username:  req.UserID,
		Password:  pass,
	})
	if err != nil {
		return Profile{}, err
	}
	
	err = s.repo.Create(ctx, entity.Profile{
		Id:        id,
		User_id:    req.UserID,
		Email:     req.Email,
		Address:   req.Address,
		Created_at: now,
	})
	if err != nil {
		return Profile{}, err
	}

	return s.Get(ctx, id)
}

// Update updates the profile with the specified ID.
func (s service) Update(ctx context.Context, id string, req UpdateProfileRequest) (Profile, error) {
	if err := req.Validate(); err != nil {
		return Profile{}, err
	}

	profile, err := s.Get(ctx, id)
	if err != nil {
		return profile, err
	}
	// user, err := s.user.Get(ctx, profile.User_id)
	// if err != nil {
	// 	return profile, err
	// }
	// pass := utils.GetHashedPassword(req.Password)
	// user.Password = pass

	// if err := s.user.Update(ctx, user); err != nil {
	// 	return profile, err
	// }

	profile.Email = req.Email
	profile.Address = req.Address
	profile.Updated_at = time.Now()

	if err := s.repo.Update(ctx, profile.Profile); err != nil {
		return profile, err
	}
	return profile, nil
}

// Delete deletes the profile with the specified ID.
func (s service) Delete(ctx context.Context, id string) (Profile, error) {
	profile, err := s.Get(ctx, id)
	if err != nil {
		return Profile{}, err
	}
	if err = s.repo.Delete(ctx, id); err != nil {
		return Profile{}, err
	}
	return profile, nil
}

// Count returns the number of profiles.
func (s service) Count(ctx context.Context) (int, error) {
	return s.repo.Count(ctx)
}

// Query returns the profiles with the specified offset and limit.
func (s service) Query(ctx context.Context, offset, limit int) ([]Profile, error) {
	items, err := s.repo.Query(ctx, offset, limit)
	if err != nil {
		return nil, err
	}
	result := []Profile{}
	for _, item := range items {
		result = append(result, Profile{item})
	}
	return result, nil
}