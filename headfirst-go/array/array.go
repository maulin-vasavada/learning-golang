package array

import "fmt"

func HelloArrays() {
	fmt.Println("arrays are awesome!")
}

func Print() {
	notes := [7]string{"do", "re", "me", "fa", "so", "la", "ti"}

	for index, note := range notes {
		fmt.Println("Note at index", index, "is", note)
	}
}

func Average(numbers ...float64) float64 {
	var sum float64 = 0
	for _, number := range numbers {
		sum += number
	}
	sampleCount := float64(len(numbers))
	average := sum / sampleCount
	fmt.Printf("SampleCount is %d, Sum is %0.2f, Avg is %0.2f\n", len(numbers), sum, average)
	return average
}
