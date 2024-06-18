package main

import (
	"log"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"os"
)

func main() {
	file, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)
    if err != nil {
        log.Fatalf("Failed to open log file: %v", err)
    }
	defer file.Close()

    log.SetOutput(file)
	selectedFiles := make([]string,1)
	localBranch := CurrentBranch()
	upStreamBranch := UpStreamBranch()
	filesStatus := CurrentBranchFileStatus(localBranch)
	branchStatus := StatusCurrentBranch(localBranch, upStreamBranch)
	app := tview.NewApplication()
	//commits := Commits()
	branches := AllBranches()
	mainArea := tview.NewButton("mainXd")
	mainArea2 := tview.NewButton("mainXd2")
	branchStatusList := BranchList(branches)
	branchStatusTextView := BranchStatusTextView(branchStatus,upStreamBranch,localBranch)
	fileStatusList := FilesStatusList(filesStatus, &selectedFiles)
	//actionButtons := ActionButtons(app)
	add:= tview.NewButton("add")
	merge:= tview.NewButton("merge")
	comm:= tview.NewButton("comm")
	grid := tview.NewGrid().
		SetRows(10, 0, 7).
		SetColumns(10, 10, 10, 0, 30).
		SetBorders(true)

	grid.AddItem(fileStatusList, 0, 0, 2, 3, 0, 100, true).
	//	AddItem(actionButtons,2,0,1,1,0,100,false).
		AddItem(add,2,0,1,1,0,100,false).
		AddItem(comm,2,1,1,1,0,100,false). 
		AddItem(merge,2,2,1,1,0,100,false).
		AddItem(branchStatusTextView,3,0,1,3,0,100,false).
		AddItem(branchStatusList, 4, 0, 4, 3, 0, 100, false).
		AddItem(mainArea, 0, 3, 10, 3, 0, 100, false).
		AddItem(mainArea2, 0, 6, 10, 3, 0, 100, false)

	firstRow := Rotator{[]tview.Primitive{fileStatusList,fileStatusList,fileStatusList,mainArea}}
	secondRow := Rotator{[]tview.Primitive{add,comm,merge,mainArea}}
	thirdRow := Rotator{[]tview.Primitive{branchStatusList,branchStatusList,branchStatusList,mainArea}}
	appLayout := AppLayout{[]Rotator{firstRow,secondRow,thirdRow},0,0}	

	capture := func(event *tcell.EventKey)*tcell.EventKey{
		if event.Rune() == 'l'{
			log.Println("l pressed")
			app.SetFocus(*appLayout.nextNodeHorizontal(true))
		}else if event.Rune() == 'h'{	
			log.Println("h pressed")
			app.SetFocus(*appLayout.nextNodeHorizontal(false))
		}else if event.Rune() == 'j'{
			log.Println("j pressed")
			app.SetFocus(*appLayout.nextNodeVertical(true))	
		}else if event.Rune() == 'k'{
			log.Println("k pressed")
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
