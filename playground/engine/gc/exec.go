package gc

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
	case "example":
		r.gcExample()
	case "slice":
		err = r.gcSlice(args[1])
	default:
		err = errors.New(messages.ArgErrorMessage("GC"))
	}

	if err != nil {
		return &err
	}

	return nil
}
