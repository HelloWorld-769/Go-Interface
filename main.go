package main

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type User struct {
	ID       uint   `gorm:"primaryKey"`
	Name     string `gorm:"size:255;not null"`
	Email    string `gorm:"size:100;not null;unique"`
	Password string `gorm:"size:100;not null"`
	gorm.Model
}
type UserRepository interface {
	Create(user *User) error
	Read(username string) (*User, error)
	Update(user *User) error
	Delete(username string) error
}

type UserRepositoryImpl struct {
	db *gorm.DB
}

func NewUserRepositoryImpl(db *gorm.DB) UserRepository {
	return &UserRepositoryImpl{db: db}
}

func (r *UserRepositoryImpl) Create(user *User) error {
	return r.db.Create(user).Error
}

func (r *UserRepositoryImpl) Read(username string) (*User, error) {
	var user User
	err := r.db.Where("name = ?", username).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepositoryImpl) Update(user *User) error {
	return r.db.Save(user).Error
}

func (r *UserRepositoryImpl) Delete(username string) error {
	var user User
	err := r.db.Where("name = ?", username).Delete(&user).Error
	return err
}

type Session struct {
	ID        uint   `gorm:"primaryKey"`
	UserID    uint   `gorm:"not null"`
	Token     string `gorm:"size:255;not null;unique"`
	ExpiresAt int64  `gorm:"not null"`
}

type SessionRepository interface {
	Create(session *Session) error
	Read(token string) (*Session, error)
	Update(session *Session) error
	Delete(token string) error
}

type SessionRepositoryImpl struct {
	db *gorm.DB
}

func NewSessionRepositoryImpl(db *gorm.DB) SessionRepository {
	return &SessionRepositoryImpl{db: db}
}

func (r *SessionRepositoryImpl) Create(session *Session) error {
	return r.db.Create(session).Error
}

func (r *SessionRepositoryImpl) Read(token string) (*Session, error) {
	var session Session
	err := r.db.Where("token = ?", token).First(&session).Error
	if err != nil {
		return nil, err
	}
	return &session, nil
}

func (r *SessionRepositoryImpl) Update(session *Session) error {
	return r.db.Save(session).Error
}

func (r *SessionRepositoryImpl) Delete(token string) error {
	var session Session
	err := r.db.Where("token = ?", token).Delete(&session).Error
	return err
}

func main() {
	dsn := "host=localhost user=postgres password=12345678 dbname=testing port=5433 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database: " + err.Error())
	}

	// Migrate the schema
	db.AutoMigrate(&User{}, &Session{})

	// Initialize the UserRepository
	userRepo := NewUserRepositoryImpl(db)

	// Create a new user
	newUser := &User{Name: "John Doe", Email: "john.doe@example.com", Password: "secret"}
	err = userRepo.Create(newUser)
	if err != nil {
		panic("failed to create user: " + err.Error())
	}

	// // Read the user
	readUser, err := userRepo.Read("John Doe")
	if err != nil {
		panic("failed to read user: " + err.Error())
	}

	fmt.Println(readUser)

	// // Update the user
	readUser.Name = "Jane Doe"
	err = userRepo.Update(readUser)
	if err != nil {
		panic("failed to update user: " + err.Error())
	}

	// // Delete the user
	err = userRepo.Delete("Jane Doe")
	if err != nil {
		panic("failed to delete user: " + err.Error())
	}
}
