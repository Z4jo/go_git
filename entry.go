package main

import (
	//"github.com/gdamore/tcell/v2"
	"fmt"
	"os/exec"
	"strings"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)


type commitList struct {
	date string
	commitHash string
	author string
	msg string
}



func getGitCommits() *tview.List{
	list := tview.NewList()	
	// git log --pretty=format:"%h - %an, %ad : %s" --date=short
	cmd := exec.Command("git", "log", "--pretty=format:\"%h - %an, %ad : %s\"" ,"--date=short")
	cmd.Dir = "./"
	commits,err := cmd.Output()
	if err != nil{
		panic(err)
	}

	for _,commit := range strings.Split(string(commits), "\n"){
		splitedCommitMsg := strings.Split(commit, ":")

		splitedCommit := strings.Split(splitedCommitMsg[0], " ")
		c := commitList{
			date :splitedCommit[2],
			commitHash: splitedCommit[0],
			author: splitedCommit[2],
			msg:splitedCommitMsg[1],
		} 
		list.AddItem(c.commitHash,c.date,'a',nil)
	}
	fmt.Println(string(commits))
	
	return list
}


func main() {  
	app := tview.NewApplication()
	list := getGitCommits()
	grid := tview.NewGrid()
	capture := func(event *tcell.EventKey) *tcell.EventKey {
		if event.Rune() == 'l'{
			app.SetFocus(grid.)				
		}
		return event
	}
	grid.SetInputCapture(capture)
	grid.SetColumns()
	grid.SetRows()
	grid.AddItem(list,0,0,3,3,0,0,false)
	app.SetRoot(grid,false).Run()
}
