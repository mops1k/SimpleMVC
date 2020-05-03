package command

import (
    "os"

    "SimpleMVC/app/service/command"
)

type ExitCommand struct {}

func (e *ExitCommand) Name() string {
    return "exit"
}

func (e *ExitCommand) Description() string {
    return "Exiting from application"
}

func (e *ExitCommand) Action(ctx command.Context) {
    os.Exit(0)
}
