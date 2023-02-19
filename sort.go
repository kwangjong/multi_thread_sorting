package main

import (
	"flag"
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

var MAX_THREAD int = runtime.NumCPU() //# of thread is limited to # of cores

func quickSort(arr *[]int32, low, high int) {
	if high-low <= 15 {
		insertionSort(arr, low, high)
		return
	}
	pivotIndex := medianOfThree(arr, low, high)
	pivot := (*arr)[pivotIndex]
	(*arr)[pivotIndex], (*arr)[high] = (*arr)[high], (*arr)[pivotIndex]

	i := low
	j := high - 1
	for {
		for i < high && (*arr)[i] <= pivot {
			i++
		}
		for j >= low && (*arr)[j] > pivot {
			j--
		}
		if i >= j {
			break
		}
		(*arr)[i], (*arr)[j] = (*arr)[j], (*arr)[i]
	}
	(*arr)[i], (*arr)[high] = (*arr)[high], (*arr)[i]
	quickSort(arr, low, i-1)
	quickSort(arr, i+1, high)
}

func insertionSort(arr *[]int32, low, high int) {
	for i := low + 1; i <= high; i++ {
		key := (*arr)[i]
		j := i - 1
		for j >= low && (*arr)[j] > key {
			(*arr)[j+1] = (*arr)[j]
			j--
		}
		(*arr)[j+1] = key
	}
}

func medianOfThree(arr *[]int32, low, high int) int {
	mid := (low + high) / 2
	if (*arr)[low] > (*arr)[mid] {
		(*arr)[low], (*arr)[mid] = (*arr)[mid], (*arr)[low]
	}
	if (*arr)[low] > (*arr)[high] {
		(*arr)[low], (*arr)[high] = (*arr)[high], (*arr)[low]
	}
	if (*arr)[mid] > (*arr)[high] {
		(*arr)[mid], (*arr)[high] = (*arr)[high], (*arr)[mid]
	}
	return mid
}

func quickSortPar(arr *[]int32, low, high int, wait *sync.WaitGroup, num_thread int) {
	defer wait.Done()

	if high-low <= 15 {
		insertionSort(arr, low, high)
		return
	}
	pivotIndex := medianOfThree(arr, low, high)
	pivot := (*arr)[pivotIndex]
	(*arr)[pivotIndex], (*arr)[high] = (*arr)[high], (*arr)[pivotIndex]

	i := low
	j := high - 1
	for {
		for i < high && (*arr)[i] <= pivot {
			i++
		}
		for j >= low && (*arr)[j] > pivot {
			j--
		}
		if i >= j {
			break
		}
		(*arr)[i], (*arr)[j] = (*arr)[j], (*arr)[i]
	}
	(*arr)[i], (*arr)[high] = (*arr)[high], (*arr)[i]

	if num_thread < MAX_THREAD {
		wait.Add(1)
		go quickSortPar(arr, low, i-1, wait, num_thread+1)
	} else {
		quickSort(arr, low, i-1)
	}
	quickSort(arr, i+1, high)
}

func sorted(arr *[]int32) bool {
	for i := 0; i < len(*arr)-1; i++ {
		if (*arr)[i] > (*arr)[i+1] {
			return false
		}
	}
	return true
}

func start_quickSortPar(arr *[]int32) {
	low, high := 0, len(*arr)-1

	start := time.Now()

	var wait sync.WaitGroup
	wait.Add(1)
	quickSortPar(arr, low, high, &wait, 1)
	wait.Wait()
	duration := time.Since(start)

	fmt.Printf("multi-thread quick sort\n")
	fmt.Printf("  sorted: %t\n", sorted(arr))
	fmt.Printf("  duration: %s\n", duration)
}

func start_quickSort(arr *[]int32) {
	low, high := 0, len(*arr)-1

	start := time.Now()
	quickSort(arr, low, high)
	duration := time.Since(start)

	fmt.Printf("single-thread quick sort\n")
	fmt.Printf("  sorted: %t\n", sorted(arr))
	fmt.Printf("  duration: %s\n", duration)
}

func make_data(size int) *[]int32 {
	data := make([]int32, size) //10M int
	for i := range data {
		data[i] = rand.Int31()
	}

	return &data
}

func main() {
	size := flag.Int("d", 10000000, "number of int32 to sort")
	flag.Parse()

	arr := make_data(*size)
	fmt.Println()
	start_quickSort(arr)
	fmt.Println()
	start_quickSort(arr)
	fmt.Println()
}
