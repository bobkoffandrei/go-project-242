package main

import (
	"context"
	"fmt"
    "github.com/bobkoffandrei/go-project-242/code"
	"github.com/urfave/cli/v3"
	"log"
	"os"
)

func main() {

	cmd := &cli.Command{

		Name: "ADS",

		Usage: "Analyze Disk Size",

		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:    "human",
				Aliases: []string{"H"},
				Value:   false,
				Usage:   "human-readable sizes (auto-select unit)",
			},
			&cli.BoolFlag{
				Name:    "all",
				Aliases: []string{"a"},
				Value:   false,
				Usage:   "include hidden files and directories",
			},
			&cli.BoolFlag{
				Name:    "recursive",
				Aliases: []string{"r"},
				Value:   false,
				Usage:   "recursive size of directories",
			},
		},

		Action: func(ctx context.Context, c *cli.Command) error {
			result, err := code.GetPathSize(c.Args().Get(0), c.Bool("recursive"), c.Bool("human"), c.Bool("all"))
			if err != nil {
				return err
			}

			fmt.Printf("%s\t%s\n", result, c.Args().Get(0))
			return nil
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}

}
