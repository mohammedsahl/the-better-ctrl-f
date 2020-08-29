package main

import (
	"strconv"
	"github.com/fabioberger/chrome"
	. "github.com/siongui/godom"
)

func main() {
	c := chrome.NewChrome()

	tabDetails := chrome.Object{
		"active": false,
	}
	c.Tabs.Create(tabDetails, func(tab chrome.Tab) {
		notification := "Tab with id: " + strconv.Itoa(tab.Id) + " created!"
		Document.GetElementById("notification").SetInnerHTML(notification)
	})

}