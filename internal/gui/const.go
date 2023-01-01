package gui

import "fyne.io/fyne/v2"

// Tutorial defines the data structure for a tutorial
type Tutorial struct {
	Title, Intro string
	View         func(w fyne.Window) fyne.CanvasObject
	SupportWeb   bool
}

var (
	// Tutorials defines the metadata for each tutorial
	Tutorials = map[string]Tutorial{
		"welcome": {
			"Welcome",
			"",
			welcomeScreen,
			true,
		},
		"installed": {
			"Install Themes",
			"List of current themes",
			currentThemesScreen,
			true,
		},
		"download": {
			"Download Themes",
			"Download themes from gnome-look.org",
			downloadScreen,
			true,
		},
	}

	// TutorialIndex  defines how our tutorials should be laid out in the index tree
	TutorialIndex = map[string][]string{
		"":            {"welcome", "installed", "download"},
		"collections": {"list", "table", "tree"},
		// "containers":  {"apptabs", "border", "box", "center", "doctabs", "grid", "scroll", "split"},
		// "widgets":     {"accordion", "button", "card", "entry", "form", "input", "progress", "text", "toolbar"},
	}
)
