package main

import (
	"log"
	"os"

	"github.com/Kashaan-Ekhlas/Key-Bored-Party/backend/internal/config"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	migration, err := migrate.New("file://migrations", cfg.DatabaseURL())
	if err != nil {
		log.Fatal(err)
	}

	if len(os.Args) != 2 {
		log.Fatal("usage: migrate up/down")
	}

	switch os.Args[1] {

	case "up":
		if err := migration.Up(); err != nil && err != migrate.ErrNoChange {
			log.Fatal(err)
		}

	case "down":
		if err := migration.Down(); err != nil && err != migrate.ErrNoChange {
			log.Fatal(err)
		}

	default:
		log.Fatal("usage: migrate up/down")

	}
}
