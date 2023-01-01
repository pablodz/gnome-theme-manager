package gui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
	"github.com/pablodz/gnome-theme-manager/internal/gnomelook"
)

func downloadScreen(_ fyne.Window) fyne.CanvasObject {

	err := gnomelook.GetDataFromScraper()
	if err != nil {
		panic(err)
	}

	data := gnomelook.GetNames()

	list := widget.NewList(
		func() int {
			return len(data)
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("template")
		},
		func(i widget.ListItemID, o fyne.CanvasObject) {
			o.(*widget.Label).SetText(data[i])
		})

	return list
}
