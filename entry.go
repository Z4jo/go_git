package main

import (
	"log"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func main() {
	selectedFiles := make([]string,1)
	localBranch := CurrentBranch()
	upStreamBranch := UpStreamBranch()
	filesStatus := CurrentBranchFileStatus(localBranch)
	branchStatus := StatusCurrentBranch(localBranch, upStreamBranch)
	app := tview.NewApplication()
	//commits := Commits()
	branches := AllBranches()
	mainArea := tview.NewButton("mainXd")
	branchStatusList := BranchList(branches)
	branchStatusTextView := BranchStatusTextView(branchStatus,upStreamBranch,localBranch)
	fileStatusList := FilesStatusList(filesStatus, &selectedFiles)
	actionButtons := ActionButtons(app)

	grid := tview.NewGrid().
		SetRows(10, 0, 7).
		SetColumns(30, 0, 30).
		SetBorders(true)

	grid.AddItem(fileStatusList, 0, 0, 4, 1, 0, 100, true).
		AddItem(actionButtons,4,0,1,1,0,100,false).
		AddItem(branchStatusTextView,5,0,1,1,0,100,false).
		AddItem(branchStatusList, 6, 0, 4, 1, 0, 100, false).
		AddItem(mainArea, 0, 1, 10, 5, 0, 100, false)

	firstRow := Rotator{[]tview.Primitive{fileStatusList,mainArea}}
	secondRow := Rotator{[]tview.Primitive{actionButtons,mainArea}}
	thirdRow := Rotator{[]tview.Primitive{branchStatusList,mainArea}}
	appLayout := AppLayout{[]Rotator{firstRow,secondRow,thirdRow},0,0}	

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
	log.Println(selectedFiles)
}
