package err

import (
	"errors"
	"fmt"
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

func (repository) returnError() {
	err := returnError(1, 2)
	if err != nil {
		fmt.Println("returnError() ended normally")
	} else {
		fmt.Println(err)
	}
	err = returnError(10, 10)
	if err == nil {
		fmt.Println("returnError() ended normally")
	} else {
		fmt.Println(err)
	}

	if err.Error() == "Error in returnError() function" {
		fmt.Println("!!")
	}
}

func returnError(a, b int) error {
	if a == b {
		err := errors.New("Error in returnError() function")
		return err
	}

	return nil
}

//In an ideal world - os.Exit () should be used in the "main" function
func (repository) exampleError() {
	if len(os.Args) == 1 {
		fmt.Println("Need args")
		os.Exit(1)
	}

	args := os.Args
	err := errors.New("An error")
	k := 1
	var n float64
	for err != nil {
		if k >= len(args) {
			fmt.Println("None of the args is a float")
			return
		}
		n, err = strconv.ParseFloat(args[k], 64)
		k++
	}

	min, max := n, n

	for i := 2; i < len(args); i++ {
		n, err := strconv.ParseFloat(args[i], 64)
		if err == nil {
			if n < min {
				min = n
			}
			if n > max {
				max = n
			}
		}
	}

	fmt.Println("Min:", min)
	fmt.Println("Max:", max)
}
