package gui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/cmd/fyne_demo/data"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/pablodz/gnome-theme-manager/utils"
)

func welcomeScreen(_ fyne.Window) fyne.CanvasObject {
	logo := canvas.NewImageFromResource(data.FyneScene)
	logo.FillMode = canvas.ImageFillContain
	if fyne.CurrentDevice().IsMobile() {
		logo.SetMinSize(fyne.NewSize(192, 192))
	} else {
		logo.SetMinSize(fyne.NewSize(256, 256))
	}

	return container.NewCenter(container.NewVBox(
		widget.NewLabelWithStyle("Welcome to the Gnome Theme Manager", fyne.TextAlignCenter, fyne.TextStyle{Bold: true}),
		widget.NewLabelWithStyle("Desktop app to download, install and erase latests version of your loved themes", fyne.TextAlignCenter, fyne.TextStyle{Bold: false}),
		logo,
		widget.NewHyperlink("Source", utils.ParseURL("https://github.com/pablodz/gnome-theme-manager")),

		// container.NewHBox(
		// 	widget.NewHyperlink("Source", utils.ParseURL("https://github.com/pablodz/gnome-theme-manager")),
		// 	// widget.NewLabel("-"),
		// 	// widget.NewHyperlink("documentation", utils.ParseURL("https://developer.fyne.io/")),
		// 	// widget.NewLabel("-"),
		// 	// widget.NewHyperlink("sponsor", utils.ParseURL("https://fyne.io/sponsor/")),
		// ),
		widget.NewLabel(""), // balance the header on the tutorial screen we leave blank on this content
	))
}
