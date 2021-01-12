package workspace

import (
	"work_space/playground/draft"
	"work_space/playground/engine/gc"
	"work_space/playground/pack/err"
	"work_space/playground/pack/imports"
	"work_space/playground/pack/logger"
	"work_space/playground/pack/stream"
)

func getPGRepoCollection() *PGRepositoriesCollection {
	return &PGRepositoriesCollection{
		Str: stream.NewRepo(),
		Log: logger.NewRepo(),
		Err: err.NewRepo(),
		Imp: imports.NewRepo(),
		GC:  gc.NewRepo(),
		D:   draft.NewRepo(),
	}
}
