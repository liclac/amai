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

// Makes an adapter for the specified game
func MakeAdapter(c *cli.Context) base.Adapter {
	game := c.GlobalString("game")
	
	switch game {
	case "ffxiv":
		return ffxiv.NewAdapter()
	default:
		log.Fatal("Unknown game: ", game)
	}
	return nil
}

func main() {
	app := cli.NewApp()
	app.Name = "amai"
	app.Usage = "Parse and process data from FFXIV"
	app.Version = "0.0.1"
	app.Flags = []cli.Flag {
		cli.StringFlag {
			Name: "game",
			Usage: "Game to connect to",
			Value: "ffxiv",
		},
	}
	app.Commands = []cli.Command{
		{
			Name: "char",
			Usage: "Shows information about a character",
			Aliases: []string { "c" },
			Action: func(c *cli.Context) { GetCharacter(MakeAdapter(c), c) },
		},
	}
	
	app.Run(os.Args)
}
