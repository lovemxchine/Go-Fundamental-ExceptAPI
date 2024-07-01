package lovemxchine

import "fmt"

type Teacher struct {
	FirstName string
	LastName  string
	Age       int
	Subject   string
}
type Teacher1 struct {
	FirstName string
	LastName  string
	Age       int
	Subject   string
}


// แนวคิดการทำ method ในภาษา go การเขียนแบบข้างล่างนี้คือกำหนดว่าเป็น method ของ teacher
func (s Teacher) FullName() string {
	return s.FirstName + " " + s.LastName
}
func (s Teacher1) FullName() string {
	return s.FirstName + " " + s.LastName
}

func SayTeacher() {
	var mrWilson = Teacher{FirstName:"Wilson",LastName:"Walt",Age:20,Subject:"Math"}
	var mrWilson1 = Teacher1{FirstName:"Wilson",LastName:"Walt",Age:20,Subject:"Math"}
	fmt.Println(mrWilson) 
	fmt.Println(mrWilson.FullName()) 
	fmt.Println(mrWilson1.FullName()) 
}