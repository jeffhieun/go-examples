package main

import (
	"fmt"
	"strings"
)

func main() {
	var cars = [4]string{"Volvo", "BMW", "Ford", "Mazda"}
	fmt.Printf("My output %s\n", strings.Join(cars[:], ","))

}
