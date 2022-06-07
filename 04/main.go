package main

// Реализовать постоянную запись данных в канал (главный поток). Реализовать
// набор из N воркеров, которые читают произвольные данные из канала и выводят в stdout.
// Необходима возможность выбора количества воркеров при старте.

// Программа должна завершаться по нажатию Ctrl+C. Выбрать и обосновать способ завершения работы всех воркеров.

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"context"
)

func writer(ctx context.Context, ch chan int) {
	for i := 0; ; i++ {
		select {
		default:
			ch <- i
		case <-ctx.Done():
			return
		}
	}
}

func worker(ctx context.Context, ch chan int) {
	for {
		select {
		default:
			fmt.Println(<-ch)
		case <-ctx.Done():
			return
		}
	}
}

func main() {
	ch := make(chan int)
	// контекс для остановки горутин
	ctx, cancel := context.WithCancel(context.Background())

	// запускаем воркеров
	for i := 0; i < 10; i++ {
		go worker(ctx, ch)
	}
	// запускаем канал с данными для воркеров
	go writer(ctx, ch)

	stop := make(chan os.Signal)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	select {
	case <-stop:
		// cancel() когда нажали SIGINT
		cancel()
		fmt.Println("До свидания.")
	}
}