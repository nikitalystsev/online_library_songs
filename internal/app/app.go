package app

import (
	"LibSongs/internal/api/handlers"
	"LibSongs/internal/config"
	"LibSongs/internal/repositories"
	"LibSongs/internal/repositories/implementation"
	repoimplementation "LibSongs/internal/services/implementation"
	"LibSongs/pkg/logging"
	"fmt"
	_ "github.com/lib/pq"
)

func Run(configDir string) {
	cfg, err := config.Init(configDir)
	if err != nil {
		panic(err)
	}

	fmt.Printf("postgres db name: %s\n", cfg.Postgres.DBName)
	fmt.Printf("postgres db user: %s\n", cfg.Postgres.Username)
	fmt.Printf("postgres db password: %s\n", cfg.Postgres.Password)
	fmt.Printf("postgres host: %s\n", cfg.Postgres.Host)
	fmt.Printf("postgres ssl mode: %s\n", cfg.Postgres.SSLMode)

	logger, err := logging.NewLogger()
	if err != nil {
		panic(err)
	}

	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.Postgres.Host, cfg.Postgres.Port, cfg.Postgres.Username, cfg.Postgres.DBName,
		cfg.Postgres.Password, cfg.Postgres.SSLMode)

	fmt.Printf("dsn: %s\n", dsn)
	db, err := repoPostgres.NewClient(dsn)
	if err != nil {
		logger.Errorf("error connect to postgres: %v", err)
		return
	}

	songRepo := implementation.NewSongRepo(db, logger)

	songService := repoimplementation.NewSongService(songRepo, logger)

	handler := handlers.NewHandler(
		songService,
	)

	router := handler.InitRoutes()

	err = router.Run(":" + cfg.Port)
	if err != nil {
		logger.Errorf("error running server: %v", err)
		return
	}
}
