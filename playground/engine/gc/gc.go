package gc

import (
	"fmt"
	"runtime"
	"time"
)

type GarbageCollection interface {
	SimpleImportModule()
}

type Repository struct {}

func NewRepo() *Repository  {
	return &Repository{}
}

//GODEBUG=gctrace=1 go run *.go
func (Repository) GCExample()  {
	var mem runtime.MemStats
	printStats(mem)

	for i := 0; i < 10; i++ {
		s := make([]byte, 50000000)
		if s == nil {
			fmt.Println("Operation failed!")
		}
	}
	printStats(mem)

	for i := 0; i < 10; i++ {
		s := make([]byte, 100000000)
		if s == nil {
			fmt.Println("Operation failed!")
		}
		time.Sleep(time.Second * 5)
	}
	printStats(mem)
}


func printStats(mem runtime.MemStats) {
	runtime.ReadMemStats(&mem)
	fmt.Println("mem.Alloc:", mem.Alloc)
	fmt.Println("mem.TotalAlloc:", mem.TotalAlloc)
	fmt.Println("mem.HeapAlloc:", mem.HeapAlloc)
	fmt.Println("mem.NumGC:", mem.NumGC)
	fmt.Println("----------")
}