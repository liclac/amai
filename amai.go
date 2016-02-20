package main

import (
	"os"
	"fmt"
	"log"
	"github.com/codegangsta/cli"
	"github.com/PuerkitoBio/goquery"
	"github.com/uppfinnarn/amai/adapters"
	"github.com/uppfinnarn/amai/ffxiv"
)

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
			Action: func(c *cli.Context) {
				id := c.Args()[0]
				url := fmt.Sprintf("http://na.finalfantasyxiv.com/lodestone/character/%s/", id)
				
				res, err := adapter.Get(url)
				if err != nil {
					log.Fatal(err)
				}
				defer res.Body.Close()
				
				if res.StatusCode != 200 {
					log.Fatal("Error: ", res.Status)
				}
				
				doc, err := goquery.NewDocumentFromResponse(res)
				if err != nil {
					log.Fatal(err)
				}
				
				name := doc.Find(".txt_charaname").Text()
				fmt.Printf("Name: %s\n", name)
			},
		},
	}
	
	app.Run(os.Args)
	
	/*doc, err := goquery.NewDocument("http://na.finalfantasyxiv.com/lodestone/character/7248246/")
	if err != nil {
		log.Fatal(err)
	}

	nameBox := doc.Find(".player_name_txt")
	name := nameBox.Find("h2 a").Text()
	fmt.Printf("Name: %s\n", name)*/
}
