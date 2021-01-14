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

func (e Exec) Exec(args []string) *error {
	switch args[0] {
	case "v":
		e.repo.checkGlobalVar()
	case "f":
		e.repo.ftpFileChecker()
	}
	return nil
}

const durationSec = 5

var file FileMeta

func (repository) checkGlobalVar() {
	tic := time.Tick(time.Second * durationSec)
	var i int
	for range tic {
		fmt.Println(file.Size)
		i++
		file.Size = uint64(i)
	}
}
