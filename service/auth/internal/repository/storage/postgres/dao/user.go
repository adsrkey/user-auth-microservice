package dao

import (
	"github.com/google/uuid"
	"time"
)

type User struct {
	ID         uuid.UUID `db:"id"`
	CreatedAt  time.Time `db:"created_at"`
	ModifiedAt time.Time `db:"modified_at"`
	Email      string    `db:"email"`
	Hash       []byte    `db:"hash"`
}
