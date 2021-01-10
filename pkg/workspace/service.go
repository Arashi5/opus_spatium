package workspace

import (
	"errors"
	"fmt"

	"work_space/playground/engine/gc"
	er "work_space/playground/pack/err"
	"work_space/playground/pack/imports"
	"work_space/playground/pack/logger"
	"work_space/playground/pack/stream"
)

type WorkSpace interface {
	GetImports()
	GetStreams() error
	GetLogger() error
	GetError() error
	GetGarbageCollection()
}

type service struct {
	Arg string
	Rep *PGRepositories
}

type Config struct {
	Arg string
	Rep *PGRepositories
}

type PGRepositories struct {
	Err *er.Repository
	Log *logger.Repository
	Imp *imports.Repository
	Str *stream.Repository
	GC  *gc.Repository
}

func NewService(cfg *Config) *service  {
	return &service{
		Arg: cfg.Arg,
		Rep: cfg.Rep,
	}
}

func (s service) GetImports()  {
	s.Rep.Imp.SimpleImportModule()
}

func (s service) GetStreams() error {
	var err error
	r := s.Rep.Str
	switch s.Arg {
	case "in":
		r.StreamIn()
		break
	case "out":
		r.StreamOut()
		break
	case "err":
		r.StreamError()
		break
	default:
		err = errors.New(errorMessage("Stream"))
	}

	if err != nil {
		return err
	} else {
		return nil
	}
}

func (s service) GetLogger() error {
	var err error
	r := s.Rep.Log
	switch s.Arg {
	case "log":
		r.Log()
		break
	case "fatal":
		r.LogFatal()
		break
	case "panic":
		r.LogPanic()
		break
	case "custom":
		r.CustomLog()
		break
	default:
		err = errors.New(errorMessage("Logger"))
	}

	if err != nil {
		return err
	} else {
		return nil
	}
}

func (s service) GetError() error {
	var err error
	r := s.Rep.Err
	switch s.Arg {
	case "return":
		r.ReturnError()
		break
	case "example":
		r.ExampleError()
		break
	default:
		err = errors.New(errorMessage("Err"))
	}

	if err != nil {
		return err
	} else {
		return nil
	}
}

func (s service) GetGarbageCollection()  {
	s.Rep.GC.GCExample()
}

func errorMessage(p string) string {
	s := fmt.Sprintf("There is no such argument in packet %s", p)
	return s
}