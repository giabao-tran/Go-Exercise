//Viết hàm tạo ra 1 struct về 1 người (gồm tên, nghề nghiệp, năm sinh),
//và struct có method tính tuổi,
//method kiểm tra người ấy có hợp với nghề của mình không ( năm sinh chia hết cho số chữ trong tên)

package main

import (
	"fmt"
	"strings"
	"time"
	"unicode/utf8"
)

type Person struct {
	name, job string
	birthYear int
}

func (p Person) CalculateAge() int {
	currentYear := time.Now().Year()
	return currentYear - p.birthYear
}

func (p Person) IsCompatibleJob() bool {
	nameWithoutSpace := strings.ReplaceAll(p.name, " ", "")
	nameLength := utf8.RuneCountInString(nameWithoutSpace)
	if p.birthYear%nameLength == 0 {
		return true
	} else {
		return false
	}
}

func main() {
	var p Person

	fmt.Println("Enter your name:")
	fmt.Scanln(&p.name)

	fmt.Println("Enter your job:")
	fmt.Scanln(&p.job)
	fmt.Println(p.job)

	fmt.Println("Enter your birth year:")
	fmt.Scanln(&p.birthYear)

	fmt.Println("Person age is:", p.CalculateAge())

	fmt.Println("Is this person compatible job:", p.IsCompatibleJob())
}
