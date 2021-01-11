package draft

import (
	"fmt"
	"time"
)
const durationSec = 5

type Repository struct {}

func NewRepo() *Repository  {
	return &Repository{}
}


type FileMeta struct {
	Size int
}

var file FileMeta

func(Repository) CheckVar() {
	tic := time.Tick(time.Second * durationSec)
	var i int
	for range tic {
		fmt.Println(file.Size)
		i++
		file.Size = i
	}
}