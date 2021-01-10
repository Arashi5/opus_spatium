package gc

import (
	"errors"
	"runtime"
)

var N = 1000000

type data struct {
	i, j int
}

func (Repository) GCSlice(s string) error {
	var err error
	switch s {
	case "s":
		slice()
		break
	case "ms":
		sliceMapWithStar()
		break
	case "mns":
		sliceMapWithNoStar()
		break
	case "mst":
		sliceMapSplit()
		break
	default:
		err = errors.New("There is no such argument in packet gc; function GCSlice()")
	}

	if err != nil {
		return err
	} else {
		return nil
	}
}

func  slice()  {
	var structure []data
	for i :=0; i < N; i++ {
		v := int(i)
		structure = append(structure, data{v, v})
	}

	runtime.GC()
	// предотвращение преждевременной очистики сборщика мус
	_ = structure[0]
}

func sliceMapWithStar() {
	m := make(map[int]*int)
	for i := 0; i < N; i++ {
		v := int(i)
		m[v] = &v
	}
	runtime.GC()
	_ = m[0]
}

func sliceMapWithNoStar() {
	m := make(map[int]int)
	for i := 0; i < N; i++ {
		v := int(i)
		m[v] = v
	}
	runtime.GC()
	_ = m[0]
}

func sliceMapSplit() {
	s := make([]map[int]int, 200)
	for i := range s {
		s[i] = make(map[int]int)
	}
	for i :=0; i < N; i ++ {
		v := int(i)
		s[i%200][v] = v
	}

	runtime.GC()
	_=s[0][0]
}