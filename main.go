package main

import (
	"fmt"
	"os"

	_ "github.com/JojiiOfficial/SystemdGoService"
	"github.com/mkideal/cli"
)

var help = cli.HelpCommand("display help information")

type argT struct {
	cli.Helper
}

var root = &cli.Command{
	Argv: func() interface{} { return new(argT) },
	Fn: func(ctx *cli.Context) error {
		fmt.Println("Usage: ezservice <install/disable/start/stop>")
		return nil
	},
}

func main() {
	if err := cli.Root(root,
		cli.Tree(help),
	).Run(os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
