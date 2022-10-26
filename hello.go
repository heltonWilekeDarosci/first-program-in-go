package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {

	showIntro()
	showMenu()
	command := seeCommand()

	switch command {
	case 1:
		startMonitoring()
	case 2:
		fmt.Println("Exhibiting Logs...")
	case 0:
		fmt.Println("Exiting...")
		os.Exit(0) // os.Exit(0) To tell to exit because the result was a success, nothing went wrong <<---
	default:
		fmt.Println("Command unknown")
		os.Exit(-1) // os.Exit(-1) To tell something went wrong <<---
	}
}

func showIntro() {
	name := "Wileke"
	version := 1.1
	fmt.Println("Hello Mr.", name)
	fmt.Println("This program is on version", version)
}

func showMenu() {
	fmt.Println("1- Start monitoring")
	fmt.Println("2- Exhibit Logs")
	fmt.Println("0- Exit")
}

func seeCommand() int {
	var commandRead int
	fmt.Scan(&commandRead)

	return commandRead
}

func startMonitoring() {
	fmt.Println("Monitoring...")
	site := "http://random-status-code.herokuapp.com/"
	reply, _ := http.Get(site)

	if reply.StatusCode == 200 {
		fmt.Println("The website:", site, "loaded successfully!")
	} else {
		fmt.Println("The website:", site, "is having problems. Status code:", reply.StatusCode)
	}
}
