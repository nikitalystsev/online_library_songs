package implementation

import (
	"LibSongs/internal/core/dto"
	"LibSongs/internal/core/models"
	repomodels "LibSongs/internal/repositories/core/models"
	"LibSongs/internal/services/errs"
	"LibSongs/internal/services/interfacesRepo"
	"context"
	"database/sql"
	"errors"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type SongRepo struct {
	db     *sqlx.DB
	logger *logrus.Entry
}

func (ss *SongRepo) GetSongTextByID(ctx context.Context, songDTO *dto.SongDTO) (string, error) {
	ss.logger.Infof("repo: get song text")
	ss.logger.Debugf("repo: get song params: %v", songDTO)

	query := `select _text from ls.song where id = $1`

	var song repomodels.SongModel

	err := ss.db.GetContext(ctx, &song, query, songDTO.ID)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		ss.logger.Errorf("repo: error selecting song text with ID: %v", err)
		return "", err
	}

	if errors.Is(err, sql.ErrNoRows) {
		ss.logger.Warnf("repo: song text was not found")
		return "", errs.ErrSongDoesNotExists
	}

	ss.logger.Infof("selected song text")

	return song.Text, nil
}

func (ss *SongRepo) DeleteByID(ctx context.Context, ID uuid.UUID) error {
	ss.logger.Infof("repo: delete song with ID: %s", ID)

	query := `delete from ls.song where id = $1`

	result, err := ss.db.ExecContext(ctx, query, ID)
	if err != nil {
		ss.logger.Errorf("repo: error deleting song: %v", err)
		return err
	}
	rows, err := result.RowsAffected()
	if err != nil {
		ss.logger.Errorf("repo: error deleting song: %v", err)
		return err
	}
	if rows != 1 {
		ss.logger.Errorf("repo: error deleting song: expected 1 row affected, got %d", rows)
		return errors.New("songRepo.Delete: expected 1 row affected")
	}

	ss.logger.Infof("repo: deleted song with ID: %s", ID)

	return nil
}

func (ss *SongRepo) Update(ctx context.Context, song *models.SongModel) error {
	ss.logger.Info("repo: update song")
	ss.logger.Debugf("repo: song: %v", song)

	query := `update ls.song 
			  set _group = $1,
			      _song = $2,
			      _release_date = $3,
			      _text = $4,
			      _link = $5
			  where id = $6`

	result, err := ss.db.ExecContext(
		ctx, query,
		song.Group,
		song.Song,
		song.ReleaseDate,
		song.Text,
		song.Link,
		song.ID,
	)
	if err != nil {
		ss.logger.Errorf("error updating song: %v", err)
		return err
	}
	rows, err := result.RowsAffected()
	if err != nil {
		ss.logger.Errorf("error updating song: %v", err)
		return err
	}
	if rows != 1 {
		ss.logger.Errorf("error updating song: expected 1 row affected, got %d", rows)
		return errors.New("bookRepo.Update: expected 1 row affected")
	}

	ss.logger.Infof("updated song with ID: %s", song.ID)

	return nil
}

func (ss *SongRepo) Create(ctx context.Context, song *models.SongModel) error {
	ss.logger.Info("repo: insert song")
	ss.logger.Debugf("repo: song: %v", song)

	query := `insert into ls.song values ($1, $2, $3, $4, $5, $6)`

	result, err := ss.db.ExecContext(
		ctx, query,
		song.ID,
		song.Group,
		song.Song,
		song.ReleaseDate,
		song.Text,
		song.Link,
	)
	if err != nil {
		ss.logger.Errorf("repo: error insert song: %v", err)
		return err
	}
	rows, err := result.RowsAffected()
	if err != nil {
		ss.logger.Errorf("repo: error inserting book: %v", err)
		return err
	}
	if rows != 1 {
		ss.logger.Errorf("repo: error inserting book: expected 1 row affected, got %d", rows)
		return errors.New("songRepo.Create: expected 1 row affected")
	}

	ss.logger.Infof("repo: inserted book with ID: %s", song.ID)

	return nil
}

func (ss *SongRepo) GetByID(ctx context.Context, ID uuid.UUID) (*models.SongModel, error) {
	ss.logger.Info("repo: get song by ID")
	ss.logger.Debug("repo: song ID", ID)

	query := `select id, _group, _song, _release_date, _text, _link from ls.song where id = $1`

	var song repomodels.SongModel

	err := ss.db.GetContext(ctx, &song, query, ID)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		ss.logger.Errorf("error selecting song with ID: %v", err)
		return nil, err
	}

	if errors.Is(err, sql.ErrNoRows) {
		ss.logger.Warnf("song with this ID not found %s", ID)
		return nil, errs.ErrSongDoesNotExists
	}

	ss.logger.Infof("selected song with ID: %s", ID)

	return ss.convertToSongModel(&song), nil
}

func (ss *SongRepo) GetPageByParams(ctx context.Context, params *dto.SongsPageDTO) ([]*models.SongModel, error) {
	ss.logger.Info("repo: get page by params")
	ss.logger.Debug("repo: params: ", params)

	query := `select * 
	          from ls.song 
	          where ($1 = '' or _group ilike '%' || $1 || '%') and 
	                ($2 = '' or _song ilike '%' || $2 || '%') and 
	                ($3::date is null or _release_date = $3::date) and 
	                ($4 = '' or _text ilike '%' || $4 || '%') and 
	                ($5 = '' or _link ilike '%' || $5 || '%')
	          limit $6 offset $7`

	var coreSongs []*repomodels.SongModel

	err := ss.db.SelectContext(
		ctx, &coreSongs, query,
		params.Group,
		params.Song,
		params.ReleaseDate,
		params.Text,
		params.Link,
		params.Limit,
		params.Offset,
	)

	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		ss.logger.Errorf("repo: error select songs by params: %s", err.Error())
		return nil, err
	}

	if len(coreSongs) == 0 {
		ss.logger.Warn("repo: songs not found")
		return nil, errs.ErrSongsDoesNotExists
	}

	ss.logger.Infof("repo: was found %d songs", len(coreSongs))

	return ss.convertArrayToSongModel(coreSongs), nil
}

func (ss *SongRepo) convertArrayToSongModel(songs []*repomodels.SongModel) []*models.SongModel {
	_songs := make([]*models.SongModel, len(songs))
	for i, song := range songs {
		_songs[i] = ss.convertToSongModel(song)
	}

	return _songs
}

func (ss *SongRepo) convertToSongModel(song *repomodels.SongModel) *models.SongModel {
	return &models.SongModel{
		ID:          song.ID,
		Group:       song.Group,
		Song:        song.Song,
		ReleaseDate: song.ReleaseDate,
		Text:        song.Text,
		Link:        song.Link,
	}
}

func NewSongRepo(db *sqlx.DB, logger *logrus.Entry) interfacesRepo.ISongRepo {
	return &SongRepo{db: db, logger: logger}
}
