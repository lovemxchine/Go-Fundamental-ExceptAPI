package lovemxchine

import "fmt"

func privateTest() {
	fmt.Println("this is private function")
}

func ExportTest() {
	privateTest()
}
