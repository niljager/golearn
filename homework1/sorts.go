package main

import (
	"fmt"
	"math"
	"math/rand"
	"sort"
)

func main() {
    numbers := make([]int, 0, 10)
    for i := 1; i <= cap(numbers); i++ {
        numbers = append(numbers, rand.Intn(10))
    }

	numbers1 := make([]int, 0, 10)
    for i := 1; i <= cap(numbers1); i++ {
        numbers1 = append(numbers1, rand.Intn(10))
    }

	numbers2 := make([]int, 0, 10)
    for i := 1; i <= cap(numbers2); i++ {
        numbers2 = append(numbers2, rand.Intn(10))
    }

	numbers3 := make([]int, 0, 10)
    for i := 1; i <= cap(numbers3); i++ {
        numbers3 = append(numbers3, rand.Intn(10))
    }

	bubbleSort(numbers)
    fmt.Println(numbers, sort.IntsAreSorted(numbers))

	selectionSort(numbers1)
	fmt.Println(numbers1, sort.IntsAreSorted(numbers1))

	insertionSort(numbers2)
	fmt.Println(numbers2, sort.IntsAreSorted(numbers2))

	quickSort(numbers3, 0, len(numbers3) - 1)
	fmt.Println(numbers3, sort.IntsAreSorted(numbers3))
}

func bubbleSort(numbers []int) {
	for i := 0; i < len(numbers) - 1; i++ {
        for j := 0; j < len(numbers) - i - 1; j++ {
            if numbers[j] > numbers[j+1] {
                tmp_val := numbers[j]
                numbers[j] = numbers[j+1]
                numbers[j+1] = tmp_val
            }
        }
    }
}

func selectionSort(numbers1 []int) {
	for i := 0; i < len(numbers1) - 1; i++ {
        min_index := i;
        for j := i + 1; j< len(numbers1); j++ {
            if(numbers1[min_index] > numbers1[j]) {
                min_index = j
            }
        }
        if(min_index != i) {
            tmp_val := numbers1[i]
            numbers1[i] = numbers1[min_index]
            numbers1[min_index] = tmp_val
        }
    }
}

func insertionSort(numbers2 []int) {
	for i := 1; i < len(numbers2); i++ {
        sorted := i - 1
        for sorted >= 0 {
            if numbers2[sorted] > numbers2[sorted + 1] {
                temp_val := numbers2[sorted]
                numbers2[sorted] = numbers2[sorted + 1]
                numbers2[sorted + 1] = temp_val
                sorted--
            } else {
                break
            }
        }
    }
}

func quickSort(numbers3 []int, start , end int) {
	if start < end {
        pivot := partOfNumbers(numbers3, start, end)
        quickSort(numbers3, start, pivot - 1)
        quickSort(numbers3, pivot, end)
    }
}

func partOfNumbers(numbers []int, start, end int) int {
    pivot := int(math.Floor(float64((start + end) / 2)))

    for start <= end {
        for numbers[start] < numbers[pivot] {
            start++
        }

        for numbers[end] > numbers[pivot] {
            end--
        }

        if(start <= end) {
            temp_val := numbers[start]
            numbers[start] = numbers[end]
            numbers[end] = temp_val
            start++
            end--
        }
    }

    return start
}