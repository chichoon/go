package main

import "fmt"

func main() {
	fmt.Println("함수의 시작이야")

	for i := 0; i < 10; i++ {
		defer fmt.Println("치춘짱 나이스 *", i)
	}

	fmt.Println("함수의 끝이고")
}
