package main

import (
	"errors"
	"fmt"

	"github.com/golang-migrate/migrate/v4"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/olegtemek/currency-task/internal/config"
)

func main() {

	cfg := config.New()

	if cfg.DbUrl == "" {
		panic("cannot get db url")
	}

	m, err := migrate.New("file://migrations", fmt.Sprintf("%s?sslmode=disable", cfg.DbUrl))
	if err != nil {
		panic(err)
	}

	if err := m.Up(); err != nil {
		if errors.Is(err, migrate.ErrNoChange) {
			fmt.Println("no migrations to apply")

			return
		}

		panic(err)
	}

	fmt.Println("migrations applied")

}
