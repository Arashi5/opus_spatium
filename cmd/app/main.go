package main

import (
	"errors"
	"fmt"
	"os"
	"work_space/pkg/workspace"
	"work_space/playground/engine/gc"
	"work_space/playground/pack/imports"
	"work_space/playground/pack/logger"
	er "work_space/playground/pack/err"
	"work_space/playground/pack/stream"
)

func main() {
	var err error
	args := os.Args
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

	ws := workspace.NewService(&workspace.Config{
		Arg: aa,
		Rep: getPGRepoList() ,
	})
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

func getPGRepoList() *workspace.PGRepositoriesCollection {
	return  &workspace.PGRepositoriesCollection{
		Str: stream.NewRepo(),
		Log: logger.NewRepo(),
		Err: er.NewRepo(),
		Imp: imports.NewRepo(),
		GC:  gc.NewRepo(),
	}
}