package main
import "github.com/rivo/tview"

type AppLayout struct {
	row Rotator
	column Rotator 
}

type Rotator struct{
	arr []tview.Primitive 	
	lastIndex int
}

// direction == true goes up/ direction == false goes down
func (r *Rotator) nextPrimitive(direction bool) *tview.Primitive{
	var primitiveRet *tview.Primitive
	if direction {
		if r.lastIndex+1 > len(r.arr)-1{
			primitiveRet = &r.arr[0]
			r.lastIndex = 0
		}else{
			primitiveRet = &r.arr[r.lastIndex+1]
			r.lastIndex += 1
		}
	}else {
		if r.lastIndex-1 < 0 {
			primitiveRet = &r.arr[len(r.arr)-1]
			r.lastIndex = len(r.arr)-1
		}else{
			primitiveRet = &r.arr[r.lastIndex-1]
			r.lastIndex -= 1
		}	
	}
	return primitiveRet
}


