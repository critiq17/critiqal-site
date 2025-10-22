package user

import (
	"github.com/critiq17/critiqal-site/internal/domain/user/dto"
)

func ToUserApi(u *dto.User) *dto.UserApi {
	if u == nil {
		return nil
	}
	return &dto.UserApi{
		Username:  u.Username,
		Email:     u.Email,
		Password:  u.Password,
		FirstName: u.FirstName,
		LastName:  u.LastName,
	}
}

func ToUsersApi(users []dto.User) []dto.UserApi {
	dtos := make([]dto.UserApi, len(users))
	for i, u := range users {
		dtos[i] = *ToUserApi(&u)
	}
	return dtos
}

func ToDBModel(u *dto.User) *dto.User {
	if u == nil {
		return nil
	}

	return &dto.User{
		ID:        u.ID,
		Username:  u.Username,
		Email:     u.Email,
		Password:  u.Password,
		FirstName: u.FirstName,
		LastName:  u.LastName,
	}
}
