package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")
}

func recursion() {
	var testcount func(int) int
	testcount = func(x int) int {
		if x == 11 {
			return 0
		}
		fmt.Println(x)
		return testcount(x + 1)
	}
	testcount(1)
}

func Named_Return_Values() {
	// In Go, you can name the return values of a function.

	var myFunc = func(x int, y int) (result int) {
		result = x + y
		return
	}

	fmt.Println(myFunc(5, 5))
}

func retrun_value() {
	/*
		func FunctionName(param1 type, param2 type) type {
		  return output
		}
	*/

	add := func(x int, y int) int {
		return x + y
	}
	fmt.Println(add(5, 5))
}

func createFn() {
	/*
		func FunctionName(param1 type, param2 type, param3 type) {
		  // code to be executed
		}
	*/
	var familyName = func(fname string) {
		fmt.Println("Hello", fname, "Refsnes")
	}

	familyName("John")
}

func range_loop() {
	fruits := [3]string{"apple", "orange", "banana"}
	for idx, val := range fruits {
		fmt.Printf("%v\t%v\n", idx, val)
	}
}
func nests() {
	adj := [2]string{"big", "tasty"}
	fruits := [3]string{"apple", "orange", "banana"}
	for i := 0; i < len(adj); i++ {
		for j := 0; j < len(fruits); j++ {
			fmt.Println(adj[i], fruits[j])
		}
	}
}

func loops() {
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
}
