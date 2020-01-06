package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/JojiiOfficial/SystemdGoService"
	"github.com/mkideal/cli"
)

type renameT struct {
	cli.Helper
	Name      string `cli:"*N,name" usage:"Specify the name of the service"`
	NewName   string `cli:"*R,new-name" usage:"Specify the new-name of the service"`
	Overwrite bool   `cli:"o,overwrite" usage:"Overwrite an existing servic"`
}

var renameCMD = &cli.Command{
	Name:    "rename",
	Aliases: []string{"ren", "r"},
	Desc:    "Rename a service",
	Argv:    func() interface{} { return new(renameT) },
	Fn: func(ctx *cli.Context) error {
		if os.Getuid() != 0 {
			return errors.New("you need to be root")
		}
		argv := ctx.Argv().(*renameT)
		if !SystemdGoService.SystemfileExists(argv.Name) {
			return errors.New("Service does not exist")
		}
		if SystemdGoService.SystemfileExists(argv.NewName) {
			if !argv.Overwrite {
				return errors.New("Service already exists. Use -o to overwrite it")
			}
			os.Remove(serviceFolder + SystemdGoService.NameToServiceFile(argv.NewName))
		}
		err := os.Rename(serviceFolder+SystemdGoService.NameToServiceFile(argv.Name), serviceFolder+SystemdGoService.NameToServiceFile(argv.NewName))
		if err != nil {
			fmt.Println("Error renaming service:", err.Error())
			return nil
		}
		err = SystemdGoService.DaemonReload()
		if err != nil {
			fmt.Println("Error reloading daemon:", err.Error())
			return nil
		}
		fmt.Println("Daemon reload succesful")

		return nil
	},
}
