package process

import (
	"fmt"
	"log"
	"time"

	"github.com/dhrp/moulin/client"
	pb "github.com/dhrp/moulin/protobuf"
)

//
// func Close(grpcDriver *client.GRPCDriver) {
//
// 	log.Println("closing connection")
// 	defer grpcDriver.Connection.Close()
// }

func count() {
	seconds := 0
	for {
		time.Sleep(1000 * time.Millisecond)
		log.Print(seconds)
		seconds = seconds + 1
	}
}

// Work manages getting, heartbeating, and completing or failing
// items, in a loop
func Work(grpcDriver *client.GRPCDriver, queueID, workType string) (result int, err error) {

	// defer Close(grpcDriver)

	var exit bool
	if workType == "once" {
		exit = true
	} else if workType == "until-finished" {
		exit = false
	} else if workType == "forever" {
		exit = false
	}

	// forever loop, until exit == true
	// go count()

	for {

		task, err := grpcDriver.LoadTask(queueID)
		if err != nil {
			log.Panic("failed loading task")
		}
		fmt.Printf("  received taskID %s from queue\n", task.TaskID)

		// let the exec function do the hard work
		result, err := Exec(task)
		if err != nil {
			fmt.Printf("  failed to exec %v!!", result)
			// ToDo: mark as failed
		}

		if result != 0 {
			fmt.Printf("  Task failed with code %d ?!? (still marking as complete for now)", result)
		}

		status := grpcDriver.Complete(queueID, task.TaskID)
		if status.Status != pb.Status_SUCCESS {
			log.Panic("failed marking as complete")
		}

		if exit == true {
			return 0, nil
		}
	}

}
