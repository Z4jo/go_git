package main

import (
	"log"
	"slices"
	"strings"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func branchStatusFormer(branchStatus BranchStatus, localBranch string, upstreamBranch string) string{
	ret := branchStatus.commitsAhead + "↑ "+localBranch + branchStatus.commitsBehind + "↓ " + upstreamBranch
	return ret
}

func BranchList(branches string) *tview.List{
	list := tview.NewList()
	for _,branch := range strings.Split(branches,"\n"){
		list.AddItem(branch,"",' ',nil)
	}
	return list
}

func BranchStatusTextView(branchStatus BranchStatus, upstreamBranch string, localBranch string)*tview.TextView{
	textView := tview.NewTextView()
	formatedBranchStatus := branchStatusFormer(branchStatus,localBranch,upstreamBranch)
	textView.SetText(formatedBranchStatus)
	return textView
}

func FilesStatusList(filesStatus []FileStatus, selectedFiles *[]string)*tview.List{
	log.Println(filesStatus)
	list := tview.NewList()

	for _,file := range filesStatus{
		formatedString := file.marking + " " + file.fileName
		list.AddItem(formatedString,"",' ',nil)
	}
	list.SetSelectedFunc(func(i int, s1, s2 string, r rune) {
		isSelected := slices.Index(*selectedFiles,s1)
		if isSelected > -1{
			list.RemoveItem(i)
			list.AddItem(s1,"",' ',nil)
			updatedSlice := slices.Delete(*selectedFiles,isSelected,isSelected+1)	
			*selectedFiles = updatedSlice	
		}else {
			list.RemoveItem(i)
			list.AddItem(s1,"[green] added",'✓',nil)
			updatedSlice := append(*selectedFiles,s1)
			*selectedFiles = updatedSlice
		}
	})
	return list
}

func ActionButtons(app *tview.Application) *tview.Flex {
	mergeButton := tview.NewButton("MERGE")
	addButton := tview.NewButton("ADD")
	commitButton := tview.NewButton("COMMIT")
	flex := tview.NewFlex().AddItem(mergeButton,0,1,true).
					AddItem(addButton,0,1,false).
					AddItem(commitButton,0,1,false)
	
	buttons := [3]tview.Primitive{mergeButton,addButton,commitButton}
	lastIndex := 0

	capture := func  (event *tcell.EventKey) *tcell.EventKey{
		if event.Rune() == 'j'{
			return event}
		if event.Rune() == 'k'{
			return event}
		if event.Rune() == 'h'{
			if lastIndex - 1 < 0{
				return event
			}else{
				app.SetFocus(buttons[lastIndex-1])	
				lastIndex-=1
			}}
		if event.Rune() == 'l'{
			if lastIndex + 1 > len(buttons)-1{
				return event
			}else{
				app.SetFocus(buttons[lastIndex+1])	
				lastIndex+=1
			}
		}
		
		return event	
	}
	flex.SetInputCapture(capture)
	return flex 
}
