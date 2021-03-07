package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var (
	stack []string
	page []string
	fault int
)

func initialized() {
	stack = make([]string, 3)
	page = make([]string, 3)
	stack[0] = "-"
	stack[1] = "-"
	stack[2] = "-"
	page[0] = "-"
	page[1] = "-"
	page[2] = "-"
	fault = 0
}

func showProcess() {
	fmt.Printf("\n--------------------\n")
	fmt.Printf("|  %s  |  %s  |  %s  |", page[0], page[1], page[2])
	fmt.Printf("\n--------------------\n")
	fmt.Printf("\nPage-fault = %d\n\n", fault)
	fmt.Printf("(|  %s  |  %s  |  %s  |)\n\n", stack[0], stack[1], stack[2])
	fmt.Printf("\nCommand > ")
}

func getCommand() string {
	reader := bufio.NewReader(os.Stdin)
	data, _ := reader.ReadString('\n')
	data = strings.Trim(data, "\n")
	return data
}

func command_create(p string) {
	if stack[0] == "-" && stack[1] == "-" && stack[2] == "-" {
		stack[0] = p
		page[0] = stack[0]
		fault++
	} else if stack[0] != "-" && stack[1] == "-" && stack[2] == "-" {
		stack[1] = p
		page[1] = stack[1]
		fault++
	} else if stack[0] != "-" && stack[1] != "-" && stack[2] == "-" {
		stack[2] = p
		page[2] = stack[2]
		fault++
	} else {
		if p != stack[0] && p != stack[1] &&  p != stack[2] {
			for i := range page {
				if page[i] == stack[0] {
					page[i] = p
					copy(stack[0:], stack[1:])
					stack[2] = p
					fault++
					break
				}
			} 
		} else if p == stack[2] {
			stack[2] = p
			page[2] = stack[2]
			//fmt.Printf("push")
		} else if p == stack[1] {
			stack[1] = stack[2]
			stack[2] = p
		} else {
			copy(stack[0:], stack[1:])
			stack[2] = p
		}
	}
}

func main() {
	initialized()
	for {
		showProcess()
		command := getCommand()
		switch command {
		case "1","2","3","4","5","6","7","8","9","0": command_create(command)
		case "exit": return
		default: fmt.Printf("\nSystax error!\n")			
		}
	}
}