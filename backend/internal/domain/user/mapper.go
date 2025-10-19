package user

import (
	"github.com/critiq17/critiqal-site/internal/domain/user/dto"
)

func ToUserApi(u *dto.UserApi) *dto.UserApi {
	if u == nil {
		return nil
	}

	return &dto.UserApi{
		ID:        u.ID,
		Username:  u.Username,
		Email:     u.Email,
		Password:  u.Password,
		FirstName: u.FirstName,
		LastName:  u.LastName,
	}
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
