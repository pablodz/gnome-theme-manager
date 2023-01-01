package gui

import (
	"log"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/go-git/go-git/v5"
	"github.com/pablodz/gnome-theme-manager/internal/gnomelook"
)

var THEME_USER_DIR = os.Getenv("HOME") + "/.themes/"

func downloadScreen(_ fyne.Window) fyne.CanvasObject {

	err := gnomelook.GetDataFromScraper()
	if err != nil {
		panic(err)
	}

	longRepos, shortRepos, zipUrls := gnomelook.GetItems()

	list := widget.NewList(
		func() int {
			return len(shortRepos)
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("template")
		},
		func(i widget.ListItemID, o fyne.CanvasObject) {
			o.(*widget.Label).SetText(shortRepos[i])
		},
	)

	list.OnSelected = func(id widget.ListItemID) {
		log.Println("-------------" + shortRepos[id] + "-------------" + longRepos[id])

		progressBar := widget.NewProgressBar()
		w := fyne.CurrentApp().NewWindow("Downloading zip repo")
		// change size
		w.Resize(fyne.NewSize(600, 120))
		// create container
		container := container.NewVBox(
			widget.NewLabel("Downloading zip repo"),
			widget.NewLabel(zipUrls[id]),
			progressBar,
		)

		w.SetContent(container)
		w.Show()

		// clone the repository
		_, err := git.PlainClone(THEME_USER_DIR+shortRepos[id], false, &git.CloneOptions{
			URL:               longRepos[id],
			RecurseSubmodules: git.DefaultSubmoduleRecursionDepth,
			Progress:          os.Stdout,
			// Auth:              &http.BasicAuth{user: "user", Password: "pass"},
		})
		if err != nil {
			// show in window
			log.Println(err)
			w.SetContent(widget.NewLabel(err.Error()))
		}
		progressBar.SetValue(1.0)
		w.Close()
	}

	return list
}
