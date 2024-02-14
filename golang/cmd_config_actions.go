package main

import (
	"github.com/brettski/go-ipwatch/config"
	cli "github.com/jawher/mow.cli"
)

func cmdConfigActions(cmd *cli.Cmd) {
	cmd.Command("dump", "Writes current configuration to  stdout", cmdDumpConfig)
}

func cmdDumpConfig(cmd *cli.Cmd) {
	cmd.Action = config.DumpConfigToStdOut
}
