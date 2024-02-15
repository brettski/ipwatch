package main

import (
	"log"

	"github.com/brettski/go-ipwatch/datastore"
	"github.com/brettski/go-ipwatch/watcher"
	cli "github.com/jawher/mow.cli"
)

func cmdRun(cmd *cli.Cmd) {
	verbose := cmd.BoolOpt("v verbose", false, "Show verbose  progress on stdout")
	cmd.Action = func() { fullCheck(*verbose) }
}

func fullCheck(isVerbose bool) {
	foundIp, err := watcher.RunIpWatcherCheck(isVerbose)
	if err != nil {
		// log for now
		log.Fatalf("Error running watcher: %s", err)
	}

	log.Printf("Retrieved IP: %s", foundIp)

	// next step database
	ipExists := datastore.IsIpInStore(foundIp)
	if ipExists {
		// update datastore as seen
		seen, err := datastore.IncrementSeen(foundIp)
		if err != nil {
			log.Fatalf("Error incrementing record: %s", err)
		}
		log.Printf("IP %s has now been seen %d times", foundIp, seen)
	} else {
		// Create new record, send email, whatever else
		err := datastore.AddNew(foundIp)
		if err != nil {
			log.Fatalf("Error adding new record: %s", err)
		}
		// TODO: Send email
		log.Printf("IP %s new record added", foundIp)
	}
}
