package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Print("이름 입력 : ")
	in := bufio.NewReader(os.Stdin)
	name, err := in.ReadString('\n')
	// Shadowing problem
	// var float32 float32 = 9.1
	// fmt.Println(float32)
	// var number float32
	// fmt.Println(number)

	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println(name)
	}
	// var fmt float32 = 9.1
	// fmt.Println(float32)
}
