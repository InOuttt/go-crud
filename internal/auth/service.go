package auth

import (
	"context"
	"github.com/dgrijalva/jwt-go"
	"github.com/inouttt/go-crud/internal/errors"
	"github.com/inouttt/go-crud/pkg/log"
	"github.com/inouttt/go-crud/pkg/utils"
	"time"
)

type Service interface {
	Login(ctx context.Context, username, password string) (string, error)
}

type Identity interface {
	GetID() string
	GetName() string
	GetPassword() string
}

type service struct {
	signingKey      string
	tokenExpiration int
	logger          log.Logger
	repo            Repository
}

func NewService(signingKey string, tokenExpiration int, logger log.Logger, repo Repository) Service {
	return service{signingKey, tokenExpiration, logger, repo}
}

func (s service) Login(ctx context.Context, username, password string) (string, error) {
	if identity := s.authenticate(ctx, username, password); identity != nil {
		return s.generateJWT(identity)
	}
	return "", errors.Unauthorized("")
}

func (s service) authenticate(ctx context.Context, username, password string) Identity {
	logger := s.logger.With(ctx, "user", username)
	user := s.getUser(ctx, username)
	if user != nil {
		userPass := []byte(user.GetPassword())
		pass := utils.ComparePassword(userPass, password)
		if pass == nil {
			logger.Infof("authentication successful")
			return user
		}
	}

	logger.Infof("authentication failed")
	return nil
}

func (s service) getUser(ctx context.Context, username string) Identity {
	user, err := s.repo.Get(ctx, username)
	if err != nil {
		return nil
	}
	return user
}

// generateJWT generates a JWT that encodes an identity.
func (s service) generateJWT(identity Identity) (string, error) {
	return jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":   identity.GetID(),
		"name": identity.GetName(),
		"exp":  time.Now().Add(time.Duration(s.tokenExpiration) * time.Hour).Unix(),
	}).SignedString([]byte(s.signingKey))
}
