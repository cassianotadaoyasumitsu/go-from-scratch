package main

func main() {

}

// Syntax
func name_of_function(list_of_parameters) (list_of_results) {
	// function body
}

// Examples
func add(x, y int) int {
	return x + y
}
result := add(2, 3)

// Example 2
func rectanglePerimeter(x, y int) int {
	result := add(x, y)
	return result
}

// Example 3
func generateReport(filepath string) {
	// report generation code
}

// Example 4
func generateReport(filepath string) (size int, err error) {
	// report generation code
}

// Example 5 "RECURSION"
func fibo(x int) int {
	if x <= 1 {
		return x
	}
	return fibo(n-1) + fibo(n-2)
}

// Example leetcode.com
func twoSum(nums []int, target int) []int {
	x := make(map[int][]int)
	for i, v := range nums {
		x[v] = append(x[v], i)
	}

	var result []int
	for k, v := range x {
		second := target - k
		if len(x[k]) > 1 && (second == k) {
			result = append(result, x[k][0], x[k][1])
			break
		} else if _, ok := x[second]; ok && x[second][0] != x[k][0] {
			result = append(result, v[0], x[second][0])
			break
		}
	}

	sort.Ints(result)
	return result
}

// Anonymous Functions

// Example
func main() {
	proc := func() {
		fmt.Println("Inside anonymous function")
	}
	proc()
}

// Example 2
go func(x int) {
	fmt.Println(x * x)
} (10)

// Example 3
fruits := []string{"banana", "apple", "pineapple"}
comparator := func(i, j int) bool {
	return len(fruits[i]) < len(fruits[j])
}
sort.Slice(fruits, comparator)