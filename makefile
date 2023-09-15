# Документация https://github.com/pressly/goose
# Примеры:
#     Создание миграции:   make migration GOOSE_CMD="create migration_name go"
#     Применение миграций: make migration GOOSE_CMD="up"
# 	  Откат миграции:      make migration GOOSE_CMD="down"
migration:
	docker exec -it menu-cli go run cmd/main.go -cmd=goose $(GOOSE_CMD)
