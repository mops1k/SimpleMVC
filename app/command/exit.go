package command

import (
    "os"

    "SimpleMVC/app/service"
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
    _ = service.Container.GetLogger().App.Info("Bye Bye...")
    os.Exit(0)
}
