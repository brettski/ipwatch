package main

import (
	"log"

	"github.com/brettski/go-ipwatch/watcher"
	cli "github.com/jawher/mow.cli"
)

func cmdRun(cmd *cli.Cmd) {
	cmd.Action = fullCheck
}

func fullCheck() {
	foundIp, err := watcher.RunIpWatcherCheck()
	if err != nil {
		// log for now
		log.Fatalf("Error running watcher: %s", err)
	}

	log.Printf("Retrieved IP: %s", foundIp)

	// next step database
}
