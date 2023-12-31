package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Commands: []*cli.Command{
			{
				Name:  "a",
				Usage: "git add",
				Action: func(ctx *cli.Context) error {

					cmd := exec.Command("git", "add", ".")
					stdout, err := cmd.Output()
					
				if err != nil {
					return err
				}

				fmt.Print(string(stdout))
				return nil
				},
			},
			{
				Name:  "c",
				Usage: "git commit",
				Action: func(ctx *cli.Context) error {
					cmd := exec.Command("git", "commit", "-am", "It worked!")
					stdout, err := cmd.Output()
					
				if err != nil {
					return err
				}

				fmt.Print(string(stdout))
				return nil
				},
			},
			{
				Name:  "p",
				Usage: "git push",
				Action: func(ctx *cli.Context) error {
					cmd := exec.Command("git", "push")
					stdout, err := cmd.Output()
					
				if err != nil {
					return err
				}

				fmt.Print(string(stdout))
				return nil
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}