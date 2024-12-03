package errs

import "errors"

var (
	ErrSongsDoesNotExists  = errors.New("songs does not exist")
	ErrSongDoesNotExists   = errors.New("song does not exist")
	ErrSongAlreadyExists   = errors.New("song already exists")
	ErrVersesDoesNotExists = errors.New("verses does not exists")
)
