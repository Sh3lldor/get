package main

import (
	"os"

	"github.com/Sh3lldor/get/db"
	"github.com/urfave/cli"
)

func main() {
	get := cli.NewApp()
	get.Name = "get"
	get.Usage = "A little tool to store all your One-Liners"
	get.UsageText = "get <flags>"
	get.Version = "1.0.0"
	get.Author = "elad_pt"
	get.ArgsUsage = "dsa"
	get.Commands = []cli.Command{
		{
			Name:  "list",
			Usage: "List all stored commands",
			Action: func(c *cli.Context) {
				db.GetAllCommands()
			},
		},
		{
			Name:  "set",
			Usage: "Set new command <CommandName> <FullCommand>",
			Action: func(c *cli.Context) {
				commandName := c.Args().Get(0)
				fullCommand := c.Args().Get(1)
				if commandName != "" && fullCommand != "" {
					db.SetNewCommand(commandName, fullCommand)
				}
			},
		},
		{
			Name:  "show",
			Usage: "Show specific command (By ID or Name)",
			Action: func(c *cli.Context) {
				commandIdentifier := c.Args().Get(0)
				if commandIdentifier != "" {
					db.ShowSpesificCommand(commandIdentifier)
				}
			},
		},
		{
			Name:  "delete",
			Usage: "Delete specific commnad (By ID or Name)",
			Action: func(c *cli.Context) {
				commandIdentifier := c.Args().Get(0)
				if commandIdentifier != "" {
					db.DeleteSpesificCommand(commandIdentifier)
				}
			},
		},
		{
			Name:  "copy",
			Usage: "Copy specific command to clipboard (By ID)",
			Action: func(c *cli.Context) {
				commandIdentifier := c.Args().Get(0)
				if commandIdentifier != "" {
					db.CopySpesificCommand(commandIdentifier)
				}
			},
		},
		{
			Name:  "reset",
			Usage: "Reset the DB",
			Action: func(c *cli.Context) {
				db.ResetDB()
			},
		},
	}
	err := get.Run(os.Args)
	if err != nil {
		os.Exit(1)
	}

}
