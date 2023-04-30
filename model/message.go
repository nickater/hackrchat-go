package model

import (
	"time"

	"github.com/google/uuid"
)

type Message struct {
	Id        int
	ChatId    uuid.UUID
	Content   string
	CreatedAt time.Time
	UpdatedAt time.Time
	SenderId  uuid.UUID
	Messages  []Message
}
