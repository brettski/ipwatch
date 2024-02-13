package main

import (
	cli "github.com/jawher/mow.cli"
)

func cmdConfigActions(cmd *cli.Cmd) {
	cmd.Command("dump", "Writes current configuration to  stdout", cmdDumpConfig)
}

func cmdDumpConfig(cmd *cli.Cmd) {
	cmd.Action = dumpConfigToStdOut
}
