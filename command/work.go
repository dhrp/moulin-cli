package command

import (
	"fmt"

	"github.com/mitchellh/cli"
	"github.com/nerdalize/moulin/client"
	"github.com/nerdalize/moulincli/process"
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

	if len(args) > 1 {
		fmt.Println("received too many arguments for queue")
		return -1
	} else if len(args) < 1 {
		fmt.Println("received too few arguments for queue")
		return -1
	}

	task := grpcDriver.LoadTask(args[0])
	fmt.Printf("  received taskID %s from queue\n", task.TaskID)
	fmt.Printf("  %s\n", task.Body)

	result, err := process.Exec(task)
	if err != nil {
		return result
	}

	fmt.Println("  Task done. Marking as complete.")
	status := grpcDriver.Complete(args[0], task.TaskID)
	fmt.Println(status)

	return 0
}

// Help (LoadCommand) shows help
func (w *Work) Help() string {
	return "Work is for loading, executing, heartbeating and completing tasks from a queue"
}

// Synopsis is the short description
func (w *Work) Synopsis() string {
	return "Work off items from queue"
}
