package main

import (
	"errors"
	"fmt"
	"os"
	"work_space/pkg/workspace"
)

func main() {
	var err *error
	args := os.Args

	if len(args) == 1 {
		fmt.Println(errors.New("Need more args"))
		os.Exit(1)
	}

	ws := workspace.NewService(&workspace.Config{Arguments: args})
	switch args[1] {
	case "imports":
		ws.GetImports()
	case "stream":
		err = ws.GetStreams()
	case "logger":
		err = ws.GetLogger()
	case "error":
		err = ws.GetError()
	case "gc":
		err = ws.GetGarbageCollection()
	case "draft":
		ws.GetDraft()
	default:
		e := errors.New("There is no such argument")
		err = &e
	}

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
