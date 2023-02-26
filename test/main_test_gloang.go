package main

import "fmt"

func main() {

	fmt.Println("\n=======数组slice======")
	var arr1 = []string{"a", "b", "c", "d", "e"}

	fmt.Println(arr1)
	fmt.Println(arr1[2:])
	//fmt.Println(arr1[2:]...)
	fmt.Println(append([]string{"insert"}, arr1[1:]...))

	fmt.Println("\n=======指针======")

	var a int = 1
	var b *int = &a
	var c **int = &b
	var x int = *b
	//https://juejin.cn/post/6844903802210877454
	//*和& 可以互相抵消 但是注意
	//【 *& 】可以抵消掉，但【 &* 】是不可以抵消的
	//a和 *&a   是一样的,都是a的值，值为1 (因为*&互相抵消掉了)
	//a和 *&*&*&*&a是一样的，都是1     (因为4个*&互相抵消掉了)
	fmt.Println("a = ", a)
	fmt.Println("&a = ", &a)
	fmt.Println("*&a = ", *&a)
	fmt.Println("b = ", b)
	fmt.Println("&b = ", &b)
	fmt.Println("*&b = ", *&b)
	fmt.Println("*b = ", *b)
	fmt.Println("c = ", c)
	fmt.Println("*c = ", *c)
	fmt.Println("&c = ", &c)
	fmt.Println("*&c = ", *&c)
	fmt.Println("**c = ", **c)
	fmt.Println("***&*&*&*&c = ", ***&*&*&*&*&c)
	fmt.Println("x = ", x)

}
