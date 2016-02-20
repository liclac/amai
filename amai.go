package main

import (
	"os"
	"fmt"
	"net/http"
	"log"
	"github.com/codegangsta/cli"
	"github.com/PuerkitoBio/goquery"
)

func main() {
	client := &http.Client{}
	
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
				
				req, err := http.NewRequest("GET", url, nil)
				if err != nil {
					log.Fatal(err)
				}
				req.Header.Add("Cookie", "ldst_touchstone=1;ldst_is_support_browser=1;ldst_visit=1")
				req.Header.Add("User-Agent", "Mozilla/5.0 (iPhone; CPU iPhone OS 9_1 like Mac OS X) AppleWebKit/601.1.46 (KHTML, like Gecko) Version/9.0 Mobile/13B143 Safari/601.1")
				
				res, err := client.Do(req)
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
