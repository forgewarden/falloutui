package main

import (
	"github.com/rivo/tview"
)

func main() {

	stationList := tview.NewList()
	stationList.AddItem("Galaxy News Radio", "", '\n', nil)
	stationList.AddItem("Enclave Radio", "", '\n', nil)

	lMapTab := newTab("Local Map")
	wMapTab := newTab("World Map")
	questsTab := newTab("Quests")
	notesTab := newTab("Notes")
	radioTab := newTab("Radio")

	radioTab.SetBorder(true)

	grid := tview.NewGrid()

	grid.SetTitle("DATA")
	grid.SetRows(0, 3)
	grid.SetColumns(0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0)
	grid.SetBorder(true)

	grid.AddItem(stationList, 0, 0, 1, 11, 0, 0, true)

	grid.AddItem(lMapTab, 1, 1, 1, 1, 0, 0, false)
	grid.AddItem(wMapTab, 1, 3, 1, 1, 0, 0, false)
	grid.AddItem(questsTab, 1, 5, 1, 1, 0, 0, false)
	grid.AddItem(notesTab, 1, 7, 1, 1, 0, 0, false)
	grid.AddItem(radioTab, 1, 9, 1, 1, 0, 0, false)

	if err := tview.NewApplication().SetRoot(grid, true).SetFocus(grid).Run(); err != nil {
		panic(err)
	}
}

func newTab(text string) *tview.TextView {
	tab := tview.NewTextView()

	tab.SetTextAlign(tview.AlignCenter)
	tab.SetText(text)

	return tab
}
