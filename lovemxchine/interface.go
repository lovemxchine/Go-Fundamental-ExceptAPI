package lovemxchine

import "fmt"

// Speaker interface
// เพิ่มเติมแนวคิด interface คือการกำหนดคุณสมบัติกลางได้ตราบเท่าที่ทำตามที่ Speaker (interface) กำหนด
// ไว้ทำ config
type Speaker interface {
	Speak() string
}

// Dog struct
type Dog struct {
	Name string
}

// Dog กลายเป็นกลุ่มของ Speaker เพราะสร้าง method Speak() ที่อยู่ภายใน Speaker interface ทันทีดูตรงบรรทัดที่ 6 
func (d Dog) Speak() string {
	return "Woof!"
}

// Person struct
type Person struct {
	Name string
}

// Person กลายเป็นกลุ่มของ Speaker ทันทีดูตรงบรรทัดที่ 6 
func (p Person) Speak() string {
	return "Hello!"
}

// function that accepts Speaker interface
func makeSound(s Speaker) {
	fmt.Println(s.Speak())
}


func Run_this() {
	dog := Dog{Name: "Buddy"}
	person := Person{Name: "Alice"}

	makeSound(dog)
	makeSound(person)
}