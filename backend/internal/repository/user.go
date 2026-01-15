package repository

import (
	"errors"

	"github.com/critiq17/critiqal-site/internal/domain/user"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	ID        string `gorm:"primaryKey"`
	Username  string `gorm:"uniqueIndex;not null"`
	Email     string `gorm:"uniqueIndex;not null"`
	Password  string `gorm:"not null"`
	FirstName string
	LastName  string
	PhotoURL  string         `gorm:"default:null"`
	CreatedAt int64          `gorm:"autoCreateTime:milli"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type UserRepository struct {
	db *gorm.DB
}

func (m *User) toDomain() *user.User {
	return &user.User{
		ID:        m.ID,
		Username:  m.Username,
		Email:     m.Email,
		Password:  m.Password,
		FirstName: m.FirstName,
		LastName:  m.LastName,
		PhotoURL:  m.PhotoURL,
		CreatedAt: m.CreatedAt,
		DeletedAt: m.DeletedAt,
	}
}
func toDomainUsers(models []User) []user.User {
	users := make([]user.User, len(models))
	for i, m := range models {
		users[i] = *m.toDomain()
	}
	return users
}

func fromDomain(u *user.User) *User {
	return &User{
		ID:        u.ID,
		Username:  u.Username,
		Email:     u.Email,
		Password:  u.Password,
		FirstName: u.FirstName,
		LastName:  u.LastName,
		PhotoURL:  u.PhotoURL,
		CreatedAt: u.CreatedAt,
		DeletedAt: u.DeletedAt,
	}
}

func fromDomainUsers(users []user.User) []User {
	models := make([]User, len(users))
	for i, u := range users {
		models[i] = *fromDomain(&u)
	}
	return models
}

func NewRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) Create(user *user.User) error {
	err := r.db.Create(fromDomain(user)).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *UserRepository) Delete(id string) error {
	return r.db.Where("id = ?", id).Delete(&user.User{}).Error
}

func (r *UserRepository) GetUsers() ([]user.User, error) {
	var models []User

	err := r.db.
		Where("deleted_at IS NULL").
		Order("created_at DESC").
		Find(&models).Error

	if err != nil {
		return nil, err
	}

	return toDomainUsers(models), nil
}

func (r *UserRepository) GetByID(id string) (*user.User, error) {
	var model User

	err := r.db.Where("id = ? AND deleted_at IS NULL", id).First(&model).Error

	if err != nil {
		return nil, err
	}

	return model.toDomain(), nil
}
func (r *UserRepository) GetUserByUsername(username string) (*user.User, error) {

	var model User

	err := r.db.Where("username = ? AND deleted_at IS NULL", username).First(&model).Error

	if err != nil {
		return nil, err
	}

	return model.toDomain(), nil
}

func (r *UserRepository) Search(username string) ([]user.User, error) {
	var models []User
	err := r.db.
		Where("username ILIKE ?", username+"%").
		Find(&models).Error
	return toDomainUsers(models), err
}

func (r *UserRepository) UpdatePhoto(username, photo_url string) error {
	if photo_url == "" {
		return errors.New("photo_url is empty")
	}

	return r.db.Model(&User{}).Where("username = ?", username).Update("photo_url", photo_url).Error
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.NewString()

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)

	return nil
}

func (u *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}
