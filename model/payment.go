package model

import "time"

type Payment struct {
	ID        string
	Name      string
	Photo     []byte
	IsActive  bool
	CreatedAt time.Time
	DeletedAt *time.Time
}
