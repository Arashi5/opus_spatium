package workspace

import (
	"errors"
	"fmt"
)

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
	switch s.Arg[0] {
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
	switch s.Arg[0]  {
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
	switch s.Arg[0]  {
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

//time go run main.go gc slice s|ms|mns|mst
func (s service) GetGarbageCollection() error {
	var err error
	r := s.Rep.GC
	switch s.Arg[0]  {
	case "example":
		r.GCExample()
		break
	case "slice":
		err = r.GCSlice(s.Arg[1])
		break
	default:
		err = errors.New(errorMessage("GC"))
	}

	if err != nil {
		return err
	} else {
		return nil
	}
}

func errorMessage(p string) string {
	s := fmt.Sprintf("There is no such argument in packet %s", p)
	return s
}
