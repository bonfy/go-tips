package main

import (
	"encoding/json"
	"fmt"

	"github.com/fatih/structs"
	"github.com/mitchellh/mapstructure"
)

// Person Struct
type Person struct {
	Name string
	Age  int
}

func main() {
	fmt.Println("=====Single====")
	// Single Object
	structPerson := Person{Name: "bonfy", Age: 18}
	fmt.Println("Origin:", structPerson)

	var mapPerson map[string]interface{}

	// Struct to map
	mapPerson2 := structs.Map(structPerson)
	fmt.Println("Map:", mapPerson2)

	// Struct to byte
	bytePerson, err := json.Marshal(structPerson)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Byte: ", bytePerson)
	fmt.Println("String: ", string(bytePerson))

	// Byte to map
	err = json.Unmarshal(bytePerson, &mapPerson)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Map:", mapPerson)

	// Byte to struct
	var bSturctPerson Person
	err = json.Unmarshal(bytePerson, &bSturctPerson)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Byte to Struct:", bSturctPerson)
	fmt.Println("structPerson == bSturctPerson ?", structPerson == bSturctPerson)

	// Map to Stuct
	var mStructPerson Person
	mapstructure.Decode(mapPerson, &mStructPerson)

	fmt.Println("Map to Struct:", mStructPerson)
	fmt.Println("structPerson == mStructPerson ?", structPerson == mStructPerson)

	fmt.Println("=====List====")

	// List Object

	structPeople := []Person{
		Person{Name: "bonfy", Age: 18},
		Person{Name: "john", Age: 16},
	}
	fmt.Println("Origin:", structPeople)

	// List Struct to map
	var mapPeople []map[string]interface{}
	for _, structP := range structPeople {
		mapP := structs.Map(structP)
		mapPeople = append(mapPeople, mapP)
	}
	fmt.Println("List Sturct to Map:", mapPeople)

	// List Struct to byte
	bytePeople, err := json.Marshal(structPeople)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Byte: ", bytePeople)
	fmt.Println("String: ", string(bytePeople))

	// Byte to map
	// var mapPeople map[string]interface{}
	// err = json.Unmarshal(bytePerson, &mapPeople)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// fmt.Println("Map:", mapPeople)

	// Byte to List struct
	var bSturctPeople []Person
	err = json.Unmarshal(bytePeople, &bSturctPeople)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Byte to List Struct:", bSturctPeople)

	// Map to Stuct
	var mStructPeople []Person
	mapstructure.Decode(mapPeople, &mStructPeople)

	fmt.Println("Map to List Struct:", mStructPeople)
}

// Result

// =====Single====
// Origin: {bonfy 18}
// Map: map[Name:bonfy Age:18]
// Byte:  [123 34 78 97 109 101 34 58 34 98 111 110 102 121 34 44 34 65 103 101 34 58 49 56 125]
// String:  {"Name":"bonfy","Age":18}
// Map: map[Name:bonfy Age:18]
// Byte to Struct: {bonfy 18}
// structPerson == bSturctPerson ? true
// Map to Struct: {bonfy 18}
// structPerson == mStructPerson ? true
// =====List====
// Origin: [{bonfy 18} {john 16}]
// List Sturct to Map: [map[Name:bonfy Age:18] map[Name:john Age:16]]
// Byte:  [91 123 34 78 97 109 101 34 58 34 98 111 110 102 121 34 44 34 65 103 101 34 58 49 56 125 44 123 34 78 97 109 101 34 58 34 106 111 104 110 34 44 34 65 103 101 34 58 49 54 125 93]
// String:  [{"Name":"bonfy","Age":18},{"Name":"john","Age":16}]
// Byte to List Struct: [{bonfy 18} {john 16}]
// Map to List Struct: [{bonfy 18} {john 16}]
