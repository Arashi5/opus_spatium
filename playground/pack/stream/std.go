package stream

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
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

func (repository) streamIn() {
	var f *os.File
	f = os.Stdin
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		fmt.Println(">", scanner.Text())
	}
}

func (repository) streamOut() {
	if len(os.Args) == 1 {
		fmt.Println("Need Arg")
		os.Exit(1)
	}

	args := os.Args
	min, err := strconv.ParseFloat(args[1], 64)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	max, err := strconv.ParseFloat(args[2], 64)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for i := 2; i < len(args); i++ {
		n, _ := strconv.ParseFloat(args[i], 64)

		if n < min {
			min = n
		}

		if n > max {
			max = n
		}
	}

	fmt.Println("Min:", min)
	fmt.Println("Max:", max)
}

/**
go run main.go 2>/tmp/stdError
or
go run main.go /tmp/output 2>&1
*/
func (repository) streamError() {
	ms := ""
	args := os.Args
	if len(args) == 1 {
		ms = "Need args"
	} else {
		ms = args[1]
	}

	io.WriteString(os.Stdout, "Standart output")
	io.WriteString(os.Stderr, ms)
	io.WriteString(os.Stderr, "\n")
}
