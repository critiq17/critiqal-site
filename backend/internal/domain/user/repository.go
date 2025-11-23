package user

type Repository interface {
	// CRUD base
	Create(u *User) error
	Delete(id string) error
	GetByID(id string) (*User, error)

	// Get, search, update
	GetUsers() ([]User, error)
	Search(username string) ([]User, error)
	GetUserByUsername(username string) (*User, error)
	UpdatePhoto(username, photo_url string) error
}
