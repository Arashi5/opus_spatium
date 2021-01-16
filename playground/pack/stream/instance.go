package stream

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
	case "in":
		r.streamIn()
	case "out":
		r.streamOut()
	case "err":
		r.streamError()
	default:
		err = errors.New(messages.ArgErrorMessage("Stream"))
	}

	if err != nil {
		return &err
	} else {
		return nil
	}
}

