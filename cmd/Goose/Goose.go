package Goose

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/pressly/goose/v3"
	"log"
	"menu/cmd/BaseCommand"
)

// Goose Инструмент для миграций https://github.com/pressly/goose
type Goose struct {
	BaseCommand.BaseCommand
}

// migrationsLocation Путь для сохранения миграций в контейнере
var migrationsLocation = "/usr/src/menu/migrations"

func (g Goose) GetDescription() string {
	return "Инструмент для управления миграциями"
}

func (g Goose) GetHelp() string {
	return "https://github.com/pressly/goose"
}

func (g Goose) IsArgsRequired() bool {
	return true
}

func (g Goose) Run(args []string) {
	command := args[0]
	gooseArgs := args[1:]
	dbconfig := g.GetDi().Config.Database
	dbstring := fmt.Sprintf(
		"%s:%s@tcp(%s)/%s?parseTime=true",
		dbconfig.Username,
		dbconfig.Password,
		dbconfig.Hostname,
		dbconfig.Name)

	db, err := goose.OpenDBWithDriver("mysql", dbstring)
	if err != nil {
		log.Fatalf("Goose: ошибка подключения к БД: %v\n", err)
	}

	defer func() {
		if err := db.Close(); err != nil {
			log.Fatalf("Goose: ошибка закрытия подключения к БД: %v\n", err)
		}
	}()

	if err := goose.Run(command, db, migrationsLocation, gooseArgs...); err != nil {
		log.Fatalf("Goose %v: %v", command, err)
	}
}
