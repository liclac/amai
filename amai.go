package main

import (
	"os"
	"fmt"
	"log"
	"github.com/codegangsta/cli"
	"github.com/uppfinnarn/amai/adapters"
	"github.com/uppfinnarn/amai/ffxiv"
)

func GetCharacter(adapter adapters.Adapter, c *cli.Context) {
	id := c.Args()[0]
	
	char, err := adapter.GetCharacter(id)
	if err != nil {
		log.Fatal(err)
	}
	
	fmt.Printf("Name: %s\n", char.Name)
}

func main() {
	var adapter adapters.Adapter = ffxiv.NewAdapter()
	
	app := cli.NewApp()
	app.Name = "amai"
	app.Usage = "Parse and process data from FFXIV"
	app.Version = "0.0.1"
	app.Commands = []cli.Command{
		{
			Name: "char",
			Usage: "Shows information about a character",
			Aliases: []string { "c" },
			Action: func(c *cli.Context) { GetCharacter(adapter, c) },
		},
	}
	
	app.Run(os.Args)
}
