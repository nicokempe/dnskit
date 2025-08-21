package output

import (
	"fmt"

	"github.com/fatih/color"
)

var (
	infoColorFormatter    = color.New(color.FgBlue).SprintFunc()
	successColorFormatter = color.New(color.FgGreen).SprintFunc()
	errorColorFormatter   = color.New(color.FgRed).SprintFunc()
)

func Info(message string) {
	fmt.Println(infoColorFormatter(message))
}

func Success(message string) {
	fmt.Println(successColorFormatter(message))
}

func Error(message string) {
	fmt.Println(errorColorFormatter(message))
}
