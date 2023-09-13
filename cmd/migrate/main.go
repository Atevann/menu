package main

import (
	"flag"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/pressly/goose/v3"
	"log"
	_ "menu/migrations"
	"menu/src/di"
	"os"
)

var (
	flags   = flag.NewFlagSet("goose", flag.ExitOnError)
	dir     = flags.String("dir", "/usr/src/menu/migrations", "директория с миграциями")
	command = flags.String("command", "version", "команда")
)

func main() {
	flags.Parse(os.Args[1:])
	args := flags.Args()

	container := di.NewDi()
	dbconfig := container.Config.Database
	dbstring := fmt.Sprintf(
		"%s:%s@tcp(%s)/%s?parseTime=true",
		dbconfig.Username,
		dbconfig.Password,
		dbconfig.Hostname,
		dbconfig.Name)

	db, err := goose.OpenDBWithDriver("mysql", dbstring)
	if err != nil {
		log.Fatalf("goose: ошибка подключения к БД: %v\n", err)
	}

	defer func() {
		if err := db.Close(); err != nil {
			log.Fatalf("goose: ошибка закрытия подключения к БД: %v\n", err)
		}
	}()

	if err := goose.Run(*command, db, *dir, args...); err != nil {
		log.Fatalf("goose %v: %v", *command, err)
	}
}
