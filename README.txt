Go command 

ใน Go 1 folder มีแค่ชื่อ package ได้แค่ชื่อเดียว


Go ต้องมี 2 อย่างล่างในโค้ดเสมอ (ตัว package กับ func จะตั้งเป็นชื่ออะไรก็ได้ เหมือน C language)
: package main
: func main(){}


Run code , Build execute file / การรันโค้ดและการสร้างไฟล์ execute
: go run main.go 
: go build main.go 


import library / วิธีลง library
: go get "github.com/example_user/example_library" 


check imported library / วิธีเช็ค library ที่มีการลงไว้
: go list -m all


init project / เริ่มต้นการใช้ Go module ในโปรเจ็กและจัดการ dependency ในโปรเจ็ก
: go mod init github.com/example_user/example_repository // แบบนี้ไว้เผื่อว่าจะนำไปขึ้นบน github
: go mod init (start) 


nodemon for go
: npx nodemon --exec go run main.go --signal SIGTERM

public , private function / การทำ private function ตัวหน้าต้องเป็นตัวเล็ก
: func Say() {} // public 
: func say() {} // private  

private function ใน package เดียวกันสามารถเรียกใช้ได้ใน root level เดียวกันได้ 
เช่น เรียกใช้ privateTest() ของ test.go ใน lovemxchine.go ได้


struct is define new type // struct คือการสร้าง type ใหม่ 
: type NewType struct {
    Name string 
	Weight int 
	Height int 
	Grade string
}
: var newName NewType
: newName.Name = "noey"
: fmt.Println(newName.Name) // output : noey


แนวคคิดการทำ method ไปดูในไฟล์ lovemxchine/method.go 
null ในภาษา go คือ nil

เพิ่มเติมเกี่ยวกับ pointer 
การทำ parameter function กับตัวแปรทั่วไป ตอนจะเรียกใช้ใน function ต้องมี * นำหน้า แต่ถ้าเป็น struct จะสามารถใช้ได้เลย
example (normal): 
	func abc(a *int){
		*a += 1
	}
example (struct): 
	func abc(a *struct){
		a.amount += 1
	}	
