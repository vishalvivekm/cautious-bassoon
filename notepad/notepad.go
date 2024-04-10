package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type notepad struct {
	Notes                           []string
	Response                        string
	Size                            int
	inpMsg, addMsg, delMsg, listFmt string
	errFul, errMpt, errCmd, errVal  string
	exitMsg, exitCmd                string
}

func newPad() (*notepad, *bufio.Scanner) {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter the maximum number of notes: ")
	scanner.Scan()
	size, _ := strconv.Atoi(scanner.Text())

	myPad := notepad{
		inpMsg:  "Enter a command and data: ",
		addMsg:  "[OK] The note was successfully created",
		delMsg:  "[OK] All notes were successfully deleted",
		listFmt: "[Info] %d: %s",
		errFul:  "[Error] Notepad is full",
		errMpt:  "[Info] Notepad is empty",
		errCmd:  "[Error] Unknown command",
		errVal:  "[Error] Missing note argument",
		exitMsg: "[Info] Bye!", exitCmd: "exit",
		Notes: newSlice(size), Size: size,
	}
	return &myPad, scanner
}

func newSlice(arrLimit int) []string {
	return make([]string, 0, arrLimit)
}

func add(pad *notepad, newContent string) ([]string, string) {
	if len(pad.Notes) == pad.Size {
		return pad.Notes, pad.errFul
	}
	if newContent == "" {
		return pad.Notes, pad.errVal
	}
	pad.Notes = append(pad.Notes, newContent)
	return pad.Notes, pad.addMsg
}

func list(pad *notepad, listed []string) string {
	if len(pad.Notes) == 0 {
		return pad.errMpt
	}
	for i, note := range pad.Notes {
		listed = append(listed, fmt.Sprintf(pad.listFmt, i+1, note))
	}
	return strings.Join(listed, "\n")
}

func clear(pad *notepad) ([]string, string) {
	pad.Notes = nil
	return pad.Notes, pad.delMsg
}

func run(pad *notepad, scanner *bufio.Scanner) {
	var command string
mainLoop:
	for {
		var content string
		fmt.Printf("%s", pad.inpMsg)
		scanner.Scan()

		command, content = func() (string, string) {
			entry := regexp.MustCompile(`\s+`).Split(scanner.Text(), 2)
			entry = append(entry, content)
			return entry[0], entry[1]
		}()

		switch command {
		case "create":
			pad.Notes, pad.Response = add(pad, content)
		case "list":
			pad.Response = list(pad, newSlice(pad.Size))
		case "clear":
			pad.Notes, pad.Response = clear(pad)
		case pad.exitCmd:
			break mainLoop
		default:
			fmt.Print(pad.errCmd)
		}
		fmt.Println(pad.Response)
	}

	fmt.Printf(pad.exitMsg)
}

func main() {
	lilPad, scanner := newPad()
	run(lilPad, scanner)
}
