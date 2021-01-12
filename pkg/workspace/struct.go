package workspace

import (
	"work_space/playground/draft"
	"work_space/playground/engine/gc"
	"work_space/playground/pack/err"
	"work_space/playground/pack/imports"
	"work_space/playground/pack/logger"
	"work_space/playground/pack/stream"
)

type service struct {
	Arguments  Arguments
	Repository *PGRepositoriesCollection
}

type Config struct {
	Arguments Arguments
}

type Arguments []string

type PGRepositoriesCollection struct {
	Err *err.Exec
	Log *logger.Exec
	Imp *imports.Exec
	Str *stream.Exec
	GC  *gc.Exec
	D   *draft.Exec
}
