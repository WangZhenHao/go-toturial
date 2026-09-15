package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")
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
