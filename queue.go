package main

import(
	"github.com/rivo/tview"
)

type AppLayout struct {
	arr []Rotator
	lastIndexRow int
	lastIndexColumn int
}

type Rotator struct{
	arr []tview.Primitive 	
}

// direction == true goes up/ direction == false goes down
func(a *AppLayout) nextNodeVertical(direction bool) *tview.Primitive{
	lastColumnIndex := a.lastIndexColumn
	var primitiveRet *tview.Primitive
	if direction {
		if a.lastIndexRow+1 > len(a.arr)-1{
			rotator := a.arr[0]	
			primitiveRet = &rotator.arr[lastColumnIndex]
			a.lastIndexRow = 0
		}else{
			rotator := a.arr[a.lastIndexRow+1]	
			primitiveRet = &rotator.arr[lastColumnIndex]
			a.lastIndexRow += 1
		}
	}else{
		if a.lastIndexRow-1 < 0 {
			rotator := a.arr[len(a.arr)-1]	
			primitiveRet = &rotator.arr[lastColumnIndex]
			a.lastIndexRow = len(a.arr)-1
		}else{
			rotator := a.arr[a.lastIndexRow-1]	
			primitiveRet = &rotator.arr[lastColumnIndex]
			a.lastIndexRow -= 1
		}
	}
	return primitiveRet
}


func (a *AppLayout) nextNodeHorizontal(direction bool) *tview.Primitive{
	rowArr := a.arr[a.lastIndexRow].arr
	var primitiveRet *tview.Primitive
	if direction {
		if a.lastIndexColumn+1 > len(rowArr)-1{
			primitiveRet = &rowArr[0]
			a.lastIndexColumn = 0
		}else{
			primitiveRet = &rowArr[a.lastIndexColumn+1]
			a.lastIndexColumn += 1
		}
	}else{
		if a.lastIndexColumn-1 < 0 {
			primitiveRet = &rowArr[len(rowArr)-1]
			a.lastIndexColumn = len(rowArr)-1
		}else{
			primitiveRet = &rowArr[a.lastIndexColumn - 1]
			a.lastIndexColumn -= 1
		}	
	}
	return primitiveRet
}
