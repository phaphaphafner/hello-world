package main

import (
	"fmt"
	"reflect"
)

var otherThing float64 = -32.34

func init() {
}

func main() {
	fmt.Println(doAnotherThing(otherThing))
}

func doAnotherThing(float float64) string {
	return fmt.Sprintf("%f - is the %x", float, reflect.TypeOf(float))
}
