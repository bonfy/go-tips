package main

import (
	"encoding/json"
	"fmt"
)

type Cat struct {
	ID   uint
	Age  int
	Name string
}

var cats = []Cat{
	Cat{ID: 1, Age: 2, Name: "cat-1"},
	Cat{ID: 2, Age: 3, Name: "cat-2"},
}

type Dog struct {
	ID   uint
	Name string
}

var dogs = []Dog{
	Dog{ID: 1, Name: "dog-1"},
	Dog{ID: 2, Name: "dog-2"},
}

func FindCatById(id uint) Cat {
	for _, item := range cats {
		if item.ID == id {
			return item
		}
	}
	return Cat{}
}

func FindDogById(id uint) Dog {
	for _, item := range dogs {
		if item.ID == id {
			return item
		}
	}
	return Dog{}
}

// 类似 Python 的 DuckType

// GetModelById :封装通用
func GetModelById(id uint, tp string) interface{} {
	if tp == "dog" {
		return FindDogById(id)
	}

	if tp == "cat" {
		return FindCatById(id)
	}

	return nil
}

// GetJSON 通用函数的适用
func GetJSON(id uint, tp string) {
	item := GetModelById(id, tp)

	js, err := json.Marshal(item)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(js))
}

func main() {
	fmt.Println(FindCatById(1))
	fmt.Println(FindDogById(2))

	fmt.Println(GetModelById(1, "dog"))
	fmt.Println(GetModelById(2, "cat"))
	fmt.Println(GetModelById(3, "dog"))

	GetJSON(1, "dog")
	GetJSON(3, "cat")

	GetJSON(4, "tiger")
}
