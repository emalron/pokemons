package main

import (
    "net/http"
    "github.com/PuerkitoBio/goquery"
    "log"
)

type Town struct {
    Name string `json:"name"`
    Url string  `json:"url"`
}

func GetPage(url string) *goquery.Document {
    res, err := http.Get(url)
    if err != nil {
        log.Fatal(err)
    }
    defer res.Body.Close()
    if res.StatusCode != 200 {
        log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
    }
    doc, err := goquery.NewDocumentFromReader(res.Body)
    if err != nil {
        log.Fatal(err)
    }
    return doc
}

func GetRegions() []Town {
    towns := []Town{}
    doc := GetPage("https://pokemon.fandom.com/ko/wiki/%EC%84%B1%EB%8F%84%EC%A7%80%EB%B0%A9")
    context := doc.Find("table")
    context.Each(func(i int, s *goquery.Selection) {
        if i == 17 {
            s.Find("tr").Find("td").Find("a").Each(func(i int, s *goquery.Selection) {
                href, ok := s.Attr("href")
                if ok {
                    log.Println(i, goquery.NodeName(s), s.Text(), href)
                    item := Town{Name: s.Text(), Url: href}
                    towns = append(towns, item)
                }
            })
        }
    })
    return towns
}
