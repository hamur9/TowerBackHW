package main

//Для реализации выбраны:
//Context + WaitGroup + Signals
//Преимущества:
//Context - уведомление всех горутин
//WaitGroup - гарантированное ожидание завершения
//Signal - корректная реакция на Ctrl+C

//Как работает:
//Ctrl+C → signal.Notify → cancel()
//Все горутины получают <-ctx.Done()
//WaitGroup дожидается wg.Done() от всех

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"syscall"
	"time"
)

type WeatherData struct {
	City        string
	Temperature int
	Humidity    int
	Condition   string
}

func main() {
	if len(os.Args) != 2 {
		log.Fatal("Неверное количество аргументов")
	}

	numWorkers, err := strconv.Atoi(os.Args[1])
	if err != nil || numWorkers <= 0 {
		log.Fatal("Неверное количество воркеров")
	}

	ctx, cancel := context.WithCancel(context.Background())
	dataCh := make(chan string)
	var wg sync.WaitGroup

	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go worker(ctx, i, dataCh, &wg)
	}

	wg.Add(1)
	go producer(ctx, dataCh, &wg)

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)

	<-sigCh
	fmt.Println("\nВоркеры останавливаются...")
	cancel()

	wg.Wait()
	fmt.Println("Все воркеры остановлены.")
}

func worker(ctx context.Context, id int, dataCh <-chan string, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		select {
		case <-ctx.Done():
			fmt.Printf("Воркер %d: завершение работы\n", id)
			return
		case data, ok := <-dataCh:
			if !ok {
				fmt.Printf("Воркер %d: канал закрыт\n", id)
				return
			}
			fmt.Printf("Воркер %d: %s\n", id, data)
		}
	}
}

func producer(ctx context.Context, dataCh chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()
	defer close(dataCh)

	cities := []string{"Москва", "Санкт-Петербург", "Новосибирск", "Екатеринбург", "Казань"}
	conditions := []string{"солнечно", "облачно", "дождь", "снег", "туман"}
	weather := make(map[string]WeatherData)

	for _, city := range cities {
		weather[city] = WeatherData{
			City:        city,
			Temperature: -20 + rand.Intn(45),
			Humidity:    20 + rand.Intn(80),
			Condition:   conditions[rand.Intn(len(conditions))],
		}
	}

	for {
		select {
		case <-ctx.Done():
			fmt.Println("Продюсер: завершение работы")
			return
		default:
			for city, data := range weather {
				newTemp := data.Temperature + (rand.Intn(5) - 2)
				newHumidity := data.Humidity + (rand.Intn(10) - 5)
				if newHumidity < 0 {
					newHumidity = 0
				}
				if newHumidity > 100 {
					newHumidity = 100
				}

				weather[city] = WeatherData{
					City:        city,
					Temperature: newTemp,
					Humidity:    newHumidity,
					Condition:   conditions[rand.Intn(len(conditions))],
				}
			}

			for _, data := range weather {
				weatherStr := fmt.Sprintf("Погода в %s: %d°C, влажность %d%%, %s",
					data.City, data.Temperature, data.Humidity, data.Condition)
				select {
				case dataCh <- weatherStr:
				case <-ctx.Done():
					return
				}
				time.Sleep(200 * time.Millisecond)
			}
			time.Sleep(2 * time.Second)
		}
	}
}
