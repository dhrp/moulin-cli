package command

import (
	"fmt"
	"strconv"

	"github.com/dhrp/moulin/client"
	"github.com/mitchellh/cli"
)

// CreateTask is for creating a single task
type Peek struct {
	Ui cli.Ui
}

// Run (LoadCommand) executes the actual action
func (c *Peek) Run(args []string) int {
	c.Ui.Output("Peeking into queue")

	grpcDriver := client.NewGRPCDriver()
	defer grpcDriver.Connection.Close()

	var limit int64 = 10
	var err error

	if len(args) > 3 {
		fmt.Println("received too many arguments")
		return -1
	} else if len(args) < 2 {
		fmt.Println("received too few arguments")
		return -1
	} else if len(args) == 3 {
		limit, err = strconv.ParseInt(args[2], 10, 32)
		if err != nil {
			fmt.Println("third argument should be a number")
			return -1
		}
	}
	queueID := args[0]
	phase := args[1]

	taskList, err := grpcDriver.Peek(queueID, phase, int32(limit))
	if err != nil {
		fmt.Println("an error occured")
		return -1
	}

	tasks := taskList.Tasks
	for i := 0; i < len(tasks); i++ {
		fmt.Println(tasks[i])
	}

	fmt.Printf("items:     %d\n", taskList.TotalItems)

	// fmt.Printf("running:   %d\n", status.NonExpiredCount)
	// fmt.Printf("expired:   %d\n", status.ExpiredCount)
	// fmt.Printf("completed: %d\n", status.CompletedCount)

	return 0
}

// Help (LoadCommand) shows help
func (c *Peek) Help() string {
	return "Peek into a queue, show the next n items"
}

// Synopsis is the short description
func (c *Peek) Synopsis() string {
	return "Get a peek into a queue"
}
