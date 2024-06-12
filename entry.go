package main

import (
	//"github.com/gdamore/tcell/v2"
	//"github.com/gdamore/tcell/v2"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)



func main() {
<<<<<<< HEAD

=======
>>>>>>> e753d9cc194589bdf0bd8ca69b455adec46b42d9
	app := tview.NewApplication()
	commits := Commits()
	branches := AllBranches()
	newPrimitive := func(text string) tview.Primitive {
		return tview.NewTextView().
			SetTextAlign(tview.AlignCenter).
			SetText(text)
	}
	main := newPrimitive("Main content")

<<<<<<< HEAD
	header := newPrimitive("header")	
	footer := newPrimitive("footer")	
=======
>>>>>>> e753d9cc194589bdf0bd8ca69b455adec46b42d9
	
	grid := tview.NewGrid().
		SetRows(3, 0, 3).
		SetColumns(30, 0, 30).
		SetBorders(true).
<<<<<<< HEAD
		AddItem(header, 0, 0, 1, 3, 0, 0, false).
		AddItem(footer, 2, 0, 1, 3, 0, 0, false)
=======
		AddItem(newPrimitive("Header"), 0, 0, 1, 3, 0, 0, false).
		AddItem(newPrimitive("Footer"), 2, 0, 1, 3, 0, 0, false)
>>>>>>> e753d9cc194589bdf0bd8ca69b455adec46b42d9

	// Layout for screens narrower than 100 cells (menu and side bar are hidden).
	grid.AddItem(commits, 0, 0, 0, 0, 0, 0, false).
		AddItem(main, 1, 0, 1, 3, 0, 0, false).
		AddItem(branches, 0, 0, 0, 0, 0, 0, false)

	// Layout for screens wider than 100 cells.
	grid.AddItem(commits, 1, 0, 1, 1, 0, 100, false).
<<<<<<< HEAD
		AddItem(main, 1, 1, 1, 1, 0, 100, false).
		AddItem(branches, 1, 2, 1, 1, 0, 100, false)

	
	firstRow := []tview.Primitive{commits,header,branches}
	secondRow := []tview.Primitive{commits,main,branches}
	thirdRow := []tview.Primitive{commits,footer,branches}
	rows := Rotator{firstRow,0}	
	rows := Rotator{secondRow,0}	
	rows := Rotator{thirdRow,0}	
	

=======
		AddItem(commits, 1, 0, 1, 1, 0, 100, false).
		AddItem(main, 1, 1, 1, 1, 0, 100, false).
		AddItem(branches, 1, 2, 1, 1, 0, 100, false)

>>>>>>> e753d9cc194589bdf0bd8ca69b455adec46b42d9
	capture := func(event *tcell.EventKey)*tcell.EventKey{
		if event.Rune() == 'l'{
			app.SetFocus(commits)	
		}else if event.Rune() == 'j'{
			app.SetFocus(branches)
		}

		
		return event
	}
	grid.SetInputCapture(capture)	
	if err := app.SetRoot(grid, true).Run(); err != nil {
		panic(err)
	}
}
