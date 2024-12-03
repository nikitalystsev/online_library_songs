package dto

import (
	"github.com/google/uuid"
	"time"
)

type SongsPageDTO struct {
	Group       string
	Song        string
	ReleaseDate *time.Time
	Text        string
	Link        string
	Limit       int
	Offset      int
}

type SongDTO struct {
	ID     uuid.UUID
	Limit  int
	Offset int
}

type SongParamDTO struct {
	Group       string
	Song        string
	ReleaseDate string
	Text        string
	Link        string
}
