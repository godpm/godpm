package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/godpm/godpm/pkg/http"
	"github.com/godpm/godpm/pkg/log"

	"github.com/urfave/cli/v2"
)

func main() {
	var addr string
	app := &cli.App{
		Name:  "godpm",
		Usage: "Golang based deploy and process manager",
		Action: func(c *cli.Context) error {
			_ = handleInput(addr)
			return nil
		},

		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "remote_addr",
				Aliases:     []string{"r"},
				Value:       "http://127.0.0.1:10086",
				Destination: &addr,
			},
		},

		Commands: []*cli.Command{
			{
				Name:  "status",
				Usage: "get process's status",
				Action: func(c *cli.Context) error {
					args := c.Args().Slice()
					return handleStatus(args, addr)
				},
			},
			{
				Name:  "stop",
				Usage: "stop a process",
				Action: func(c *cli.Context) error {
					return handleStop(c.Args().Slice(), addr)
				},
			},
			{
				Name:  "start",
				Usage: "stop a process",
				Action: func(c *cli.Context) error {
					return handleStart(c.Args().Slice(), addr)
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal().Fatal(err)
	}
}

func handleInput(addr string) error {
	_ = handleStatus([]string{}, addr)

	buf := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("godpm> ")
		sentence, err := buf.ReadBytes('\n')
		if err != nil {
			return err
		}

		if len(sentence) != 0 && sentence[0] != '\n' {
			splits := strings.Split(strings.TrimSuffix(string(sentence), "\n"), " ")
			if len(splits) == 0 {
				continue
			}
			args := []string{}
			if len(splits) > 1 {
				args = splits[1:]
			}

			switch splits[0] {
			case "status":
				err = handleStatus(args, addr)
			case "start":
				err = handleStart(args, addr)
				if err != nil {
					fmt.Println(err)
					continue
				}

				err = handleStatus([]string{}, addr)

			case "stop":
				err = handleStop(args, addr)
				if err != nil {
					fmt.Println(err)
					continue
				}
				err = handleStatus(args, addr)
			default:
				fmt.Printf("command '%s' does not support \n", splits[0])
			}

			if err != nil {
				fmt.Println(err)
			}
		}
	}
}

var svc *http.Service

func getService(addr string) *http.Service {
	if svc == nil {
		svc = http.NewService(addr, "", true)
	}
	return svc
}

func handleStop(args []string, addr string) error {
	svc := getService(addr)
	if len(args) > 0 {
		return svc.Stop(args[0])
	}

	err := fmt.Errorf("no args for stop")
	return err
}

func handleStart(args []string, addr string) error {
	svc := getService(addr)
	if len(args) > 0 {
		return svc.Start(args[0])
	}

	err := fmt.Errorf("no args for start")
	return err
}

func handleStatus(args []string, addr string) error {
	svc := getService(addr)

	procs, err := svc.Status(args)
	if err != nil {
		return err
	}

	prettyPrintProcessInfo(procs)
	return nil
}

func prettyPrintProcessInfo(procs []http.ProcStatus) {
	nameFieldLength := 30
	for _, proc := range procs {
		fmt.Print(proc.Name)
		fmt.Print(strings.Repeat(" ", nameFieldLength-len(proc.Name)))
		fmt.Print(proc.State)
		fmt.Print(strings.Repeat(" ", 12-len(proc.State)))
		pid := fmt.Sprintf("pid %d", proc.Pid)
		fmt.Print(pid)
		fmt.Print(strings.Repeat(" ", 12-len(pid)))
		uptime := time.Since(proc.Uptime).String()
		fmt.Print(uptime)
		fmt.Println("")
	}
}
