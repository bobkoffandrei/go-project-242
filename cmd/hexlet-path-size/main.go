package main

import (
//	"fmt"
	"github.com/urfave/cli/v3"
	"os"
    "context"
	"github.com/bobkoffandrei/go-project-242/pathsize"
	"fmt"
	"log"
)



func main() {

    cmd := &cli.Command{

        Name:  "ADS",

        Usage: "Analyze Disk Size",

        Action: func(ctx context.Context, c *cli.Command) error {
		result, err := pathsize.GetPathSize(c.Args().Get(0), false, false, false)
        	if err != nil {
		//fmt.Println("Error:", err)
		return err
	}
        fmt.Printf("%sB\t%s\n", result, c.Args().Get(0))
            return nil
        },
        // Флаги, которые можно передать программе
        Flags: []cli.Flag{
            &cli.StringFlag{
                
            },
        },
    }

    // Запускаем приложение, передаём аргументы командной строки
    if err := cmd.Run(context.Background(), os.Args); err != nil {
        log.Fatal(err)
    }

	
}
