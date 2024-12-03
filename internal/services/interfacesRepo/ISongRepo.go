package interfacesRepo

import (
	"LibSongs/internal/core/dto"
	"LibSongs/internal/core/models"
	"context"
	"github.com/google/uuid"
)

type ISongRepo interface {
	GetPageByParams(ctx context.Context, params *dto.SongsPageDTO) ([]*models.SongModel, error)
	GetSongTextByID(ctx context.Context, songDTO *dto.SongDTO) (string, error)
	DeleteByID(ctx context.Context, ID uuid.UUID) error
	Update(ctx context.Context, song *models.SongModel) error
	Create(ctx context.Context, song *models.SongModel) error
	GetByID(ctx context.Context, ID uuid.UUID) (*models.SongModel, error)
}
