# Документация https://github.com/pressly/goose
# Примеры:
#     Создание миграции:   make migration GOOSE_CLI="-command=create migration_name go"
#     Применение миграций: make migration GOOSE_CLI="-command=up"
# 	  Откат миграции:      make migration GOOSE_CLI="-command=down"
migration:
	docker exec -it menu-cli go run cmd/migrate/main.go $(GOOSE_CLI)
