package main

import (
	"fmt"
	"os"

	cli "github.com/jawher/mow.cli"
)

func main() {
	fmt.Printf("brettski's ipwatch\n\n")

	app := cli.App("go-ipwatch", "brettski's ipwatch (golang edition)")
	app.Version("v version", "0.1.0")
	app.Action = func() { fmt.Println("Use -h for command help.") }
	app.Command("run", "run get against endpoint", cmdRun)
	app.Command("config", "Configuration actions", cmdConfigActions)
	app.Command("data", "Interact with datastore", cmdDataActions)

	app.Run(os.Args)

	// config := getEnvConfig()

}
