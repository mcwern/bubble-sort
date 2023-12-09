package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	s := makeRandomSlice(10, 100)
	bubbleSort(s)
	printSlice(s, 11)
	isSorted := checkSorted(s)

	if isSorted {
		fmt.Println("The slice is sorted!")
	} else {
		fmt.Println("The slice is NOT sorted!")
	}
}

func makeRandomSlice(numItems, max int) []int {
	r := rand.New(rand.NewSource(time.Now().UnixNano())) // Initialize with a changing seed

	slice := make([]int, numItems)

	for i := 0; i < numItems; i++ {
		slice[i] = r.Intn(max)
	}

	return slice
}

func printSlice(s []int, numItems int) {
	if len(s) < numItems {
		fmt.Println(s)
	} else {
		fmt.Println(s[:numItems])
	}
}

func checkSorted(s []int) bool {
	isSorted := false
	for i := 0; i < len(s); i++ {
		if i == len(s)-1 {
			isSorted = true
		} else if s[i] < s[i+1] || s[i] == s[i+1] {
			isSorted = true
		} else {
			isSorted = false
			return isSorted
		}
	}

	return isSorted
}

func bubbleSort(slice []int) {
	for !checkSorted(slice) {
		for i := 1; i < len(slice); i++ {
			if slice[i-1] > slice[i] {
				slice[i-1], slice[i] = slice[i], slice[i-1]
			}
		}
	}
}
