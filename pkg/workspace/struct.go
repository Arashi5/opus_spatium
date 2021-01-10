package workspace

import (
	"work_space/playground/engine/gc"
	"work_space/playground/pack/err"
	"work_space/playground/pack/imports"
	"work_space/playground/pack/logger"
	"work_space/playground/pack/stream"
)

type service struct {
	Arg string
	Rep *PGRepositories
}

type Config struct {
	Arg string
	Rep *PGRepositories
}

type PGRepositories struct {
	Err *err.Repository
	Log *logger.Repository
	Imp *imports.Repository
	Str *stream.Repository
	GC  *gc.Repository
}
