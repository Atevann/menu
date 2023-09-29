package di

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"menu/config"
	"menu/src/di/repositories"
)

// Di DI контейнер
type Di struct {
	Config       config.Config
	Database     *gorm.DB
	Repositories repositories.Repositories
}

// NewDi Инициализация DI
func NewDi() Di {
	cfg := newConfig()
	db := newDatabase(cfg.Database)
	repos := repositories.InitRepositories(db)

	return Di{
		Config:       cfg,
		Database:     db,
		Repositories: repos,
	}
}

// newConfig Чтение конфига
func newConfig() config.Config {
	cfg := &config.Config{}
	_, err := toml.DecodeFile("config/config.toml", cfg)
	if err != nil {
		log.Fatalf("Ошибка декодирования файла конфигов %v", err)
	}

	return *cfg
}

// newDatabase Подключение к базе данных
func newDatabase(database config.Database) *gorm.DB {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		database.Username,
		database.Password,
		database.Hostname,
		database.Name)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{Logger: logger.Default.LogMode(logger.Info)})
	if err != nil {
		log.Fatalf("Ошибка подключения к базе данных %v", err)
	}

	return db
}
