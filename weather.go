// vim: ft=go ts=4 sw=4 et ai:
package main

import (
    "net/http"
    "fmt"
    "log"
    "io/ioutil"
    "encoding/xml"
    )

var feed_url = "http://weather.gc.ca/rss/city/on-118_e.xml"

type RSS struct {
    XMLName xml.Name `xml:"feed"`
    ItemList []Item `xml:"entry"`
}

type Item struct {
    Title string `xml:"title"`
    Link string `xml:link"`
    Summary string `xml:"summary"`
}

func main() {
    var rss RSS
    res, err := http.Get(feed_url)
    if err != nil {
        log.Fatal(err)
    }
    body, err := ioutil.ReadAll(res.Body)
    res.Body.Close()
    document := string(body)
    //log.Println(document)
    err = xml.Unmarshal([]byte(document), &rss)
    if err != nil {
        log.Fatal(err)
    }
    //log.Printf("%#v", rss)

    for _, item := range rss.ItemList {
        fmt.Printf("%s\n", item.Title)
    }
}
