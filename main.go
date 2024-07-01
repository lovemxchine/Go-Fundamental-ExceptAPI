package main

import (
	"fmt"

	"github.com/google/uuid"
	first_package "github.com/lovemxchine/go-example/lovemxchine" // define lovemxchine package as first_package ?
	// lovemxchine "github.com/lovemxchine/go-example/package" // this should define as package name
)

// Create Struct
type Student struct {
	Name   string
	Weight int
	Height int
	Grade  string
}

func main() {
	id := uuid.New()
	name := "channarong"
	fmt.Println(name)
	fmt.Printf("UUID : %s \n", id)
	first_package.Saylovemxchine()
	first_package.ExportTest()

	var arr = [5]int{5, 0, 2, 1, 7}
	var slic = []string{"noey", "test", "again"}
	var myArray [3]string
	myArray[0] = "noey"
	myArray[1] = "test"
	myArray[2] = "array"
	var myArray2 [3]string
	myArray2[0] = "noey"
	myArray2[1] = "test"
	myArray2[2] = "array2"
	myArray = myArray2
	slic = myArray[:]
	// var slic []int
	for i := 0; i < len(myArray); i++ {
		fmt.Println(myArray[i])
	}
	fmt.Println(slic)
	fmt.Println(arr)
	//map
	myMap := make(map[string]int)
	myMap["go"] = 3
	myMap["lang"] = 2
	myMap["geez"] = 31
	fmt.Println(myMap["go"])
	for key, value := range myMap {
		fmt.Printf("%s => %d\n", key, value)
	}
	for _, value := range myMap { // การทำ for _ , value แบบนี้คือเอาแต่ค่าที่เป็นแค่ value ไม่เอา key
		fmt.Printf("%d \n", value)
	}
	val, ok := myMap["geez"]
	if ok {
		fmt.Println(val, ok)
	}

	// use Struct
	var studentNoey Student
	studentNoey.Name = "noey"
	studentNoey.Weight = 55
	studentNoey.Height = 172
	studentNoey.Grade = "A"
	fmt.Println(studentNoey)
	fmt.Println(studentNoey.Name)
	first_package.SayTeacher()
	first_package.CheckOutput()
	first_package.SetWeapon()
	// var mySlice = []int{1, 2, 3, 4, 5}
	// mySlice = append(mySlice[:2], mySlice[2+1:]...)
	// fmt.Println(mySlice)
}
