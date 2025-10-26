package user

import (
	"github.com/critiq17/critiqal-site/internal/domain/user/dto"
)

type UserService struct {
	repo Repository
}

func NewUserService(repo Repository) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (s *UserService) CreateUser(u *dto.UserApi) error {

	addUser := dto.User{
		Username:  u.Username,
		Email:     u.Email,
		Password:  u.Password,
		FirstName: u.FirstName,
		LastName:  u.LastName,
	}

	err := s.repo.Create(&addUser)
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

func (s *UserService) GetUsers() ([]dto.UserApi, error) {

	users, err := s.repo.GetUsers()
	if err != nil {
		return nil, err
	}

	usersApi := ToUsersApi(users)

	return usersApi, nil
}
