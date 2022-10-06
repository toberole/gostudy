package demo

import "fmt"

var n = 1

const (
	N1 = iota
	N2
)

func Demo1(f string, args ...interface{}) {
	fmt.Printf("Demo1 ......")
	fmt.Printf(f, args)
}
