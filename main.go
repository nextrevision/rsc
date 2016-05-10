package main

import (
	"fmt"
	"os"

	"github.com/nextrevision/go-runscope"
	"github.com/nextrevision/rsc/cmd"
)

var client *runscope.Client

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}
