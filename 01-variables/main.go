package main

import "fmt"

func main() {

}

func slices_copy() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	// Original slice
	fmt.Printf("numbers = %v\n", numbers)
	fmt.Printf("length = %d\n", len(numbers))
	fmt.Printf("capacity = %d\n", cap(numbers))

	// Create copy with only needed numbers
	neededNumbers := numbers[:len(numbers)-10]
	numbersCopy := make([]int, len(neededNumbers))
	copy(numbersCopy, neededNumbers)

	fmt.Printf("numbersCopy = %v\n", numbersCopy)
	fmt.Printf("length = %d\n", len(numbersCopy))
	fmt.Printf("capacity = %d\n", cap(numbersCopy))

	/*
		numbers = [1 2 3 4 5 6 7 8 9 10 11 12 13 14 15]
		length = 15
		capacity = 15
		// New slice
		numbersCopy = [1 2 3 4 5]
		length = 5
		capacity = 5
	*/
}

func slices_append() {
	myslice1 := []int{1, 2, 3, 4, 5, 6}
	fmt.Printf("myslice1 = %v\n", myslice1)
	fmt.Printf("length = %d\n", len(myslice1))
	fmt.Printf("capacity = %d\n", cap(myslice1))

	myslice1 = append(myslice1, 20, 21)
	fmt.Printf("myslice1 = %v\n", myslice1)
	fmt.Printf("length = %d\n", len(myslice1))
	fmt.Printf("capacity = %d\n", cap(myslice1))
	/*
		myslice1 = [1 2 3 4 5 6]
		length = 6
		capacity = 6
		myslice1 = [1 2 3 4 5 6 20 21]
		length = 8
		capacity = 12

	*/
}
func slices_make() {
	/*
		myslice1 = [0 0 0 0 0]
		length = 5
		capacity = 10
		myslice2 = [0 0 0 0 0]
		length = 5
		capacity = 5
	*/
	myslice1 := make([]int, 5, 10)
	fmt.Printf("myslice1 = %v\n", myslice1)
	fmt.Printf("length = %d\n", len(myslice1))
	fmt.Printf("capacity = %d\n", cap(myslice1))

	// with omitted capacity
	myslice2 := make([]int, 5)
	fmt.Printf("myslice2 = %v\n", myslice2)
	fmt.Printf("length = %d\n", len(myslice2))
	fmt.Printf("capacity = %d\n", cap(myslice2))
}

func slices() {
	/*
		Slices are similar to arrays, but are more powerful and flexible.

		Like arrays, slices are also used to store multiple values of the same type in a single variable.

		However, unlike arrays, the length of a slice can grow and shrink as you see fit.

		In Go, there are several ways to create a slice:

		Using the []datatype{values} format
		Create a slice from an array
		Using the make() function
	*/
	myslice1 := []int{}
	fmt.Println(len(myslice1))
	fmt.Println(cap(myslice1))
	fmt.Println(myslice1)

	myslice2 := []string{"Go", "Slices", "Are", "Powerful"}
	fmt.Println(len(myslice2))
	fmt.Println(cap(myslice2))
	fmt.Println(myslice2)

	arr1 := [6]int{10, 11, 12, 13, 14, 15}
	myslice := arr1[2:4]

	fmt.Printf("myslice = %v\n", myslice)
	fmt.Printf("length = %d\n", len(myslice))
	fmt.Printf("capacity = %d\n", cap(myslice))

	/*
		myslice = [12 13]
		length = 2
		capacity = 4 // 从切片起始位置到原数组末尾的元素个数 6 - 2 = 4
	*/

}

func constants() {
	// Go Constants
	const (
		A int = 1
		B     = 3.14
		C     = "Hi!"
	)
}

func variable() {
	var a string
	var b int
	var c bool
	var d int = 1
	var f float32 = 2.4
	e := 2

	e = 3

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(f)
	fmt.Println(e)
}

func array() {
	// array
	var arr1 = [3]int{1, 2, 3}
	arr2 := [5]int{4, 5, 6, 7, 8}
	fmt.Println(arr1)
	fmt.Println(arr2)

	//This example declares two arrays (arr1 and arr2) with inferred lengths:

	var arr3 = [...]int{1, 2, 3}
	arr4 := [...]int{4, 5, 6, 7, 8}

	fmt.Println(arr1)
	fmt.Println(arr2)
	fmt.Println(arr3)
	fmt.Println(arr4)

	prices := [3]int{10, 20, 30}

	prices[2] = 50
	fmt.Println(prices)

	fmt.Println(prices[0])
	fmt.Println(prices[2])

	// Initialize Only Specific Elements,This example initializes only the second and third elements of the array:

	arr5 := [5]int{1: 10, 2: 40}

	fmt.Println(arr5)

}
