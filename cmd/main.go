package main

import (
	"context"
	"errors"
	"github.com/IliaSobolev/Torrseal/internal/bot"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"gopkg.in/telebot.v3"
	"gopkg.in/telebot.v3/layout"
	"os"
	"strconv"
)

func main() {
	//Log setup
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	pretty, err := strconv.ParseBool(os.Getenv("LOG_PRETTY"))
	if err != nil {
		pretty = false
	}
	if pretty {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	}
	switch os.Getenv("LOG_LEVEL") {
	case "DISABLED":
		zerolog.SetGlobalLevel(zerolog.Disabled)
	case "PANIC":
		zerolog.SetGlobalLevel(zerolog.PanicLevel)
	case "FATAL":
		zerolog.SetGlobalLevel(zerolog.FatalLevel)
	case "ERROR":
		zerolog.SetGlobalLevel(zerolog.ErrorLevel)
	case "WARN":
		zerolog.SetGlobalLevel(zerolog.WarnLevel)
	case "DEBUG":
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	case "TRACE":
		zerolog.SetGlobalLevel(zerolog.TraceLevel)
	default:
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	}

	//DB connection
	pool, err := pgxpool.New(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal().Err(err).Msg("failed to connect to database")
	}
	//ping
	err = pool.Ping(context.Background())
	if err != nil {
		log.Fatal().Err(err).Msg("failed to ping database")
	}

	//Migrations
	m, err := migrate.New("file://migrations", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal().Err(err).Msg("failed to initialize to database")
	}
	err = m.Up()
	if err != nil && !errors.Is(err, migrate.ErrNoChange) {
		log.Fatal().Err(err).Msg("failed to apply migrations")
	}

	//Bot
	l, err := layout.New("assets/layout/layout.yml")
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to load layout")
	}
	tb, err := telebot.NewBot(l.Settings())
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to create telebot")
	}
	b := bot.NewBot(tb, l)

	// Start bot
	log.Info().Msg("Starting bot")
	b.StartBlocking()
}
