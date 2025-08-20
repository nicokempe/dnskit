package output

import (
	"fmt"

	"github.com/fatih/color"
)

var (
	// infoColorFormatter formats informational messages in blue.
	infoColorFormatter = color.New(color.FgBlue).SprintFunc()
	// successColorFormatter formats success messages in green.
	successColorFormatter = color.New(color.FgGreen).SprintFunc()
	// errorColorFormatter formats error messages in red.
	errorColorFormatter = color.New(color.FgRed).SprintFunc()
)

// Info prints an informational message.
func Info(message string) {
	fmt.Println(infoColorFormatter(message))
}

// Success prints a success message.
func Success(message string) {
	fmt.Println(successColorFormatter(message))
}

// Error prints an error message.
func Error(message string) {
	fmt.Println(errorColorFormatter(message))
}
