package main

import (
	"birthdate/week"
	"fmt"
	"log"
)

func main() {
	fmt.Print("Enter year (eg., 2000), month(1-12), and date (1-31):")

	var y, m, d int
	if _, err := fmt.Scan(&y, &m, &d); err != nil {
		log.Fatalln("Scan for y, m, and d failed:", err)
	}
	weekday := week.Weekday(y, m, d)
	fmt.Println("Weekday =", weekday)
}
