package utils

import "fmt"

const (
	Black   = 30
	Red     = 31
	Green   = 32
	Yellow  = 33
	Blue    = 34
	Magenta = 35
	Cyan    = 36
	White   = 37
)

const (
	Bold      = 1
	Underline = 4
	Blinking  = 5
	Reverse   = 7
)

func Print(text string, color int) {
	fmt.Printf("\033[%dm%s\033[0m\n", color, text)
}

func PrintStyled(text string, color int, style int) {
	fmt.Printf("\033[%d;%dm%s\033[0m\n", style, color, text)
}

func Error(text string) {
	PrintStyled(text, Red, Bold)
}

func Success(text string) {
	PrintStyled(text, Green, Bold)
}

func Info(text string) {
	Print(text, Blue)
}

func Loading(text string) {
	PrintStyled(text, Yellow, Blinking)
}

func Clear() {
	fmt.Print("\033[H\033[2J")
}
