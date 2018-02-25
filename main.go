package main

import (
	"fmt"
	"os"

	// or "runtime"
	"github.com/dhrp/moulincli/command"
	"github.com/mitchellh/cli"
)

// func cleanup() {
// 	fmt.Println("cleanup")
// }

func main() {

	// sigchan := make(chan os.Signal, 2)
	// signal.Notify(sigchan, os.Interrupt, syscall.SIGTERM)
	// go func() {
	// 	<-sigchan
	// 	cleanup()
	// 	os.Exit(1)
	// }()

	ui := &cli.BasicUi{
		Reader:      os.Stdin,
		Writer:      os.Stdout,
		ErrorWriter: os.Stderr,
	}

	c := cli.NewCLI("cliexample", "0.0.1")
	c.Args = os.Args[1:]

	c.Commands = map[string]cli.CommandFactory{

		"load": func() (cli.Command, error) {
			return &command.LoadCommand{
				Ui: &cli.ColoredUi{
					Ui:          ui,
					OutputColor: cli.UiColorGreen,
				},
			}, nil
		},
		"work": func() (cli.Command, error) {
			return &command.Work{
				Ui: &cli.ColoredUi{
					Ui:          ui,
					OutputColor: cli.UiColorGreen,
				},
			}, nil
		},
		"create": func() (cli.Command, error) {
			return &command.CreateTask{
				Ui: &cli.ColoredUi{
					Ui:          ui,
					OutputColor: cli.UiColorGreen,
				},
			}, nil
		},
		"progress": func() (cli.Command, error) {
			return &command.Progress{
				Ui: &cli.ColoredUi{
					Ui:          ui,
					OutputColor: cli.UiColorGreen,
				},
			}, nil
		},
		"peek": func() (cli.Command, error) {
			return &command.Peek{
				Ui: &cli.ColoredUi{
					Ui:          ui,
					OutputColor: cli.UiColorGreen,
				},
			}, nil
		},
	}

	exitStatus, err := c.Run()
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
	}

	os.Exit(exitStatus)
}
