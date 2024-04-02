package files

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func GetFloats(fileName string) ([]float64, error) {
	numbers := make([]float64, 3)
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		//fmt.Println(scanner.Text())
		number, err := strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			return numbers, err
		}
		numbers = append(numbers, number)
	}
	err = file.Close()
	if err != nil {
		return numbers, err
	}
	if scanner.Err() != nil {
		return numbers, err
	}
	return numbers, err
}
