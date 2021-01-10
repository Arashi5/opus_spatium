package main

import (
	"errors"
	"fmt"
	"os"
	"work_space/pkg/workspace"
)

func main() {
	args := os.Args
	var err error
	if len(args) == 1 {
		err = errors.New("Need more args")
		fmt.Println(err)
		os.Exit(1)
	}
	var aa string
	if len(args) < 3 {
		aa = ""
	} else {
		aa = args[2]
	}
	ws := workspace.NewService(&workspace.Config{Arg: aa})
	switch args[1] {
	case "imports":
		ws.GetImports()
		break
	case "stream":
		err = ws.GetStreams()
		break
	case "logger":
		err = ws.GetLogger()
		break
	case "error":
		err = ws.GetError()
	case "gc":
		ws.GetGarbageCollection()
		break
	default:
		err = errors.New("There is no such argument")
	}

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
