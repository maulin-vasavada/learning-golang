package workerpool

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("worker", id, "started  job", j)
		time.Sleep(time.Second)
		fmt.Println("worker", id, "finished job", j)
		results <- j * 2
	}
}

func WorkerTest() {

	const numJobs = 10
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs)

	for a := 1; a <= numJobs; a++ {
		fmt.Println(<-results)
	}
}

func worker2(id int) {
	fmt.Printf("Worker %d starting\n", id)

	time.Sleep(time.Second)
	fmt.Printf("Worker %d done\n", id)
}

func WaitGroupTest() {

	var wg sync.WaitGroup

	/*
		wg.Add(1)
		go func() {
			defer wg.Done()
			worker2(1)
		}()

		wg.Add(1)
		go func() {
			defer wg.Done()
			worker2(2)
		}()

		wg.Add(1)
		go func() {
			defer wg.Done()
			worker2(3)
		}()
	*/
	/*
		for i := 1; i <= 5; i++ {
			wg.Add(1)

			go func() {
				defer wg.Done()
				worker2(i)
			}()
		}
	*/

	j := 1
	for i := 1; i <= 5; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()
			worker2(j)
			j++
		}()
	}

	wg.Wait()

}

func runner1(wg *sync.WaitGroup) {
	defer wg.Done() // This decreases counter by 1
	fmt.Print("\nI am first runner")

}

func runner2(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Print("\nI am second runner")
}

func runner3(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Print("\nI am third runner")
}

func WaitGroupTest2() {
	wg := new(sync.WaitGroup)
	wg.Add(2)

	// We are increasing the counter by 2
	// because we have 2 goroutines
	go runner1(wg)
	go runner2(wg)
	// Add below line to see the fun, don't change the Add(counter) to 3
	go runner3(wg)

	// This Blocks the execution
	// until its counter become 0
	wg.Wait()
}
