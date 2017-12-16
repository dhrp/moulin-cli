package command

import (
	"fmt"

	"github.com/mitchellh/cli"
	"github.com/nerdalize/moulin/client"
	"github.com/nerdalize/moulincli/process"
)

// LoadCommand is for loading
type LoadCommand struct {
	Ui cli.Ui
}

// Run (LoadCommand) executes the actual action
func (c *LoadCommand) Run(args []string) int {
	c.Ui.Output("Loading item from queue")

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
	fmt.Printf("received taskID %s from queue\n", task.TaskID)
	fmt.Printf("%s\n", task.Body)

	result, err := process.Exec(task)
	if err != nil {
		return result
	}

	return 0
}

// Help (LoadCommand) shows help
func (c *LoadCommand) Help() string {
	return "Run as an agent (detailed help information here)"
}

// Synopsis is the short description
func (c *LoadCommand) Synopsis() string {
	return "Run as an agent"
}
