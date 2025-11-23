package user

type Service interface {
	// Create user from api
	CreateUser(u *User) error
	// Delete user bu id
	DeleteUser(id string) error
	// GetUsers retrieves all users
	GetUsers() ([]User, error)
	// Set photo for user
	SetUserPhoto(username, photo_url string) error
	// Get user by username
	GetByUsername(username string) (User, error)
	// Search users by username
	Search(username string) (User, error)
}
