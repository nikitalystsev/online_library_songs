package models

import (
	"github.com/google/uuid"
	"time"
)

type SongModel struct {
	ID          uuid.UUID `db:"id"`
	Group       string    `db:"_group"`
	Song        string    `db:"_song"`
	ReleaseDate time.Time `db:"_release_date"`
	Text        string    `db:"_text"`
	Link        string    `db:"_link"`
}
