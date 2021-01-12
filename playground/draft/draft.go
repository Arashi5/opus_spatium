package draft

import (
	"fmt"
	"time"
)

type Exec struct {
	repo repository
}

type repository struct{}

func NewRepo() *Exec {
	return &Exec{repo: repository{}}
}

func (e Exec) Exec(_ []string) *error {
	e.repo.checkGlobalVar()
	return nil
}

const durationSec = 5

type FileMeta struct {
	Size int
}

var file FileMeta

func (repository) checkGlobalVar() {
	tic := time.Tick(time.Second * durationSec)
	var i int
	for range tic {
		fmt.Println(file.Size)
		i++
		file.Size = i
	}
}
