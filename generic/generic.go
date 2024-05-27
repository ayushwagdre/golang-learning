package generic_learning

import (
	"fmt"
)

type number interface {
	int64 | float64
}

func sum[k comparable, v int64 | float64](data map[k]v) v {
	var sum v
	for _, value := range data {
		sum += value
	}
	return sum
}

func sumUsingNumber[k comparable, v number](data map[k]v) v {
	var sum v
	for _, value := range data {
		sum += value
	}
	return sum
}

func exampleFunc[T any](param T) {
	fmt.Printf("Type: %T, Value: %v\n", param, param)
}

func print[T any](param []T) {
	for _, v := range param {
		fmt.Printf("Type: %T, Value: %v\n", v, v)
	}
}

func compareArray[T number](param []T) T {
	var max T
	max = 0
	for _, value := range param {
		if value > max {
			max = value
		}
	}
	return max
}

func GenericLearning() {
	data := map[string]int64{"A": 1, "B": 2, "C": 3}
	fmt.Println(sum(data)) // Prints the sum of int64 values
	fmt.Println("sumUsingNumber", sumUsingNumber(data))

	data2 := map[string]float64{"A": 1.1, "B": 2.2, "C": 3.3}
	fmt.Println(sum(data2)) // Prints the sum of float64 values
	fmt.Println("sumUsingNumber", sumUsingNumber(data2))

	sumUsingNumber(data)

	exampleFunc(5)       // T is int
	exampleFunc("hello") // T is string
	exampleFunc(true)    // T is bool

	intSlice := []int{10, 5, 7, 14, 3}
	floatSlice := []float64{3.14, 2.718, 1.618, 0.577}
	stringSlice := []string{"apple", "banana", "orange", "kiwi"}
	print(intSlice)
	print(floatSlice)
	print(stringSlice)

	println(compareArray(floatSlice))
}
