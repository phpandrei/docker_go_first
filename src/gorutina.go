package main

import (
	"time"
	"fmt"
	"os"
	"sync"
)

var Wg sync.WaitGroup

func MyRutina() {
	time.Sleep(time.Millisecond * 1000)
	println("Hello gorutina")

	text := "Hello Andrei!\r\n"
    //file, err := os.Open("hello.txt")
	file, err := os.OpenFile("hello.txt", os.O_APPEND|os.O_WRONLY, 0777)
     
	//очень интересная ситуация. Нет файла в 2х горутинах, потому гонка. потому накатывает 2 раза создание.
	//Надо инициировать вначале, скорее всего
    if err != nil{
        fmt.Println("Unable to create file:", err) 

        file, _ = os.Create("hello.txt")
		os.Chmod("hello.txt", 0777)
    }

    defer file.Close() 

    file.WriteString(text)
     
    fmt.Println("Done.")

	Wg.Done()
}

func main() {
	Wg.Add(2)
	go MyRutina()
	go MyRutina()

	println("Hello, World!")

	time.Sleep(time.Millisecond * 500)
	Wg.Wait()
}
