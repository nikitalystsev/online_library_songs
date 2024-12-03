package handlers

import (
	jsondto "LibSongs/internal/api/core/dto"
	jsonmodels "LibSongs/internal/api/core/models"
	"LibSongs/internal/core/dto"
	"LibSongs/internal/core/models"
	"LibSongs/internal/services/errs"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"io"
	"net/http"
	"os"
	"strconv"
	"time"
)

// @Summary Метод получения данных библиотеки с фильтрацией по всем полям и пагинацией
// @Tags songs
// @ID getPageSongs
// @Accept json
// @Produce json
// @Param group query string false "Название группы"
// @Param song query string false "Название песни"
// @Param release_date query string false "Дата выхода песни"
// @Param text query string false "Текст книги"
// @Param link query string false "Ссылка на песню"
// @Param limit query string false "Лимит"
// @Param offset query string false "Смещение"
// @Success 200 {array} models.JSONSongModel "Список песен"
// @Failure 400 {object} dto.ErrorResponse "Неверный запрос"
// @Failure 404 {object} dto.ErrorResponse "Песни не найдены"
// @Failure 500 {object} dto.ErrorResponse "Внутренняя ошибка сервера"
// @Router /api/v1/songs [get]
func (h *Handler) getPageSongs(ctx *gin.Context) {
	fmt.Println("call getPageSongs")

	var (
		params dto.SongsPageDTO
		err    error
	)

	params.Group = ctx.Query("group")
	params.Song = ctx.Query("song")
	params.Text = ctx.Query("text")
	params.Link = ctx.Query("link")
	tmp := ctx.Query("release_date")
	if tmp != "" {
		var temp time.Time
		temp, err = time.Parse("02.01.2006", ctx.Query("release_date"))
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, jsondto.ErrorResponse{ErrorMsg: err.Error()})
			return
		}
		params.ReleaseDate = &temp
	}

	if h.isNoEmptyField(ctx.Query("limit")) {
		if params.Limit, err = strconv.Atoi(ctx.Query("limit")); err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, jsondto.ErrorResponse{ErrorMsg: err.Error()})
			return
		}
	}

	if h.isNoEmptyField(ctx.Query("offset")) {
		if params.Offset, err = strconv.Atoi(ctx.Query("offset")); err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, jsondto.ErrorResponse{ErrorMsg: err.Error()})
			return
		}
	}

	songs, err := h.songService.GetPageByParams(ctx.Request.Context(), &params)
	if err != nil && errors.Is(err, errs.ErrSongsDoesNotExists) {
		ctx.AbortWithStatusJSON(http.StatusNotFound, jsondto.ErrorResponse{ErrorMsg: err.Error()})
		return
	}
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, jsondto.ErrorResponse{ErrorMsg: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, h.convertArrayBooksToJSONSongModels(songs))
}

// @Summary Метод получения данных библиотеки с фильтрацией по всем полям и пагинацией
// @Tags songs
// @ID getSongByVerses
// @Accept json
// @Produce json
// @Param id path string true "Идентификатор песни"
// @Param limit query string false "Лимит"
// @Param offset query string false "Смещение"
// @Success 200 {array}  string "Куплеты"
// @Failure 400 {object} dto.ErrorResponse "Неверный запрос"
// @Failure 404 {object} dto.ErrorResponse "Песня не найдена"
// @Failure 500 {object} dto.ErrorResponse "Внутренняя ошибка сервера"
// @Router /api/v1/songs/{id}/verses [get]
func (h *Handler) getSongByVerses(ctx *gin.Context) {
	fmt.Println("call getSongByVerses")

	var (
		songDTO dto.SongDTO
		err     error
	)

	if songDTO.ID, err = uuid.Parse(ctx.Param("id")); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, jsondto.ErrorResponse{ErrorMsg: err.Error()})
		return
	}

	if h.isNoEmptyField(ctx.Query("limit")) {
		if songDTO.Limit, err = strconv.Atoi(ctx.Query("limit")); err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, jsondto.ErrorResponse{ErrorMsg: err.Error()})
			return
		}
	}

	if h.isNoEmptyField(ctx.Query("offset")) {
		if songDTO.Offset, err = strconv.Atoi(ctx.Query("offset")); err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, jsondto.ErrorResponse{ErrorMsg: err.Error()})
			return
		}
	}

	verses, err := h.songService.GetTextSongByVerses(ctx.Request.Context(), &songDTO)
	if err != nil && (errors.Is(err, errs.ErrSongDoesNotExists) || errors.Is(err, errs.ErrVersesDoesNotExists)) {
		ctx.AbortWithStatusJSON(http.StatusNotFound, jsondto.ErrorResponse{ErrorMsg: err.Error()})
		return
	}

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, jsondto.ErrorResponse{ErrorMsg: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, verses)
}

// @Summary Метод удаления песни
// @Tags songs
// @ID deleteByID
// @Accept json
// @Produce json
// @Param id path string true "Идентификатор песни"
// @Success 200 "Успешное удаление песни"
// @Failure 400 {object} dto.ErrorResponse "Неверный запрос"
// @Failure 404 {object} dto.ErrorResponse "Песни не найдены"
// @Failure 500 {object} dto.ErrorResponse "Внутренняя ошибка сервера"
// @Router /api/v1/songs/{id} [delete]
func (h *Handler) deleteByID(ctx *gin.Context) {
	fmt.Println("call DeleteByID")

	ID, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, jsondto.ErrorResponse{ErrorMsg: err.Error()})
		return
	}

	err = h.songService.DeleteByID(ctx.Request.Context(), ID)
	if err != nil && errors.Is(err, errs.ErrSongDoesNotExists) {
		ctx.AbortWithStatusJSON(http.StatusNotFound, jsondto.ErrorResponse{ErrorMsg: err.Error()})
		return
	}
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, jsondto.ErrorResponse{ErrorMsg: err.Error()})
		return
	}

	ctx.Status(http.StatusOK)
}

