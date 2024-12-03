package models

import (
	"github.com/google/uuid"
	"time"
)

type JSONSongModel struct {
	ID          uuid.UUID `json:"id"`
	Group       string    `json:"group"`
	Song        string    `json:"song"`
	ReleaseDate time.Time `json:"release_date"`
	Text        string    `json:"text"`
	Link        string    `json:"link"`
}
