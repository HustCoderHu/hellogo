package ch12

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"

	"github.com/sirupsen/logrus"
)

func F1_reader() {
	inputReader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter some input, plz")
	input, err := inputReader.ReadString('\n')
	if err == nil {
		fmt.Printf("input was: %s\n", input)
	}
}

func F2_file_rw() {
	inputFile, inputErr := os.Open("go.mod")
	if inputErr != nil {
		fmt.Printf("An error occurred on opening the inputfile\n" +
			"Does the file exist?\n" +
			"Have u got access to it?\n")
		return
	}
	defer inputFile.Close()

	inputReader := bufio.NewReader(inputFile)

	for {
		inputStr, inputErr := inputReader.ReadString('\n')
		// fmt.Printf("input was:%s", inputStr)
		logrus.Infof("input was:%s", inputStr)
		if inputErr == io.EOF {
			return
		}
	}
}

func F5_read_write_file1() {
	inputFile := "go.mod"
	outputFile := "go.mod.bak"
	buf, err := ioutil.ReadFile(inputFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "File err: %s\n", err)
		// panic(err.Error())
	}

	logrus.Info(string(buf))
	err = ioutil.WriteFile(outputFile, buf, 0644)
	if err != nil {
		panic(err.Error())
	}
}

func F11_os_args() {
	who := "Alice "
	if len(os.Args) >= 2 {
		who += strings.Join(os.Args[1:], " ")
	}
	fmt.Println("Good morning", who)
}
