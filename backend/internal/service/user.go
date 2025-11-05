package service

import (
	"context"
	"fmt"
	"mime/multipart"

	"github.com/critiq17/critiqal-site/internal/domain/user"
	"github.com/critiq17/critiqal-site/internal/repository"
	"github.com/critiq17/critiqal-site/internal/storage"
	"gorm.io/gorm"
)

type User struct {
	ID        string
	Username  string
	Email     string
	Password  string
	FirstName string
	LastName  string
	CreatedAt int64
	DeletedAt gorm.DeletedAt
}

type UserService struct {
	repo    user.Repository
	storage storage.Storage
}

func NewUserService(repo user.Repository, storage storage.Storage) *UserService {
	return &UserService{
		repo: repo, storage: storage,
	}
}

func (s *UserService) CreateUser(u *user.User) error {

	err := s.repo.Create(u)
	if err != nil {
		return err
	}

	return nil
}

func (s *UserService) DeleteUser(id string) error {

	err := s.repo.Delete(id)
	if err != nil {
		return err
	}

	return nil
}

func (s *UserService) GetUsers() ([]user.User, error) {
	users, err := s.repo.GetUsers()
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (s *UserService) GetByID(id string) (*user.User, error) {
	user, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *UserService) GetByUsername(username string) (*user.User, error) {
	u, err := s.repo.GetUserByUsername(username)
	if err != nil {
		return nil, err
	}

	return u, nil
}

func (s *UserService) SetUserPhoto(id, photo_url string) error {
	return s.repo.UpdatePhoto(id, photo_url)
}

func (s *UserService) UploadUserPhoto(ctx context.Context, username string, file *multipart.FileHeader) (string, error) {
	path := fmt.Sprintf("avatars/%s/%s", username, file.Filename)

	url, err := s.storage.UploadFile(ctx, file, path)
	if err != nil {
		return "", err
	}

	if err := s.repo.UpdatePhoto(username, url); err != nil {
		return "", err
	}

	return url, nil
}

func (s *UserService) SearchUsers(username string) ([]user.User, error) {
	users, err := s.repo.Search(username)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (s *UserService) Auth(username, password string) (*user.User, bool, error) {
	user, err := s.repo.GetUserByUsername(username)
	if err != nil {
		return nil, false, err
	}

	u := repository.User{
		Password: user.Password,
	}

	return user, u.CheckPassword(password), nil
}
