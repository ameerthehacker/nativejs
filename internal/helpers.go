package internal

import (
	"fmt"
	"github.com/logrusorgru/aurora"
	"log"
	"os"
)

// func Must panics with printing error in red color to stderr
func Must(err error) {
	if err != nil {
		log.Panicf("njs: %s", aurora.Red(err.Error()))
	}
}

// func PrintError prints the error string as red color to the stderr
func PrintError(err error) {
	if err != nil {
		errString := fmt.Sprintf("njs: %s", err.Error())

		_, _ = fmt.Fprintln(os.Stderr, aurora.Red(errString))
	}
}
