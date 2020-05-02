package command

import "os"

type ExitCommand struct {}

func (e *ExitCommand) Name() string {
    return "exit"
}

func (e *ExitCommand) Description() string {
    return "Exiting from application"
}

func (e *ExitCommand) Action() {
    os.Exit(0)
}
