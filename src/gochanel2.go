package main

import (
    "fmt"
    "time"
)

//параллельная обработка в рутинах и запись и чтение из канала
func main() {
    // создаем канал, в который будем отправлять сообщения
    msgCh1 := make(chan string)
    msgCh2 := make(chan string)
    
    go mysetter(msgCh1, "Hello")
    go mysetter(msgCh2, "World")

    s1, s2 := <-msgCh1, <-msgCh2

    fmt.Println(s1, s2)
    
    close(msgCh1)
    close(msgCh2)        
}

func mysetter(msgCh chan string, str string) {    
    //time.Sleep(3000 * time.Millisecond)
    msgCh <- str
    time.Sleep(3000 * time.Millisecond) //этот слип нам пофиг, так как в канал записали. То есть дождались
}