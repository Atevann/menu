package main

import (
	"flag"
	"fmt"
	"menu/cmd/BaseCommand"
	"menu/cmd/Goose"
	"menu/src/di"
	"os"
)

// CommandInterface Интерфейс команд
type CommandInterface interface {
	// Run Запускает команду
	Run(args []string)
	// IsArgsRequired Необходимы ли команде аргументы
	IsArgsRequired() bool
	// GetDescription Возвращает описание команды
	GetDescription() string
	// GetHelp Возвращает справку по команде
	GetHelp() string
}

var (
	flags = flag.NewFlagSet("cmd", flag.ExitOnError)
	help  = flags.Bool("help", false, "Отобразить справку")
	cmd   = flags.String("cmd", "", "Название команды")
)

func main() {
	flags.Parse(os.Args[1:])

	args := flags.Args()
	container := di.NewDi()
	commands := getCommands(&container)

	if *help == true {
		displayHelp(commands)

		os.Exit(0)
	}

	command := commands[*cmd]

	if command == nil {
		fmt.Printf("Команда не найдена: %s\n", *cmd)

		os.Exit(0)
	}

	if len(args) == 0 && command.IsArgsRequired() {
		fmt.Printf("Недостаточно аргументов\n")

		os.Exit(0)
	}

	command.Run(args)
}

// displayHelp Отображает справку
func displayHelp(commands map[string]CommandInterface) {
	for name, command := range commands {
		fmt.Printf(
			"Справка:\n"+
				"%v\n"+
				"\tDescription: \t%v\n"+
				"\tHelp: \t\t%v\n",
			name,
			command.GetDescription(),
			command.GetHelp())
	}
}

// getCommands Вовзращает список всех зарегистрированных команд
func getCommands(container *di.Di) map[string]CommandInterface {
	commands := make(map[string]CommandInterface)
	baseCommand := BaseCommand.New(container)

	commands["goose"] = Goose.Goose{BaseCommand: baseCommand}

	return commands
}
