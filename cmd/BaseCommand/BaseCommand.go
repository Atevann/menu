package BaseCommand

import "menu/src/di"

// BaseCommand Абстрактная комманда
type BaseCommand struct {
	container *di.Di
}

// New Конструктор
func New(container *di.Di) BaseCommand {
	return BaseCommand{container: container}
}

// GetDi Возвращает контейнер DI
func (command BaseCommand) GetDi() *di.Di {
	return command.container
}
