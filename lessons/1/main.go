package main

import "fmt"

func main() {

	// func(s string, i int) {
	// 	fmt.Println("anonimous func")
	// 	fmt.Printf("arguments: %s, %d\n", s, i)
	// }("First", 1)

	// inc := increment1(10)
	// fmt.Println(inc())
	// fmt.Println(inc())

	fWords := "Hi"

	var sWords *string = &fWords

	*sWords += " and Hello"

	fmt.Println(fWords == *sWords)
	fmt.Println(&fWords)
	fmt.Println(sWords)

	*sWords += " and Hello"
	fmt.Println(&fWords)
	fmt.Println(sWords)
}

// func increment1(i int) func() int {
// 	count := i
// 	return func() int {
// 		count++
// 		return count
// 	}
// }

// func increment2(i int) int {
// 	count := i
// 	count++
// 	return count
// }

// func findMin(nums ...int) int {

// 	if len(nums) == 0 {
// 		return 0
// 	}

// 	min := nums[0]

// 	for _, val := range nums {
// 		if val < min {
// 			min = val
// 		}
// 	}
// 	return min
// }
