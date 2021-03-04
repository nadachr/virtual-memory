package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var (
	frame []string
	fault int
)

func initialized() {
	frame = make([]string, 3)
	frame[0] = "-"
	frame[1] = "-"
	frame[2] = "-"
	fault = 0
}

func showProcess() {
	fmt.Printf("\n--------------------\n")
	fmt.Printf("|  %s  |  %s  |  %s  |", frame[0], frame[1], frame[2])
	fmt.Printf("\n--------------------\n")
	fmt.Printf("\nPage-fault = %d\n\n", fault)
	fmt.Printf("\nCommand > ")
}

func getCommand() string {
	reader := bufio.NewReader(os.Stdin)
	data, _ := reader.ReadString('\n')
	data = strings.Trim(data, "\n")
	return data
}

func command_create(p string) {
		if frame[0] == "-" && frame[1] == "-" && frame[2] == "-" {
			frame[0] = p
			fault++
		} else if frame[0] != "-" && frame[1] == "-" && frame[2] == "-" {
			frame[1] = p
			fault++
		} else if frame[0] != "-" && frame[1] != "-" && frame[2] == "-" {
			frame[2] = p
			fault++
		} else {
			if p != frame[0] && p != frame[1] &&  p != frame[2] {
				copy(frame[0:], frame[1:])
				frame[2] = p
				fault++
			} else if p == frame[2] {
				frame[2] = p
				//fmt.Printf("push")
			} else if p == frame[1] {
				frame[1] = frame[2]
				frame[2] = p
			} else {
				copy(frame[0:], frame[1:])
				frame[2] = p
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