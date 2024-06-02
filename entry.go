package main

import (
	//"github.com/gdamore/tcell/v2"
	"fmt"
	"os/exec"
	"github.com/rivo/tview"
)

func getGitCommits() *tview.List{
	list := tview.NewList()	
	// git log --pretty=format:"%h - %an, %ad : %s" --date=short
	cmd := exec.Command("git", "log", "--pretty=format:\"%h - %an, %ad : %s\"" ,"--date=short")

	cmd.Dir = "./"
	commits,err := cmd.Output()
	if err != nil{
		panic(nil)
	}
	fmt.Println(string(commits))
	
	return list
}

func main() {  
	app := tview.NewApplication()
	list := getGitCommits()
	fmt.Println(list)
	grid := tview.NewGrid()
	
	app.SetRoot(grid,false)




}
