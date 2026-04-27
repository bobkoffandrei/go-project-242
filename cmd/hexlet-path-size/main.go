package main

import (
//	"fmt"
	"github.com/urfave/cli/v3"
	"os"
    "context"
	"github.com/bobkoffandrei/go-project-242/pathsize"
)

func main(){
(&cli.Command{}).Run(context.Background(), os.Args)


GetPathSize("~/go-project-242", false, false, false)

}
