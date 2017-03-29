package main

import (
	"fmt"
	"math/rand"
	"strconv"
)

type graph struct {
	Next *graph
	secret float64
	pos string
}

func get_Random() float64 {
	return rand.Float64()
}

func xor(a bool, b bool) bool {
	return a != b
}

func Un_xor(a bool, b bool) bool {
	x := xor(a, b)
	return xor(x, true)	//Xor by 1 gives the opposite of the original number
}


func FloatToString(input_num float64) string {
	return strconv.FormatFloat(input_num, 'f', 2, 64)
}



func main() {
	ACE := true
	MACE := false
	fmt.Printf(FloatToString( get_Random() ))
	if xor(ACE, MACE) {
		fmt.Printf("\nExclusive OR TRUE\n")
	} else {
		fmt.Printf("\nExclusive OR FALSE\n")
	}

}
