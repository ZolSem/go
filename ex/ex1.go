package main

import (
	"errors"
	"fmt"
)

func main() {

	PrintSecondMax([]int{5, 5, 1, 6, 6})

}

func secondMax(arr []int) (int, error) {
	if len(arr) < 2 {
		return -1, errors.New("arr length is shorter than 2")
	}

	max1, max2 := arr[0], arr[1]

	if max1 < max2 {
		max1, max2 = max2, max1
	}

	for _, v := range arr {

		if max1 < v {
			max2 = max1
			max1 = v
		} else {
			if max2 < v && v != max1 {
				max2 = v
			}
		}
	}

	if max1 == max2 {
		return -1, errors.New("the same arguments in arr")
	}

	return max2, nil
}

func PrintSecondMax(arr []int) {

	i, err := secondMax(arr)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(i)
	}

}
