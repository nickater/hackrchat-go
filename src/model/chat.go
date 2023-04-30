package model

import (
	"time"

	"github.com/google/uuid"
)

type Chat struct {
	Id        uuid.UUID
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
	Users     []User
	Messages  []Message
}
