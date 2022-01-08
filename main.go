package main

import (
	"log"
	"strings"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

func main() {
	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer ui.Close()

	url := buildLuasUrlString("ran")
	var stopInfoXML StopInfo

	makeApiCallAndParseXml(url, &stopInfoXML)

	abbrRecs := parseAbbreviationsFromCsv("places.csv")

	abbrStringList, placesStringList := abbreviationRecordsToStringLists(abbrRecs)
	inboundInfo, outboundInfo := makeListOfStringsFromStopInfo(stopInfoXML)

	_ = abbrStringList

	// Instantiate UI elements
	inboundWidget := widgets.NewParagraph()
	inboundWidget.Text = strings.Join(inboundInfo[:], "\n")
	inboundWidget.SetRect(0, 4, 50, 14)

	outboundWidget := widgets.NewParagraph()
	outboundWidget.Text = strings.Join(outboundInfo[:], "\n")
	outboundWidget.SetRect(50, 4, 100, 14)

	footer := widgets.NewParagraph()
	footer.Text = stopInfoXML.Message
	footer.SetRect(0, 14, 100, 18)

	header := widgets.NewList()
	header.Title = "Stops"
	header.Rows = placesStringList
	header.TextStyle = ui.NewStyle(ui.ColorYellow)
	header.WrapText = false
	header.SetRect(0, 0, 100, 4)

	uiElements := UIElements{header, inboundWidget, outboundWidget, footer}

	renderUIElements(&uiElements)

	// Render setup UI
	// TODO: Move to own file

	selectedRow := 0
	selectedRowAbbr := ""
	uiEvents := ui.PollEvents()
	for {
		e := <-uiEvents
		selectedRow = uiElements.hdr.SelectedRow
		selectedRowAbbr = abbrStringList[selectedRow]
		switch e.ID {
		case "q", "<C-c>", "<Escape>":
			return
		case "j", "<Down>":
			uiElements.hdr.ScrollDown()
		case "k", "<Up>":
			uiElements.hdr.ScrollUp()
		case "<Enter>":
			updateUiWithApiCall(&stopInfoXML, selectedRowAbbr, &uiElements)
		}

		renderUIElements(&uiElements)
		// ui.Render(inboundWidget)
		// ui.Render(outboundWidget)
		// ui.Render(footer)
		// ui.Render(header)
	}

}
