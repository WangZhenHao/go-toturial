package main

import "fmt"

func main() {
	/*
		var a = map[KeyType]ValueType{key1:value1, key2:value2,...}
		b := map[KeyType]ValueType{key1:value1, key2:value2,...}
	*/
	var a = map[string]string{"brand": "Ford", "model": "Mustang", "year": "1964"}
	b := map[string]int{"Oslo": 1, "Bergen": 2, "Trondheim": 3, "Stavanger": 4}

	fmt.Printf("a\t%v\n", a)
	fmt.Printf("b\t%v\n", b)

	/*
		a   map[brand:Ford model:Mustang year:1964]
		b   map[Bergen:2 Oslo:1 Stavanger:4 Trondheim:3]
	*/
}

func remove_element() {
	/*
		delete(map_name, key)
	*/
	var a = make(map[string]string)
	a["brand"] = "Ford"
	a["model"] = "Mustang"
	a["year"] = "1964"

	fmt.Println(a)

	delete(a, "year")

	fmt.Println(a)

	/*
		map[brand:Ford model:Mustang year:1964]
		map[brand:Ford model:Mustang]
	*/
}

func emty_map() {
	/*
		var a map[KeyType]ValueType
		The default value of a map is nil.
	*/
	var a = make(map[string]string)
	var b map[string]string

	fmt.Println(a == nil)
	fmt.Println(b == nil)
	/*
		false
		true
	*/
}

func make_map() {
	/*
		var a = make(map[KeyType]ValueType)
		b := make(map[KeyType]ValueType)
	*/
	var a = make(map[string]string) // The map is empty now
	a["brand"] = "Ford"
	a["model"] = "Mustang"
	a["year"] = "1964"
	// a is no longer empty
	b := make(map[string]int)
	b["Oslo"] = 1
	b["Bergen"] = 2
	b["Trondheim"] = 3
	b["Stavanger"] = 4

	fmt.Printf("a\t%v\n", a)
	fmt.Printf("b\t%v\n", b)

	/*
		a   map[brand:Ford model:Mustang year:1964]
		b   map[Bergen:2 Oslo:1 Stavanger:4 Trondheim:3]
	*/
}
