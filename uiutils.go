package main

import (
	"strings"

	"github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

type UIElements struct {
	hdr    *widgets.List
	inbnd  *widgets.Paragraph
	outbnd *widgets.Paragraph
	ftr    *widgets.Paragraph
}

func updateUiWithApiCall(stopInfoXML *StopInfo, stopName string, uielems *UIElements) {
	url := buildLuasUrlString(stopName)
	makeApiCallAndParseXml(url, stopInfoXML)

	inboundInfo, outboundInfo := makeListOfStringsFromStopInfo(*stopInfoXML)

	uielems.outbnd.Text = strings.Join(outboundInfo[:], "\n")
	uielems.inbnd.Text = "Updated: " + strings.Join(inboundInfo[:], "\n")

	renderUIElements(uielems)
}

func renderUIElementMap(uielems map[string]termui.Drawable) {
	for _, elem := range uielems {
		termui.Render(elem)
	}
}

func renderUIElements(uielems *UIElements) {
	termui.Render(uielems.hdr)
	termui.Render(uielems.inbnd)
	termui.Render(uielems.outbnd)
	termui.Render(uielems.ftr)
}
