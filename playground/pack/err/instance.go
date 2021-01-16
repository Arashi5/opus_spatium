package err

import (
	"errors"
	"work_space/pkg/messages"
)

type Exec struct {
	repo repository
}

type repository struct{}

func NewRepo() *Exec {
	return &Exec{repo: repository{}}
}

func (e Exec) Exec(args []string) *error {
	var err error
	r := e.repo
	switch args[0] {
	case "return":
		r.returnError()
	case "example":
		r.exampleError()
	default:
		err = errors.New(messages.ArgErrorMessage("Error"))
	}

	if err != nil {
		return &err
	}

	return nil
}