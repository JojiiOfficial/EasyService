package main

import (
	"fmt"
	"os"

	_ "github.com/JojiiOfficial/SystemdGoService"
	"github.com/mkideal/cli"
)

var help = cli.HelpCommand("display help information")

const serviceFolder = "/etc/systemd/system/"

type argT struct {
	cli.Helper
}

var root = &cli.Command{
	Argv: func() interface{} { return new(argT) },
	Fn: func(ctx *cli.Context) error {
		fmt.Println("Commands:\n\n" +
			"  help     display help information" + "\n" +
			"  create   Create a systemd service(aliases creat,c" + "\n" +
			"  delete   Delete a systemd service(aliases del,d)",
		)
		return nil
	},
}

func main() {
	if err := cli.Root(root,
		cli.Tree(help),
		cli.Tree(createCMD),
		cli.Tree(deleteCMD),
	).Run(os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
