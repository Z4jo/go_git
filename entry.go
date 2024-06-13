package main

import (
	//"github.com/gdamore/tcell/v2"
	//"github.com/gdamore/tcell/v2"
//	"fmt"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)
func main() {
	app := tview.NewApplication()
	commits := Commits()
	branches := AllBranches()
	mainArea := tview.NewButton("mainUwuwa")
	currentBranch := CurrentBranch()
	statusCurrentBranch := StatusCurrentBranch(currentBranch)
	//titleColor:= tcell.ColorNames["black"]
		
	grid := tview.NewGrid().
		SetRows(7, 0, 7).
		SetColumns(30, 0, 30).
		SetBorders(true)
	
	// Layout for screens wider than 100 cells.
	grid.AddItem(commits, 0, 0, 2, 1, 0, 100, true).
		AddItem(branches, 2, 0, 2, 1, 0, 100, false).
		AddItem(mainArea, 0, 1, 2, 2, 0, 100, false)

	//Layout abstraction initialization	
	firstRow := Rotator{[]tview.Primitive{commits,mainArea}}
	secondRow := Rotator{[]tview.Primitive{branches,mainArea}}
	//thirdRow := Rotator{[]tview.Primitive{commits,footer,branches}}
	appLayout := AppLayout{[]Rotator{firstRow,secondRow},0,0}	

	capture := func(event *tcell.EventKey)*tcell.EventKey{
		if event.Rune() == 'l'{
			app.SetFocus(*appLayout.nextNodeHorizontal(true))
		}else if event.Rune() == 'h'{	
			app.SetFocus(*appLayout.nextNodeHorizontal(false))
		}else if event.Rune() == 'j'{
			app.SetFocus(*appLayout.nextNodeVertical(true))	
		}else if event.Rune() == 'k'{
			app.SetFocus(*appLayout.nextNodeVertical(false))	
		}		
		
		return event
	}
	grid.SetInputCapture(capture)	
	if err := app.SetRoot(grid, true).Run(); err != nil {
		panic(err)
	}
}
