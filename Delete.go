package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/JojiiOfficial/SystemdGoService"
	"github.com/mkideal/cli"
)

type deleteT struct {
	cli.Helper
	Name string `cli:"*N,name" usage:"Specify the name of the service"`
	Yes  bool   `cli:"y,yes" usage:"Skip confirm messages" dft:"false"`
}

var deleteCMD = &cli.Command{
	Name:    "delete",
	Desc:    "Delete a systemd service",
	Aliases: []string{"del", "d"},
	Argv:    func() interface{} { return new(deleteT) },
	Fn: func(ctx *cli.Context) error {
		argv := ctx.Argv().(*deleteT)
		reader := bufio.NewReader(os.Stdin)
		if len(argv.Name) == 0 {
			fmt.Println("No name given!")
			return nil
		}
		if _, err := os.Stat(serviceFolder + SystemdGoService.NameToServiceFile(argv.Name)); err != nil {
			fmt.Println("A service with this name doesn't exists")
			return nil
		}
		if !argv.Yes {
			y, i := confirmInput("Do you really want to delete the service \""+argv.Name+"\" [y/n]> ", reader)
			if i == -1 || !y {
				return nil
			}
		}
		err := os.Remove(serviceFolder + SystemdGoService.NameToServiceFile(argv.Name))
		if err != nil {
			fmt.Println("Error deleting file: " + err.Error())
		} else {
			fmt.Println("Service " + SystemdGoService.NameToServiceFile(argv.Name) + " deleted!")
			err = SystemdGoService.DaemonReload()
			if err != nil {
				fmt.Println("Error reloading daemon: " + err.Error())
				return nil
			}
			fmt.Println("Daemon reloaded successfully")
		}
		return nil
	},
}
