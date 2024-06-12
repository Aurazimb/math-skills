package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
)

func main() {
	file, err := os.Open("data.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	data := []float64{}
	num := 0.0
	for scanner.Scan() {
		line := scanner.Text()
		for _, char := range line {
			if char >= '0' && char <= '9' {
				num = num * 10
				num += float64(char - 48)
			}
		}
		data = append(data, num)
		num = 0
	}
	average := int(math.Round(Average(data)))
	fmt.Printf("Average: %d", average)
	fmt.Println()
	median := int(math.Round(Median(data)))
	fmt.Printf("Median: %d", median)
	fmt.Println()
	variance := int(math.Round(Variance(data)))
	fmt.Printf("Variance: %d", variance)
	fmt.Println()
	standart_deviation := int(math.Round(StandDev(data)))
	fmt.Printf("Standard Deviation: %d", standart_deviation)
	fmt.Println()
}

func Average(slice []float64) float64 {
	sum := 0.0
	for _, num := range slice {
		sum += num
	}
	x := sum / float64(len(slice))
	return x
}

func Median(slice []float64) float64 {
	sort.Float64s(slice)
	median := float64(len(slice) / 2.0)
	if len(slice)%2 == 0 {
		median = (slice[len(slice)/2-1] + slice[len(slice)/2]) / 2.0
	}
	return float64(median)
}

func Variance(slice []float64) float64 {
	av := Average(slice)
	var variance float64
	for _, el := range slice {
		dif := el - av
		variance += dif * dif
	}

	return variance / float64(len(slice))
}

func StandDev(slice []float64) float64 {
	variance := Variance(slice)
	Standdev := math.Sqrt(variance)
	return Standdev
}
