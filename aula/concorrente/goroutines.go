package main

import (
	"fmt"
	"time"
)

func trabalho(nome string) {
	fmt.Println(nome, " iniciou")
	fmt.Println(2 * time.Second)
	fmt.Println(nome, " finalizou")
}

func main2() {
	go trabalho("Goroutine 1")
	go trabalho("Goroutine 2")

	time.Sleep(3 * time.Second)
	fmt.Println("todos finalizaram!")
}
