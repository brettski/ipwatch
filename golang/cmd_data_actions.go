package main

import (
	"github.com/brettski/go-ipwatch/datastore"
	cli "github.com/jawher/mow.cli"
)

func cmdDataActions(cmd *cli.Cmd) {
	dumpall := cmd.BoolOpt("a dumpall", false, "Write all records to stdout")

	cmd.Action = func() {
		if *dumpall {
			datastore.DumpAllRecords()
		}

	}
}

// func cmdDataWriteAll(cmd *cli.Cmd) {
// 	// cmd.Action = func() { log.Fatal("Not Implemented") }
// 	cmd.Action = datastore.DumpAllRecords
// }
