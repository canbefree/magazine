package main

import (
	"github.com/canbefree/magazine/cmd"
	_ "github.com/canbefree/magazine/startup"
)

func main() {
	// log.Info(context.TODO(), "curr bin path"+vars.RootPath)
	cmd.Execute()
}
