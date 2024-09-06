//Viết ham nhập 2 cạnh của hình chữ nhật, in ra diện tích, chu vi

package main

import "fmt"

func main() {
	var width, height float64

	for {
		fmt.Print("Please type width: ")
		fmt.Scanln(&width)
		if width >= 0 {
			break
		}
		fmt.Println("Width must be non-negative. Please type again.")
	}

	for {
		fmt.Print("Please type height: ")
		fmt.Scanln(&height)
		if height >= 0 {
			break
		}
		fmt.Println("Height must be non-negative. Please type again.")
	}

	area := width * height
	perimeter := 2 * (width + height)

	fmt.Println("Area of the rectangle:", area)
	fmt.Println("Perimeter of the rectangle:", perimeter)
}
