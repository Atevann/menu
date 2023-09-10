package config

// Config структура конфигурации
type Config struct {
	HttpServer HttpServer
	Database   Database
}

type HttpServer struct {
	ListenPort int
}

type Database struct {
	Hostname string
	Name     string
	Username string
	Password string
}
