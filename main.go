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
		fmt.Println("Usage: ezservice <service/create/delete/modify> [-N,-F,-D,-U,-G,-s,-e,-d,-y]")
		return nil
	},
}

func main() {
	if err := cli.Root(root,
		cli.Tree(help),
		cli.Tree(createCMD),
	).Run(os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
