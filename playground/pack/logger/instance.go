package logger

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
	switch args[0] {
	case "log":
		e.repo.log()
	case "fatal":
		e.repo.logFatal()
	case "panic":
		e.repo.logPanic()
	case "custom":
		e.repo.customLog()
	default:
		err = errors.New(messages.ArgErrorMessage("Logger"))
	}

	if err != nil {
		return &err
	} else {
		return nil
	}
}
