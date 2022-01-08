package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

type StopInfo struct {
	XMLName   xml.Name `xml:"stopInfo"`
	Text      string   `xml:",chardata"`
	Created   string   `xml:"created,attr"`
	Stop      string   `xml:"stop,attr"`
	StopAbv   string   `xml:"stopAbv,attr"`
	Message   string   `xml:"message"`
	Direction []struct {
		Text string `xml:",chardata"`
		Name string `xml:"name,attr"`
		Tram []struct {
			Text        string `xml:",chardata"`
			Destination string `xml:"destination,attr"`
			DueMins     string `xml:"dueMins,attr"`
		} `xml:"tram"`
	} `xml:"direction"`
}

func makeListOfStringsFromStopInfo(sI StopInfo) ([]string, []string) {
	inboundInfo := make([]string, 0)
	outboundInfo := make([]string, 0)

	for _, dir := range sI.Direction {
		if strings.Contains(dir.Name, "Inbound") {
			for _, tram := range dir.Tram {
				inboundInfo = append(inboundInfo, tram.DueMins+" min(s) "+tram.Destination)
			}
		} else if strings.Contains(dir.Name, "Outbound") {
			for _, tram := range dir.Tram {
				outboundInfo = append(outboundInfo, tram.DueMins+" min(s) "+tram.Destination)
			}
		}
	}

	return inboundInfo, outboundInfo
}

func buildLuasUrlString(stopName string) string {
	return "http://luasforecasts.rpa.ie/xml/get.ashx?action=forecast&stop=" + stopName + "&encrypt=false"
}

func makeApiCallAndParseXml(url string, stopInfoXML *StopInfo) {
	resp, err := http.Get(url)

	if err != nil {
		log.Fatalln(err)
		os.Exit(1)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if err := xml.Unmarshal(body, &stopInfoXML); err != nil {
		fmt.Println("Error in parsing xml")
		os.Exit(1)
	}

}
