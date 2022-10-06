package cgo

/*
#include "cgoTest1.h"
*/
import "C"

import "fmt"

func CGoTest1_1() {
	var sum = C.add(1, 2)
	fmt.Println(sum)
}
