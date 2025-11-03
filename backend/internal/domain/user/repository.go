package user

type Repository interface {
	Create(u *User) error
	Delete(id string) error
	GetUsers() ([]User, error)
	GetByID(id string) (*User, error)
	GetUserByUsername(username string) (*User, error)
	UpdatePhoto(username, photo_url string) error
}
