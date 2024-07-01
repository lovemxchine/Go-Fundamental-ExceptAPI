package lovemxchine

import "fmt"

func getRef(addr *int) {
	*addr = 500
}

func getRef2(val int) {
	val = 500
}

func CheckOutput() {
	var x int = 20
	var b int = 20
	getRef(&x)
	getRef2(b)
	fmt.Println(x)
}