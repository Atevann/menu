package di

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"menu/config"
)

// Di DI контейнер
type Di struct {
	Config   *config.Config
	Database *gorm.DB
}

// NewDi Инициализация DI
func NewDi() Di {
	cfg := newConfig()
	db := newDatabase(cfg.Database)

	return Di{
		Config:   cfg,
		Database: db,
	}
}

// newConfig Чтение конфига
func newConfig() *config.Config {
	cfg := &config.Config{}
	_, err := toml.DecodeFile("config/config.toml", cfg)
	if err != nil {
		log.Fatalf("Ошибка декодирования файла конфигов %v", err)
	}

	return cfg
}

// newDatabase Подключение к базе данных
func newDatabase(database config.Database) *gorm.DB {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		database.Username,
		database.Password,
		database.Hostname,
		database.Name)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Ошибка подключения к базе данных %v", err)
	}

	return db
}
