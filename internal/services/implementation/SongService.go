package implementation

import (
	"LibSongs/internal/core/dto"
	"LibSongs/internal/core/models"
	"LibSongs/internal/services/errs"
	"LibSongs/internal/services/interfaces"
	"LibSongs/internal/services/interfacesRepo"
	"context"
	"errors"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"strings"
)

type SongService struct {
	songRepo interfacesRepo.ISongRepo
	logger   *logrus.Entry
}

func (ss *SongService) GetByID(ctx context.Context, ID uuid.UUID) (*models.SongModel, error) {
	ss.logger.Info("srv: get song by ID")
	ss.logger.Debugf("srv: song ID: %v", ID)

	song, err := ss.songRepo.GetByID(ctx, ID)
	if err != nil && !errors.Is(err, errs.ErrSongDoesNotExists) {
		ss.logger.Errorf("srv: error getting song: %v", err)
		return nil, err
	}

	if song == nil {
		ss.logger.Warn("srv: song don't exists")
		return nil, errs.ErrSongDoesNotExists
	}

	ss.logger.Info("srv: successfully get song by ID")

	return song, nil
}

func (ss *SongService) GetPageByParams(ctx context.Context, params *dto.SongsPageDTO) ([]*models.SongModel, error) {
	ss.logger.Info("srv: get page by params")
	ss.logger.Debugf("srv: songs params: %v", params)

	songs, err := ss.songRepo.GetPageByParams(ctx, params)
	if err != nil && !errors.Is(err, errs.ErrSongsDoesNotExists) {
		ss.logger.Errorf("srv: error getting songs: %v", err)
		return nil, err
	}

	if songs == nil {
		ss.logger.Warn("srv: songs don't exists")
		return nil, errs.ErrSongsDoesNotExists
	}

	ss.logger.Info("srv: successfully get page by params")

	return songs, nil
}

func (ss *SongService) GetTextSongByVerses(ctx context.Context, songDTO *dto.SongDTO) ([]string, error) {
	ss.logger.Info("srv: get song by verses")
	ss.logger.Debugf("srv: song params: %v", songDTO)

	songText, err := ss.songRepo.GetSongTextByID(ctx, songDTO)
	if err != nil && !errors.Is(err, errs.ErrSongDoesNotExists) {
		ss.logger.Errorf("srv: error getting song: %v", err)
		return nil, err
	}

	if songText == "" {
		ss.logger.Warn("srv: song don't exists")
		return nil, errs.ErrSongsDoesNotExists
	}

	songVerses := strings.Split(songText, "\n\n")

	if songDTO.Offset >= len(songVerses) {
		ss.logger.Warn("srv: offset out of range")
		return []string{}, nil
	}

	end := songDTO.Offset + songDTO.Limit
	if end > len(songVerses) {
		end = len(songVerses)
	}

	result := songVerses[songDTO.Offset:end]

	if len(result) == 0 {
		ss.logger.Warn("srv: verses was not found")
		return []string{}, errs.ErrVersesDoesNotExists
	}
	
	ss.logger.Info("srv: successfully get song by verses")

	return result, nil
}

func (ss *SongService) DeleteByID(ctx context.Context, ID uuid.UUID) error {
	ss.logger.Info("srv: delete song by ID")
	ss.logger.Debugf("srv: song ID: %v", ID)

	err := ss.songRepo.DeleteByID(ctx, ID)
	if err != nil && !errors.Is(err, errs.ErrSongDoesNotExists) {
		ss.logger.Errorf("srv: error delete song: %v", err)
		return err
	}

	if errors.Is(err, errs.ErrSongDoesNotExists) {
		ss.logger.Warn("srv: song don't exists")
		return errs.ErrSongDoesNotExists
	}

	ss.logger.Info("srv: successfully delete song by ID")

	return nil
}

func (ss *SongService) Update(ctx context.Context, song *models.SongModel) error {
	ss.logger.Info("srv: update song")
	ss.logger.Debugf("srv: song: %v", song)

	err := ss.songRepo.Update(ctx, song)
	if err != nil && !errors.Is(err, errs.ErrSongDoesNotExists) {
		ss.logger.Errorf("srv: error update song: %v", err)
		return err
	}

	if errors.Is(err, errs.ErrSongDoesNotExists) {
		ss.logger.Warn("srv: song don't exists")
		return errs.ErrSongsDoesNotExists
	}

	ss.logger.Info("srv: successfully update song")

	return nil
}

func (ss *SongService) Create(ctx context.Context, song *models.SongModel) error {
	ss.logger.Info("srv: create song")
	ss.logger.Debugf("srv: song: %v", song)

	err := ss.songRepo.Create(ctx, song)
	if err != nil && !errors.Is(err, errs.ErrSongAlreadyExists) {
		ss.logger.Errorf("srv: error create song: %v", err)
		return err
	}

	if errors.Is(err, errs.ErrSongAlreadyExists) {
		ss.logger.Warn("srv: song already exists")
		return errs.ErrSongsDoesNotExists
	}

	ss.logger.Info("srv: successfully create song")

	return nil
}

func NewSongService(songRepo interfacesRepo.ISongRepo, logger *logrus.Entry) interfaces.ISongService {
	return &SongService{songRepo: songRepo, logger: logger}
}
