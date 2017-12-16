package command

import (
	"fmt"

	"github.com/mitchellh/cli"
	"github.com/nerdalize/moulin/client"
	pb "github.com/nerdalize/moulin/protobuf"
)

// CreateTask is for creating a single task
type CreateTask struct {
	Ui cli.Ui
}

// Run (LoadCommand) executes the actual action
func (c *CreateTask) Run(args []string) int {
	c.Ui.Output("Creating a single task")

	grpcDriver := client.NewGRPCDriver()
	defer grpcDriver.Connection.Close()

	if len(args) > 2 {
		fmt.Println("received too many arguments")
		return -1
	} else if len(args) < 2 {
		fmt.Println("received too few arguments")
		return -1
	}

	task := new(pb.Task)
	task.QueueID = args[0]
	task.Body = args[1]
	status := grpcDriver.PushTask(task)

	fmt.Printf("submit task %v\n", status)
	return 0
}

// Help (LoadCommand) shows help
func (c *CreateTask) Help() string {
	return "Run as an agent (detailed help information here)"
}

// Synopsis is the short description
func (c *CreateTask) Synopsis() string {
	return "Create a task with body"
}
