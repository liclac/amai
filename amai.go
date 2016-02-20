package main

import (
  "fmt"
  "log"
  "github.com/PuerkitoBio/goquery"
)

func main() {
  doc, err := goquery.NewDocument("http://na.finalfantasyxiv.com/lodestone/character/7248246/") 
  if err != nil {
    log.Fatal(err)
  }
  
  nameBox := doc.Find(".player_name_txt")
  name := nameBox.Find("h2 a").Text()
  fmt.Printf("Name: %s\n", name)
}
