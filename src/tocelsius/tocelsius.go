package main

import (
	"fmt"
	"log"
	"test/src/keyboard"
)

func main() {
	fmt.Println("Enter a temperature in Fahrenheit")
	fahrenheit, err := keyboard.GetFloat()
	if err != nil {
		log.Fatal(err)
	}
	celsius := (fahrenheit - 32) * 5 / 9
	fmt.Println("%0.2f degrees Celsius\n", celsius)
}
