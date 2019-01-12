package main

import (
	"fmt"
	"time"
)

func main() {
	hours, _ := time.ParseDuration("10h")
	complex, _ := time.ParseDuration("1h10m10s")

	fmt.Println(hours)
	fmt.Println(complex)
	fmt.Printf("there are %.0f seconds in %v\n", complex.Seconds(), complex)

	h, _ := time.ParseDuration("4h30m")
	fmt.Printf("I've got %.1f hours of work left.\n", h.Hours())

	m, _ := time.ParseDuration("1h30m")
	fmt.Printf("The movie is %.0f minutes long.\n", m.Minutes())
	ns, _ := time.ParseDuration("1000ns")
	fmt.Printf("one microsecond has %d nanoseconds.", ns.Nanoseconds())
}
