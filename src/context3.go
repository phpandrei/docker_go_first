package main

import (
    "context"
    "fmt"
    "time"
)

type contextkey string

func main() {
    ctx := context.Background()
    ctx = context.WithValue(ctx, contextkey("AuthToken"), "eqwe-4545e-445454e-fg2")

    go processSensitiveData(ctx)

    time.Sleep(2 * time.Second)
}

func processSensitiveData(ctx context.Context) {
    //просто передача ключ-значение
    authToken, ok := ctx.Value(contextkey("AuthToken")).(string)

    if !ok {        
        fmt.Println("Токена нет!")

        return
    }

    fmt.Println("Токен: %s", authToken)
}