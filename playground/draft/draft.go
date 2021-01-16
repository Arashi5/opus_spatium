package draft

import (
	"fmt"
	"time"
)

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
