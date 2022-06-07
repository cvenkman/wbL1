package main

// Реализовать все возможные способы остановки выполнения горутины.

import (
	"context"
	"fmt"
	"time"
)

func main() {
	viaReturn()
	viaContext()
	viaChannel()
}

// выходим из горутины при return
func viaReturn() {
	go func() {
		return
	}()
	fmt.Println("done return") // это выполнится
}

// cancel() говорит горутинам заканчиваться
func viaContext() {
	ctx, cancel := context.WithCancel(context.Background())

	go func(ctx context.Context) {
		select {
		case <-ctx.Done():
			fmt.Println("done context") // сюда придет значение при вызове cancel()
			return
		default:
			time.Sleep(500 * time.Millisecond)
		}
	}(ctx)

	cancel()
	time.Sleep(1 * time.Second)
}

// горутина блокируется при чтении из канала, ждет, пока в него что-то отправят
func viaChannel() {
	ch := make(chan bool)
	go func() {
		time.Sleep(1 * time.Second)
		ch <- true
	}()
	<-ch
	fmt.Println("done channel")
}
