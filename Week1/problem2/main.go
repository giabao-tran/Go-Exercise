//Viết chương trình nhập 1 string, in ra true nếu độ dài chuỗi chia hết cho 2, và false nếu ngược lại

package main

import "fmt"

func main() {
	var s string
	fmt.Print("Please type the string: ")
	fmt.Scan(&s)
	if len(s)%2 == 0 {
		fmt.Print("true")
	} else {
		fmt.Print("false")
	}
}
