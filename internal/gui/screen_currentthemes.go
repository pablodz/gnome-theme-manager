package gui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
	"github.com/pablodz/gnome-theme-manager/utils"
)

func currentThemesScreen(_ fyne.Window) fyne.CanvasObject {

	data, err := utils.GetDirectories("/usr/share/themes")
	if err != nil {
		panic(err)
	}

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
