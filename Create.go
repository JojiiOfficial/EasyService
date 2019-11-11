package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path"
	"strings"

	"github.com/JojiiOfficial/SystemdGoService"
	"github.com/mkideal/cli"
)

type createT struct {
	cli.Helper
	Name        string `cli:"*N,name" usage:"Specify the name of the service"`
	ExecFile    string `cli:"*F,file" usage:"Specify the ExecStart file" `
	Description string `cli:"D,description" usage:"Specify the description of the service"`
	User        string `cli:"U,user" usage:"Specify the user for the service"`
	Group       string `cli:"G,group" usage:"Specify the group for the service"`
	Start       bool   `cli:"s,start" usage:"Starts the service after creating"`
	Enable      bool   `cli:"e,enable" usage:"Enables the service after creating"`
	Delete      bool   `cli:"!d,delete" usage:"Deletes a given service" dft:"false"`
	Yes         bool   `cli:"y,yes" usage:"Skip confirm messages" dft:"false"`
}

var createCMD = &cli.Command{
	Name:    "create",
	Aliases: []string{"create"},
	Argv:    func() interface{} { return new(createT) },
	Fn: func(ctx *cli.Context) error {
		argv := ctx.Argv().(*createT)
		if os.Getgid() != 0 {
			fmt.Println("You need to be root to run this command")
			return nil
		}
		if argv.Delete {
			if len(argv.Name) == 0 {
				fmt.Println("No name given!")
				return nil
			}
			if _, err := os.Stat("/etc/systemd/system/" + SystemdGoService.NameToServiceFile(argv.Name)); err != nil {
				fmt.Println("A service with this name doesn't exists")
				return nil
			}
			if !argv.Yes {
				reader := bufio.NewReader(os.Stdin)
				y, i := confirmInput("Do you really want to delete the service \""+argv.Name+"\" [y/n]> ", reader)
				if i == -1 || !y {
					return nil
				}
			}
			err := os.Remove("/etc/systemd/system/" + SystemdGoService.NameToServiceFile(argv.Name))
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
		}
		description := "An easy service for " + argv.Name
		if len(argv.Name) == 0 || len(argv.ExecFile) == 0 {
			fmt.Println("Missing parameter value")
			return nil
		}
		file := argv.ExecFile
		if !strings.HasPrefix(file, "/") {
			ex, err := os.Executable()
			if err != nil {
				log.Fatal(err)
			}
			dir := path.Dir(ex)
			if strings.HasPrefix(file, "./") {
				file = dir + "/" + file[2:]
			} else {
				file = dir + "/" + file
			}
		}
		if _, er := os.Stat(argv.ExecFile); er != nil {
			fmt.Println("File not found")
			return nil
		}
		if len(argv.Description) > 0 {
			description = argv.Description
		}
		if SystemdGoService.SystemfileExists(argv.Name) {
			fmt.Println("Servicename already taken")
			return nil
		}
		service := SystemdGoService.NewDefaultService(argv.Name, description, file)
		err := service.Create()
		if err != nil {
			fmt.Println("Error creating service: " + err.Error())
		} else {
			SystemdGoService.DaemonReload()
			fmt.Println("Service created successfully: \"/etc/systemd/" + SystemdGoService.NameToServiceFile(argv.Name) + "\"")
			if argv.Enable {
				err = service.Start()
				if err != nil {
					fmt.Println("Error starting service:", err.Error())
					return nil
				}
				fmt.Println("Service started successfully")
				err = service.Enable()
				if err != nil {
					fmt.Println("Error enabling service:", err.Error())
					return nil
				}
				fmt.Println("Service enabled successfully")
			}
		}
		return nil
	},
}
