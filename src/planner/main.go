package main

import (
	"fmt"
	"test/src/dates"
)

func main() {
	days := 3
	fmt.Println("Your appointment is in ", days, "days")
	fmt.Printf("with a follow-up in", days+dates.DaysInWeek, "days")
}
