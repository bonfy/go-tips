package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name string
	Age  int
}

func main() {
	people := []Person{
		Person{Name: "bonfy", Age: 18},
		Person{Name: "rene", Age: 16},
	}
	fmt.Println("Origin: ", people)

	// Marshal: struct -> []byte
	js, err := json.Marshal(people)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Byte: ", js)
	fmt.Println("String: ", string(js))

	// Unmarshal: []byte -> struct
	var newPeople []Person
	err = json.Unmarshal(js, &newPeople)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Decode: ", newPeople)
}

// Refer: http://cizixs.com/2016/12/19/golang-json-guide

// Result:

// Origin:  [{bonfy 18} {rene 16}]
// Byte:  [91 123 34 78 97 109 101 34 58 34 98 111 110 102 121 34 44 34 65 103 101 34 58 49 56 125 44 123 34 78 97 109 101 34 58 34 114 101 110 101 34 44 34 65 103 101 34 58 49 54 125 93]
// String:  [{"Name":"bonfy","Age":18},{"Name":"rene","Age":16}]
// Decode:  [{bonfy 18} {rene 16}]
