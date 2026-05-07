package main

import (
	"context"
	"fmt"
	"github.com/bobkoffandrei/go-project-242/code"
	"github.com/urfave/cli/v3"
	"log"
	"os"
	"io"
)

func main() {

	f, err := os.OpenFile("terminal.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}

	defer func() {
    if err := f.Close(); err != nil {
        fmt.Fprintf(os.Stderr, "ошибка закрытия файла логов: %v\n", err)
    }
}()


	mw := io.MultiWriter(os.Stdout, f)

	

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

			if c.Args().Get(0) == "" {
			return fmt.Errorf("отсутствуют агрументы")
			}

			result, err := code.GetPathSize(c.Args().Get(0), c.Bool("recursive"), c.Bool("human"), c.Bool("all"))
			if err != nil {
				return err
			}


			//fmt.Printf("%s\t%s\n", result, c.Args().Get(0))
			_, err = fmt.Fprintf(mw, "%s\t%s\n", result, c.Args().Get(0))
			if err != nil {
				return err
			}
			return nil
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}

	

}
