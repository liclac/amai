package main

import (
	"os"
	"fmt"
	"log"
	"encoding/json"
	"github.com/codegangsta/cli"
	"github.com/uppfinnarn/amai/base"
	"github.com/uppfinnarn/amai/ffxiv"
)

// Gets information about a character, prints a JSON blob to stdout.
func GetCharacter(adapter base.Adapter, c *cli.Context) {
	id := c.Args()[0]
	
	char, err := adapter.GetCharacter(id)
	if err != nil {
		log.Fatal(err)
	}
	
	s, err := json.MarshalIndent(char, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	
	fmt.Printf("%s\n", s)
}

func main() {
	var adapter base.Adapter = ffxiv.NewAdapter()
	
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
