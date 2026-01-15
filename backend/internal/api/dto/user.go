package dto

import "github.com/critiq17/critiqal-site/internal/domain/user"

type UserApi struct {
	//UserID    string `json:"user_id"`
	Username  string `json:"username" binding:"required"`
	Email     string `json:"email" binding:"required"`
	Password  string `json:"password"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	PhotoURL  string `json:"photo_url"`
}

type CreateRequest struct {
	Username  string `json:"username" binding:"required"`
	Email     string `json:"email" binding:"required"`
	Password  string `json:"password"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func ToUserApi(u *user.User) *UserApi {
	if u == nil {
		return nil
	}
	return &UserApi{
		//UserID:    u.ID,
		Username:  u.Username,
		Email:     u.Email,
		Password:  "",
		FirstName: u.FirstName,
		LastName:  u.LastName,
		PhotoURL:  u.PhotoURL,
	}
}

func ToUsersApi(users []user.User) []UserApi {
	dtos := make([]UserApi, len(users))
	for i, u := range users {
		dtos[i] = *ToUserApi(&u)
	}
	return dtos
}

func ToDBModel(u *UserApi) *user.User {
	if u == nil {
		return nil
	}

	return &user.User{
		Username:  u.Username,
		Email:     u.Email,
		Password:  u.Password,
		FirstName: u.FirstName,
		LastName:  u.LastName,
		PhotoURL:  u.PhotoURL,
	}
}
