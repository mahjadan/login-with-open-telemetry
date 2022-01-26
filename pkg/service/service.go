package service

import (
	"context"
	"errors"
	enc "github.com/mahjadan/login-with-open-telemetry/pkg/encrypt"
	"github.com/mahjadan/login-with-open-telemetry/pkg/repository"
)

var ErrInvalidUserOrPassword = errors.New("invalid username or password")

type UserService interface {
	Login(ctx context.Context, username, password string) error
	Register(ctx context.Context, username, password string) error
}

type service struct {
	repo    repository.UsersRepository
	appName string
}

// todo generate token on successfull login
func (s service) Login(ctx context.Context, username, password string) error {
	//var span trace.Span
	//ctx, span = otel.Tracer(s.appName).Start(ctx, "login")
	//defer span.End()

	hash, err := s.repo.Get(ctx, username)

	if errors.Is(err, repository.ErrNotFound) {
		//span.RecordError(ErrInvalidUserOrPassword)
		return ErrInvalidUserOrPassword
	}
	if err != nil {
		//span.RecordError(err)
		return err
	}

	if !enc.PasswordMatch(password, hash) {
		//span.RecordError(ErrInvalidUserOrPassword)
		return ErrInvalidUserOrPassword
	}
	return nil
}

func (s service) Register(ctx context.Context, username, password string) error {
	err := s.repo.Set(ctx, username, password)
	if err != nil {
		return err
	}
	return nil
}

func NewService(usersRepository repository.UsersRepository, name string) UserService {
	return &service{
		repo:    usersRepository,
		appName: name,
	}
}
