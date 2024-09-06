//Viết chương trình nhập một slice số, in ra tổng, số lớn nhất, số nhỏ nhất, trung bình cộng, slice đã được sắp xếp

package main

import (
	"fmt"
	"slices"
)

func main() {
	var slice []int
	var length int
	sum := 0
	fmt.Print("Please type length of slice: ")
	fmt.Scanln(&length)

	for i := 0; i < length; i++ {
		var num int
		fmt.Print("Please type number: ")
		fmt.Scanln(&num)
		slice = append(slice, num)
		sum += num
	}

	slices.Sort(slice)

	fmt.Println("The max number of slice is:", slice[len(slice)-1])
	fmt.Println("The min number of slice is:", slice[0])

	fmt.Println("Sorted slice is:", slice)

	slices.Reverse(slice)
	fmt.Println("Reversed slice is:", slice)

	fmt.Println("Sum of slice is:", sum)

	//I want learn everything about slices in go should I know

}
