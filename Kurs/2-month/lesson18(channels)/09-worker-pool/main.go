package main

import (
	"fmt"
	"sync"
	"time"
)

// Worker Pool - главный практический паттерн курса.
// Вместо горутины на каждую задачу запускаем N воркеров,
// которые разбирают задачи из общего канала.

type Job struct {
	ID   int
	Data string
}

type Result struct {
	JobID  int
	Output string
	Worker int
}

func worker(id int, jobs <-chan Job, results chan<- Result, wg *sync.WaitGroup) {
	defer wg.Done()
	for j := range jobs {
		fmt.Printf("Воркер %d взял задачу %d\n", id, j.ID)
		time.Sleep(200 * time.Millisecond) // имитация работы
		results <- Result{
			JobID:  j.ID,
			Output: "обработано: " + j.Data,
			Worker: id,
		}
	}
}

func main() {
	const numJobs = 9
	const numWorkers = 3

	jobs := make(chan Job, numJobs)
	results := make(chan Result, numJobs)
	var wg sync.WaitGroup

	// запускаем воркеров
	for w := 1; w <= numWorkers; w++ {
		wg.Add(1)
		go worker(w, jobs, results, &wg)
	}

	// раздаём задачи
	start := time.Now()
	for j := 1; j <= numJobs; j++ {
		jobs <- Job{ID: j, Data: fmt.Sprintf("данные-%d", j)}
	}
	close(jobs) // больше задач не будет

	// ждём воркеров и закрываем результаты
	go func() {
		wg.Wait()
		close(results)
	}()

	// собираем результаты
	count := 0
	for r := range results {
		count++
		fmt.Printf("  Результат: задача %d от воркера %d\n", r.JobID, r.Worker)
	}

	fmt.Printf("\nОбработано %d задач за %v (воркеров: %d)\n",
		count, time.Since(start).Round(time.Millisecond), numWorkers)

	// Поэкспериментируй с numWorkers (1, 3, 9) и покажи, как меняется время.
}
