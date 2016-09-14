package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	loc, _ := time.LoadLocation("US/Eastern")
	z := t.In(loc)
	fmt.Println(z.Format(time.RubyDate))
}
