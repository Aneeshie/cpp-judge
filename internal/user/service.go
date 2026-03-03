package user

import (
	"context"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) RegisterUser(ctx context.Context, input UserInput) (string, error) {
	email := strings.ToLower(input.Email)

	hashedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(input.HashPassword),
		bcrypt.DefaultCost,
	)

	if err != nil {
		return "", err
	}

	return s.repo.CreateUser(ctx, UserInput{
		Email:        email,
		HashPassword: string(hashedPassword),
	})
}

func (s *Service) Login(ctx context.Context, email string, enteredPassword string) bool {
	u, err := s.repo.GetUserByEmail(ctx, email)

	if err != nil {
		return false
	}

	err = bcrypt.CompareHashAndPassword([]byte(u.PasswordHash), []byte(enteredPassword))

	if err != nil {
		return false
	}

	return true
}
