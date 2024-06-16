package main

import (
	"log"
	"os/exec"
	"strings"
	"github.com/rivo/tview"
)

type CommitList struct {
	date string
	commitHash string
	author string
	msg string
}

type FileStatus struct{
	marking string
	fileName string
}

type BranchStatus struct{
	commitsAhead string
	commitsBehind string
}


func AllBranches() string{
	cmd := exec.Command("git","branch","--all")
	branches,err := cmd.Output()
	if err != nil {
		log.Fatalln(err)
	}
	return strings.TrimSpace(string(branches))
}

func Commits() *tview.List{
	list := tview.NewList()	
	// git log --pretty=format:"%h - %an, %ad : %s" --date=short
	cmd := exec.Command("git", "log", "--pretty=format:\"%h - %an, %ad : %s\"" ,"--date=short")
	commits,err := cmd.Output()
	if err != nil{
		log.Fatalln(err)
	}

	for _,commit := range strings.Split(string(commits), "\n"){
		splitedCommitMsg := strings.Split(commit, ":")

		splitedCommit := strings.Split(splitedCommitMsg[0], " ")
		c := CommitList{
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
	currentBranch,err := cmd.Output()
	if err != nil {
		log.Fatalln(err)
	}
	return string(currentBranch)
}

func UpStreamBranch() string{
	cmd:=exec.Command("git", "rev-parse","--abbrev-ref", "--symbolic-full-name","@{u}")
	upStreamBranch,err := cmd.Output()
	if err != nil{
		log.Fatalln(err)
	}
	ret := string(upStreamBranch)
	return ret
}

func CurrentBranchFileStatus(currentBranch string)[]FileStatus{
	cmd := exec.Command("git","status","-s")
	output,err := cmd.Output()
	if err != nil{
		log.Fatalln(err)
	}
	status := string(output)
	filesStatus := make([]FileStatus,0)
	for i,file := range strings.Split(status, "\n"){
		if (file == "" || file == " ") {
			if i == 0 {
				filesStatus = append(filesStatus,FileStatus{"✓","all files uptodate"})
			}
			break
		}else{
			statusArr := strings.Split(file, " ")
			if len(statusArr) > 2 {
				log.Println(len(statusArr))
				filesStatus = append(filesStatus, FileStatus{statusArr[1],statusArr[2]})
			}else{
				log.Println(statusArr)
				filesStatus = append(filesStatus, FileStatus{statusArr[0],statusArr[1]})
			}
		}
	}
	return filesStatus
}

func StatusCurrentBranch(currentBranch string,upStreamBranch string)BranchStatus{
	upStreamBranch = strings.TrimSpace(upStreamBranch)
	commitsAheadString :=  "HEAD.."+ upStreamBranch
	commitsBehindString := upStreamBranch + "..HEAD" 
	cmdAhead := exec.Command("git","rev-list","--count",commitsAheadString)	
	cmdAhead.Dir = "./"
	commitsAhead,err := cmdAhead.CombinedOutput()
	if err != nil{
		log.Println(string(commitsAhead))
		log.Fatalln(err)
	}
	cmdBehind := exec.Command("git","rev-list","--count",commitsBehindString)	
	commitsBehind,err := cmdBehind.Output()
	if err != nil{
		log.Println(string(commitsBehind))
		log.Fatalln(err)
	}

	return BranchStatus{
					strings.TrimSpace(string(commitsAhead)),
					strings.TrimSpace(string(commitsBehind))}
	}
