package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

const monitorings = 3
const delay = 7

func main() {

	showIntro()
	for {
		showMenu()
		command := seeCommand()

		switch command {
		case 1:
			startMonitoring()
		case 2:
			printLogs()
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
	fmt.Println("")

	return commandRead
}

func startMonitoring() {
	fmt.Println("Monitoring...")

	websites := readWebsiteFile()

	for i := 0; i < monitorings; i++ {
		for i, website := range websites {
			fmt.Println("Testing website", i, ":", website)
			websiteTest(website)
		}
		time.Sleep(delay * time.Second)
		fmt.Println("")
	}
	fmt.Println("")
}

func websiteTest(website string) {
	reply, err := http.Get(website)

	if err != nil {
		fmt.Println(err)
	}

	if reply.StatusCode == 200 {
		fmt.Println("The website:", website, "loaded successfully!")
		logRegister(website, true)
	} else {
		fmt.Println("The website:", website, "is having problems. Status code:", reply.StatusCode)
		logRegister(website, false)
	}
}

func readWebsiteFile() []string {

	var websites []string

	file, err := os.Open("websites.txt")

	if err != nil {
		fmt.Println(err)
	}

	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		line = strings.TrimSpace(line)
		websites = append(websites, line)

		if err == io.EOF {
			break
		}
	}

	file.Close()

	return websites
}

func logRegister(website string, status bool) {

	file, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println(err)
	}

	file.WriteString(time.Now().Format("02/01/2006 15:04:05") + " - " + website + " - online: " + strconv.FormatBool(status) + "\n")

	file.Close()
}

func printLogs() {
	fmt.Println("Exhibiting Logs...")
	fmt.Println("")
	file, err := ioutil.ReadFile("log.txt")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(file))
	fmt.Println("")
}
