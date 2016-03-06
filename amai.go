package amai

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
	
	data := make(chan interface{})
	errors := make(chan error)
	
	go adapter.GetCharacter(id, data, errors)
	
	select {
	case err := <- errors:
		log.Fatal(err)
	case char := <- data:
		s, err := json.MarshalIndent(char, "", "  ")
		if err != nil {
			log.Fatal(err)
		}
		
		fmt.Printf("%s\n", s)
	}
}

// Gets information about a guild, prints a JSON blob to stdout.
func GetGuild(adapter base.Adapter, c *cli.Context) {
	id := c.Args()[0]
	
	data := make(chan interface{})
	errors := make(chan error)
	
	go adapter.GetGuild(id, data, errors)
	
	select {
	case err := <- errors:
		log.Fatal(err)
	case char := <- data:
		s, err := json.MarshalIndent(char, "", "  ")
		if err != nil {
			log.Fatal(err)
		}
		
		fmt.Printf("%s\n", s)
	}
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
	app.Commands = []cli.Command{
		{
			Name: "guild",
			Usage: "Shows information about a guild",
			Aliases: []string { "g" },
			Action: func(c *cli.Context) { GetGuild(MakeAdapter(c), c) },
		},
	}
	
	app.Run(os.Args)
}
