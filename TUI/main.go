package main

import (
	"github.com/gdamore/tcell/v2"
	"github.com/golangci/golangci-lint/test/testdata_etc/unused_exported/lib"
	"github.com/rivo/tview"
)

// song informations

type Song struct {
	duration float64
	title    string
	artist   string
	album    string
}

type App struct {
	tview          *tview.Application
	rootContainter *tview.Flex
}

func main() {
	app := tview.NewApplication()

	// 2x2

	flex := tview.NewFlex().
		AddItem(tview.NewFlex().SetDirection(tview.FlexRow).
			AddItem(tview.NewBox().SetBorder(true).SetTitle("Library"), 0, 1, false).
			AddItem(tview.NewBox().SetBorder(true).SetTitle("Lyrics"), 0, 1, false), 0, 1, false).

		// 2x2

		AddItem(tview.NewFlex().SetDirection(tview.FlexRow).
			AddItem(tview.NewBox().SetBorder(true).SetTitle("Songs"), 0, 10, false).
			AddItem(tview.NewBox().SetBorder(true).SetTitle("Player"), 0, 1, false), 0, 3, false)

	// Set up an event loop
	app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {

		switch event.Rune() {
		case 'k':
			flex.SetBorder(true)

		case 'j':
			flex.SetBorder(false)
		}

		// Check for the q key press
		if event.Rune() == 'q' {
			app.Stop()
			return nil
		}

		return event

	})

	// The application will run until app.Stop() is called (in response to the Escape key press)

	if err := app.SetRoot(flex, true).Run(); err != nil {
		panic(err)
	}

}
