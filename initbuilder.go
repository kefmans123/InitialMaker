package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
	"time"
)

var name string

func init() {
	fmt.Println("Welcome to my perfected InitialMaker10.")
	fmt.Println()
	fmt.Println("In this InitialMaker we have 2 types:")
	time.Sleep(3 * time.Second)
	fmt.Println("The first one: EU Standard")
	fmt.Println("The second one: International")
	fmt.Println()
	time.Sleep(2 * time.Second)
}
func main() {
	fmt.Print("Which one do you choose? Please typ [1] or [2]: ")
	reader := bufio.NewReader(os.Stdin)

	// Readstring will block until the delimiter is entered
	input, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println("BeepBoop! Seems like some water went into my circuits. Please try again.", err)
		return
	}
	// Remove the delimiter from the string
	input = strings.TrimSuffix(input, "\n")
	input = strings.TrimSuffix(input, "\r")
	switch input {
	case "1":
		name = eu()
	case "2":
		name = International()
	case "close":
		fmt.Println("Closing the program. gOoDbYee.. .. .. Humann.. ^.^")
		return
	default:
		fmt.Println("")
		fmt.Println("BeeBoop! Please only use [1] or [2]. Please try again.")
		fmt.Println("")
	}
	if name == "" {
		return
	}

	fmt.Println("Your initial name: ", name)
	time.Sleep(2 * time.Second)
	fmt.Println("Closing the program. gOoDbYee.. .. .. Humann.. ^.^")
	time.Sleep(3 * time.Second)
}

func eu() string {
	var initName string

	fmt.Print("Please enter your first name(s): ")
	readerFirstname := bufio.NewReader(os.Stdin)

	// Readstring will block until the delimiter has been deliverd
	inputFirstname, err := readerFirstname.ReadString('\n')

	fmt.Print("Please enter your preposition(s), leave it blank if it does not apply to you: ")
	readerPrepositions := bufio.NewReader(os.Stdin)

	// Readstring will block until the delimiter has been deliverd
	inputPrepositions, err := readerPrepositions.ReadString('\n')

	fmt.Print("Please enter your last name: ")
	readerLastname := bufio.NewReader(os.Stdin)

	// Readstring will block until the delimiter has been deliverd
	inputLastname, err := readerLastname.ReadString('\n')

	if err != nil {
		fmt.Println("BeepBoop! Seems like some water went into my circuits. Please try again.", err)
		return ""
	}
	// Remove the delimiter from the string
	inputFirstname = strings.TrimSuffix(inputFirstname, "\n")
	inputFirstname = strings.TrimSuffix(inputFirstname, "\r")

	inputPrepositions = strings.TrimSuffix(inputPrepositions, "\n")
	inputPrepositions = strings.TrimSuffix(inputPrepositions, "\r")

	inputLastname = strings.TrimSuffix(inputLastname, "\n")
	inputLastname = strings.TrimSuffix(inputLastname, "\r")

	FirstnameArray := strings.Split(inputFirstname, " ")

	for _, c := range FirstnameArray {
		c = strings.ToUpper(c[0:1])
		initName += c + "."
	}
	count := len(initName) - 1
	initName = initName[0:count]

	if inputPrepositions == "" {
		initName += " " + inputLastname

	} else {
		initName += " " + inputPrepositions + " " + inputLastname
	}
	return initName
}

func International() string {
	var initName string
	fmt.Print("Please enter your fullname: ")

	reader := bufio.NewReader(os.Stdin)

	// Readstring will block until the delimiter is entered
	input, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println("BeepBoop! Seems like some water went into my circuits. Please try again.", err)
		return ""
	}
	// Remove the delimiter from the string
	input = strings.TrimSuffix(input, "\n")
	input = strings.TrimSuffix(input, "\r")

	inputArray := strings.Split(input, " ")

	for _, c := range inputArray {
		matched, err := regexp.Match(`^(?i)(van|der|de|den)`, []byte(c))

		if err != nil {
			fmt.Println("BeepBoop! Seems like some water went into my circuits. Please try again.", err)
			return ""
		}

		if !matched {
			c = strings.ToUpper(c[0:1])
			initName += c + "."
		}
	}

	count := len(initName) - 1
	initName = initName[0:count]

	return initName
}
