package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

func main() {

	showIntro()
	for {
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
	websites := []string{"http://random-status-code.herokuapp.com/", "http://www.alura.com.br", "http://www.caellum.com.br"}

	for i := 0; i < 5; i++ {
		for i, website := range websites {
			fmt.Println("Testing website", i, ":", website)
			websiteTest(website)
		}
		time.Sleep(10 * time.Second)
		fmt.Println("")
	}
	fmt.Println("")
}

func websiteTest(website string) {
	reply, _ := http.Get(website)

	if reply.StatusCode == 200 {
		fmt.Println("The website:", website, "loaded successfully!")
	} else {
		fmt.Println("The website:", website, "is having problems. Status code:", reply.StatusCode)
	}
}
