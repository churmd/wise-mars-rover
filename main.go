package main

import (
	"os"

	"github.com/churmd/wise-mars-rover/cli"
)

func main(){
	interactor := cli.New(os.Stdin, os.Stdout)
	interactor.Run()
}