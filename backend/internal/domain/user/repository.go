package user

type Repository interface {
	Create(u *User) error
	Delete(id string) error
	GetUsers() ([]User, error)
	GetUserByUsername(username string) (*User, error)
	UpdatePhoto(id, photo_url string) error
}
