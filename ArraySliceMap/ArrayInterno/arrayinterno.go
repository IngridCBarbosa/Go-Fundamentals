package main

import "fmt"

func main() {
	slice1 := make([]int, 10, 20)
	slice2 := append(slice1, 1, 2, 3)

	fmt.Println("Slice 1:", slice1)
	fmt.Println("Slice 2:", slice2)

	slice1[0] = 7

	fmt.Println("Slice 1:", slice1)
	fmt.Println("Slice 2:", slice2)

}
