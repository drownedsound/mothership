package main

import (
	"fmt"
	"log"
	"strings"
)

func printHeader() {
	fmt.Println(strings.Repeat(filler, lineLength))
	fmt.Println(strings.ToUpper("👾 👾 👾 Start Mothership 👾 👾 👾"))
	fmt.Println(strings.Repeat(filler, lineLength), "\n")
}

func printFooter() {
	fmt.Println(strings.Repeat(filler, lineLength))
	fmt.Println(strings.ToUpper("👾 👾 👾 End Mothership 👾 👾 👾"))
	fmt.Println(strings.Repeat(filler, lineLength))
}

func logToConsole(msg string) {
	log.Println(msg)
}
