package main

import (
    "context"
    "fmt"
    "time"
)

func main() {
    ctx := context.Background()
    ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
    defer cancel()
    //time.Sleep(3 * time.Second) это добавляет общее время, потому контекст попадет более поздний
    //то есть раскоментировать - будет Canceled 
    
    go slowTask(ctx)
    
    //give some time for the slowTask goroutine to conclude
    time.Sleep(4 * time.Second)
}

func slowTask(ctx context.Context) {
    select {
        case <-time.After(2 * time.Second):
            fmt.Println("Прога отработала.")
        case <-ctx.Done():
            fmt.Println("Canceled:", ctx.Err())
    }
}