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

const monitoring = 3
const delay = 5

func main() {

	showIntroduction()
	for {
		showMenu()
		command := readCommand()

		switch command {
		case 1:
			startMonitoring()
		case 2:
			printLogs()
		case 0:
			fmt.Println("Exiting program")
			os.Exit(0)
		default:
			fmt.Println("Invalid command")
			os.Exit(-1)
		}
	}
}

func showIntroduction() {
	name := "Ricardo"
	version := 1.1
	fmt.Println("Hello, Mr/Ms.", name)
	fmt.Println("Program version", version)
}

func showMenu() {
	fmt.Println("1 - Start Monitoring")
	fmt.Println("2 - Show Logs")
	fmt.Println("0 - Exit Program")
}

func readCommand() int {
	var readedCommand int
	fmt.Scan(&readedCommand)
	fmt.Println("The selected menu was", readedCommand)
	fmt.Println("")

	return readedCommand
}

func startMonitoring() {
	fmt.Println("Monitoring...")
	//sites := []string{"https://www.alura.com.br", "https://g1.globo.com", "https://www.cruzeiro.com.br"}
	sites := readSitesFromFile()

	for i := 0; i < monitoring; i++ {
		fmt.Println("Monitoring number:", i+1, "from:", monitoring)
		fmt.Println("")
		for i, site := range sites {
			fmt.Println("Testing site", i, ":", site)
			testSite(site)
		}
		time.Sleep(delay * time.Second)
		fmt.Println("")
	}

	fmt.Println("")

}

func testSite(site string) {
	response, err := http.Get(site)

	if err != nil {
		fmt.Println("Occurred a error in get site:", err)
	}

	if response.StatusCode == 200 {
		fmt.Println("Website:", site, "is running! StatusCode:", response.StatusCode)
		logRegister(site, true)
		fmt.Println("")
	} else {
		fmt.Println("Website:", site, "is down! StatusCode:", response.StatusCode)
		logRegister(site, false)
		fmt.Println("")
	}
}

func readSitesFromFile() []string {
	var sites []string

	file, err := os.Open("sites.txt")

	if err != nil {
		fmt.Println("Occurred a error in open file:", err)
	}

	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		line = strings.TrimSpace(line)

		sites = append(sites, line)

		if err == io.EOF {
			break
		}
	}
	file.Close()

	return sites
}

func logRegister(site string, status bool) {
	logFile, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println(err)
	}

	logFile.WriteString(time.Now().Format("02/01/2006 15:04:05") + " - " + site + " - online: " + strconv.FormatBool(status) + "\n")

	fmt.Println(logFile)
	logFile.Close()
}

func printLogs() {

	file, err := ioutil.ReadFile("log.txt")

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(file))
}
