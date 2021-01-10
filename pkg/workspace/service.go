package workspace

import (
	"errors"
	"fmt"
	"work_space/playground/engine/gc"
	"work_space/playground/pack/imports"
	"work_space/playground/pack/logger"
	"work_space/playground/pack/stream"
	e "work_space/playground/pack/err"
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
}

type Config struct {
	Arg string
}

func NewService(cfg *Config) service  {
	return service{
		Arg: cfg.Arg,
	}
}

func (s service) GetImports()  {
	imports.SimpleImportModule()
}

func (s service) GetStreams() error {
	var err error
	switch s.Arg {
	case "in":
		stream.StreamIn()
		break
	case "out":
		stream.StreamOut()
		break
	case "err":
		stream.StreamError()
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
	switch s.Arg {
	case "log":
		logger.Log()
		break
	case "fatal":
		logger.LogFatal()
		break
	case "panic":
		logger.LogPanic()
		break
	case "custom":
		logger.CustomLog()
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
	switch s.Arg {
	case "return":
		e.ReturnError()
		break
	case "example":
		e.ExampleError()
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
	gc.GCExample()
}

func errorMessage(p string) string {
	s := fmt.Sprintf("There is no such argument in packet %s", p)
	return s
}