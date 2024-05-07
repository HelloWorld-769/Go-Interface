package model

import "gorm.io/gorm"

type Session struct {
	gorm.Model
	UserID    uint
	SessionID string
}

type SessionRepository interface {
	Create(session *Session) error
	Update(session *Session) error
	Delete(id uint) error
	FindByID(id uint) (*Session, error)
	FindAll() ([]*Session, error)
}
