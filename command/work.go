package command

import (
	"fmt"

	"github.com/dhrp/moulin/client"
	"github.com/dhrp/moulincli/process"
	"github.com/mitchellh/cli"
)

// Work is for loading, executing, heartbeating and completing tasks
type Work struct {
	Ui cli.Ui
}

// Run (LoadCommand) executes the actual action
func (w *Work) Run(args []string) int {
	w.Ui.Output("Workin' from queue " + args[0])

	grpcDriver := client.NewGRPCDriver()
	defer grpcDriver.Connection.Close()

	if len(args) > 2 {
		fmt.Println("received too many arguments for queue")
		return -1
	} else if len(args) < 1 {
		fmt.Println("received too few arguments for queue")
		return -1
	}

	workType := "once"
	if len(args) == 2 {
		workType = args[1]
	}

	switch workType {
	case "once":
		process.Work(grpcDriver, args[0], "once")
		return 0
	case "until-finished":
		fmt.Println("not implemented yet ")
		return 1
	case "forever":
		process.Work(grpcDriver, args[0], "forever")
		return 1
	}
	fmt.Println("invalid work type")
	return 1
}

// Help (LoadCommand) shows help
func (w *Work) Help() string {
	return "Work is for loading, executing, heartbeating and completing tasks from a queue"
}

// Synopsis is the short description
func (w *Work) Synopsis() string {
	return "Work off items from queue"
}
