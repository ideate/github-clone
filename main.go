package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"time"
)

type (
	configuration struct {
		GITHUB_KEY string
	}
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	// OPEN JSON CONFIG FILE AND DECODE
	configFile, err := os.Open("config.json")
	check(err)
	fmt.Println("Config file opened.")

	decoder := json.NewDecoder(configFile)
	config := configuration{}
	err = decoder.Decode(&config)
	check(err)
	fmt.Println("Config file decoded.")

	// MAKE OUTPUT DIRECTORY IF IT DOES NOT EXIST
	if _, err := os.Stat("output"); err != nil {
		err = os.Mkdir("output", 0777)
		check(err)
		fmt.Println("Output directory created.")
	} else {
		check(err)
	}

	// OPEN URL INPUT FILE AND READ
	file, err := os.Open("input/url")
	check(err)
	defer file.Close()
	fmt.Println("Input file opened.")

	// READ FILE LINE BY LINE
	numLines := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		time.Sleep(1000 * time.Millisecond)
		numLines++
		outputDir := "output/" + strconv.Itoa(numLines)

		cmd := exec.Command("git", "clone", scanner.Text(), outputDir)
		err := cmd.Run()
		check(err)

		fmt.Printf("Git clone #%v successful (%v).", numLines, scanner.Text())
		fmt.Printf("\n")
	}

	err = scanner.Err()
	check(err)

}
