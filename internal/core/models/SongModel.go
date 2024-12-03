package models

import (
	"github.com/google/uuid"
	"time"
)

type SongModel struct {
	ID          uuid.UUID
	Group       string
	Song        string
	ReleaseDate time.Time
	Text        string
	Link        string
}
