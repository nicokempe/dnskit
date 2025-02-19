package output

import (
	"fmt"
	"github.com/fatih/color"
)

var (
	InfoColor    = color.New(color.FgBlue).SprintFunc()
	SuccessColor = color.New(color.FgGreen).SprintFunc()
	ErrorColor   = color.New(color.FgRed).SprintFunc()
)

func Info(msg string) {
	fmt.Println(InfoColor(msg))
}

func Success(msg string) {
	fmt.Println(SuccessColor(msg))
}

func Error(msg string) {
	fmt.Println(ErrorColor(msg))
}
