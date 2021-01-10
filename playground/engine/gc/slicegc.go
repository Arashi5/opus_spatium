package gc

import "runtime"

type data struct {
	i, j int
}

func (Repository) GCSlice()  {
	var n = 400000
	var structure []data
	for i :=0; i < n; i++ {
		val := int(i)
		structure = append(structure, data{val, val})
	}

	runtime.GC()
	// предотвращение преждевременной очистики сборщика мус
	_ = structure[0]
}