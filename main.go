package main

import (
	"fmt"
	"os"
	"runtime"

	_ "github.com/JojiiOfficial/SystemdGoService"
	"github.com/mkideal/cli"
)

var help = cli.HelpCommand("display help information")

const serviceFolder = "/etc/systemd/system/"
const version = "1.9"
const binFile = "ezservice"

type argT struct {
	cli.Helper
	Version bool `cli:"v,version" usage:"Displays the version of easyservice"`
}

var root = &cli.Command{
	Argv: func() interface{} { return new(argT) },
	Fn: func(ctx *cli.Context) error {
		argv := ctx.Argv().(*argT)
		if argv.Version {
			fmt.Println("EasyService V."+version, runtime.GOOS+"/"+runtime.GOARCH)
		} else {
			fmt.Println("Commands:\n\n" +
				"  help     display help information" + "\n" +
				"  create   Create a systemd service(aliases creat,c" + "\n" +
				"  delete   Delete a systemd service(aliases del,d)" + "\n" +
				"  rename   Rename a service(aliases ren,r)",
			)
		}
		return nil
	},
}

func main() {
	if err := cli.Root(root,
		cli.Tree(help),
		cli.Tree(createCMD),
		cli.Tree(deleteCMD),
		cli.Tree(renameCMD),
	).Run(os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
