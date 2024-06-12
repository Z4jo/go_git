package main

import(
	"github.com/rivo/tview"
	"os/exec"
	"strings"
)

type commitList struct {
	date string
	commitHash string
	author string
	msg string
}



func AllBranches() *tview.List{
	list := tview.NewList()	
	cmd := exec.Command("git","branch","--all")
	cmd.Dir = "./"
	branches,err := cmd.Output()
	if err != nil {
		panic(err)
	}
	for _,branch := range strings.Split(string(branches),"\n"){
		list.AddItem(branch," ", ' ', nil)
	}

	return list
}

func Commits() *tview.List{
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
		list.AddItem(c.commitHash,c.date,' ',nil)
	}
	return list
}

func CurrentBranch() string{
	cmd:= exec.Command("git", "rev-parse", "--abbrev-ref", "HEAD")
	cmd.Dir = "./"
	currentBranch,err := cmd.Output()
	if err != nil {
		panic(err)
	}
	return string(currentBranch)
}

func StatusCurrentBranch(currentBranch string){


}
