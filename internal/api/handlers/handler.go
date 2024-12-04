package handlers

import (
	_ "LibSongs/docs_swagger"
	"LibSongs/internal/services/interfaces"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"io"
	"net/http"
)

type Handler struct {
	songService interfaces.ISongService
}

func NewHandler(
	songService interfaces.ISongService,
) *Handler {
	return &Handler{
		songService: songService,
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	gin.DefaultWriter = io.Discard

	router := gin.Default()

	router.Use(h.corsSettings())

	api := router.Group("/api")
	{
		v1 := api.Group("/v1")
		{
			v1.GET("/songs", h.getPageSongs)
			v1.GET("/songs/:id/verses", h.getSongByVerses)
			v1.DELETE("/songs/:id", h.deleteByID)
			v1.PUT("/songs/:id", h.updateByID)
			v1.POST("/songs", h.addNewSong)
		}
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return router
}

func (h *Handler) corsSettings() gin.HandlerFunc {
	return cors.New(cors.Config{
		AllowMethods: []string{
			http.MethodPost,
			http.MethodGet,
			http.MethodPut,
			http.MethodPatch,
			http.MethodDelete,
		},
		AllowOrigins: []string{
			"*",
		},
		AllowCredentials: true,
		AllowHeaders: []string{
			"Authorization",
			"Content-Type",
		},
		ExposeHeaders: []string{
			"Content-Type",
		},
	})
}