// @Summary Метод изменения данных песни
// @Tags songs
// @ID updateByID
// @Accept json
// @Produce json
// @Param id path string true "Идентификатор песни"
// @Param newSongParams body  dto.SongParamDTO true "Параметры песни"
// @Success 200 "Успешное обновление данных песни"
// @Failure 400 {object} dto.ErrorResponse "Неверный запрос"
// @Failure 404 {object} dto.ErrorResponse "Песня не найдена или куплеты не найдены"
// @Failure 500 {object} dto.ErrorResponse "Внутренняя ошибка сервера"
// @Router /api/v1/songs/{id} [put]
func (h *Handler) updateByID(ctx *gin.Context) {
	var inp dto.SongParamDTO
	if err := ctx.BindJSON(&inp); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, jsondto.ErrorResponse{ErrorMsg: err.Error()})
		return
	}

	ID, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, jsondto.ErrorResponse{ErrorMsg: err.Error()})
		return
	}

	song, err := h.songService.GetByID(ctx.Request.Context(), ID)
	if err != nil && errors.Is(err, errs.ErrSongsDoesNotExists) {
		ctx.AbortWithStatusJSON(http.StatusNotFound, jsondto.ErrorResponse{ErrorMsg: err.Error()})
	}
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, jsondto.ErrorResponse{ErrorMsg: err.Error()})
	}

	releaseDate, err := time.Parse("02.01.2006", inp.ReleaseDate)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, jsondto.ErrorResponse{ErrorMsg: err.Error()})
		return
	}

	song.Group = inp.Group
	song.Song = inp.Song
	song.ReleaseDate = releaseDate
	song.Text = inp.Text
	song.Link = inp.Link

	err = h.songService.Update(ctx.Request.Context(), song)
	if err != nil && errors.Is(err, errs.ErrSongsDoesNotExists) {
		ctx.AbortWithStatusJSON(http.StatusNotFound, jsondto.ErrorResponse{ErrorMsg: err.Error()})
		return
	}
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, jsondto.ErrorResponse{ErrorMsg: err.Error()})
		return
	}

	ctx.Status(http.StatusOK)
}

// @Summary Метод добавления новой песни
// @Tags songs
// @ID addNewSong
// @Accept json
// @Produce json
// @Param nemSong body  dto.NewSongDTO true "Группа и название песни"
// @Success 201 "Успешное создание песни"
// @Failure 400 {object} dto.ErrorResponse "Неверный запрос"
// @Failure 409 {object} dto.ErrorResponse "Песня уже существует"
// @Failure 404 {object} dto.ErrorResponse "Песня не найдена"
// @Failure 500 {object} dto.ErrorResponse "Внутренняя ошибка сервера"
// @Router /api/v1/songs [post]
func (h *Handler) addNewSong(ctx *gin.Context) {
	var inp jsondto.NewSongDTO
	if err := ctx.BindJSON(&inp); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, jsondto.ErrorResponse{ErrorMsg: err.Error()})
		return
	}

	otherApiHost := os.Getenv("OTHER_API_HOST")
	otherApiPort := os.Getenv("OTHER_API_PORT")

	url := fmt.Sprintf("http://%s:%s/info?group=%s&song=%s", otherApiHost, otherApiPort, inp.Group, inp.Song)

	response, err := http.Get(url)
	if err != nil {
		ctx.AbortWithStatusJSON(response.StatusCode, jsondto.ErrorResponse{ErrorMsg: err.Error()})
		return
	}

	defer func(Body io.ReadCloser) {
		if err = Body.Close(); err != nil {
			return
		}

	}(response.Body)

	body, err := io.ReadAll(response.Body)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, jsondto.ErrorResponse{ErrorMsg: err.Error()})
		return
	}
	var songInfo struct {
		ReleaseDate time.Time `json:"releaseDate"`
		Text        string    `json:"text"`
		Link        string    `json:"link"`
	}

	if err = json.Unmarshal(body, &songInfo); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, jsondto.ErrorResponse{ErrorMsg: err.Error()})
		return
	}

	song := &models.SongModel{
		ID:          uuid.New(),
		Group:       inp.Group,
		Song:        inp.Song,
		ReleaseDate: songInfo.ReleaseDate,
		Text:        songInfo.Text,
		Link:        songInfo.Link,
	}

	err = h.songService.Create(ctx.Request.Context(), song)
	if err != nil && errors.Is(err, errs.ErrSongAlreadyExists) {
		ctx.AbortWithStatusJSON(http.StatusConflict, jsondto.ErrorResponse{ErrorMsg: err.Error()})
		return
	}

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, jsondto.ErrorResponse{ErrorMsg: err.Error()})
		return
	}

	ctx.Status(http.StatusCreated)
}

func (h *Handler) isNoEmptyField(field string) bool {
	return field != "" && field != "NaN" && field != "null"
}

func (h *Handler) convertArrayBooksToJSONSongModels(songs []*models.SongModel) []*jsonmodels.JSONSongModel {
	jsonBooks := make([]*jsonmodels.JSONSongModel, len(songs))
	for i, book := range songs {
		jsonBooks[i] = h.convertToJSONSongModel(book)
	}

	return jsonBooks
}

func (h *Handler) convertToJSONSongModel(song *models.SongModel) *jsonmodels.JSONSongModel {
	return &jsonmodels.JSONSongModel{
		ID:          song.ID,
		Group:       song.Group,
		Song:        song.Song,
		ReleaseDate: song.ReleaseDate,
		Text:        song.Text,
		Link:        song.Link,
	}
}
