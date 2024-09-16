package service

import (
	"checkmarx/internal/domain/entity"
	"checkmarx/internal/domain/repository"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base32"
	"errors"
	"strings"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	repo repository.AuthRepository
}

func NewAuthService(repo repository.AuthRepository) *AuthService {
	return &AuthService{repo: repo}
}

func (as *AuthService) Get(u *entity.User) (*entity.User, error) {
	if u.Email == "" || u.Password == "" {
		return nil, errors.New("missing fields")
	}

	u.Email = strings.ToLower(u.Email)

	return as.repo.Get(u.Email)
}

func (as *AuthService) Login(u *entity.User) (*entity.Credentials, error) {
	if u.Email == "" || u.Password == "" {
		return nil, errors.New("missing fields")
	}

	u.Email = strings.ToLower(u.Email)

	user, err := as.repo.Get(u.Email)
	if err != nil {
		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(u.Password)); err != nil {
		return nil, errors.New("wrong credentials")
	}

	token, err := as.GenerateToken(user)
	if err != nil {
		return nil, err
	}

	if err := as.repo.InsertToken(token); err != nil {
		return nil, err
	}

	return &entity.Credentials{UserID: user.ID, Email: user.Email, Token: *token}, nil
}

func (as *AuthService) GenerateToken(u *entity.User) (*entity.Token, error) {
	token := &entity.Token{
		UserID:    u.ID,
		Scope:     "Auth",
		ExpiresAt: time.Now().Add(72 * time.Hour),
	}

	randomBytes := make([]byte, 16)

	_, err := rand.Read(randomBytes)
	if err != nil {
		return nil, err
	}

	token.PlainText = base32.StdEncoding.WithPadding(base32.NoPadding).EncodeToString(randomBytes)

	hash := sha256.Sum256([]byte(token.PlainText))
	token.Hash = hash[:]

	return token, nil
}

func (as *AuthService) Create(u *entity.User) (*entity.Token, error) {
	if u.Email == "" || u.Password == "" {
		return nil, errors.New("missing fields")
	}

	u.Email = strings.ToLower(u.Email)

	hash, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	u.Password = string(hash)

	token, err := as.GenerateToken(u)
	if err != nil {
		return nil, err
	}

	as.repo.Insert(u, token)

	return token, nil
}

func (as *AuthService) GetByToken(t string) (*entity.User, error) {
	token := sha256.Sum256([]byte(t))

	user, err := as.repo.GetByToken(token)
	if err != nil {
		return nil, err
	}

	return user, nil
}
