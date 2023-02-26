package main

import "fmt"

func main() {

	var arr1 = []string{"a", "b", "c", "d", "e"}

	fmt.Println(arr1)
	fmt.Println(arr1[2:])
	//fmt.Println(arr1[2:]...)
	fmt.Println(append([]string{"insert"}, arr1[1:]...))

}
