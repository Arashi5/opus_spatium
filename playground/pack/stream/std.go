package stream

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

func StreamIn() {
	var f * os.File
	f = os.Stdin
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		fmt.Println(">", scanner.Text())
	}
}

func StreamOut()  {
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
		n,_ := strconv.ParseFloat(args[i],64)

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
func StreamError() {
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