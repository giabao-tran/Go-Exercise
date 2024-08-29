//Viết chương trình nhập giải bài toán twosum : https://leetcode.com/problems/two-sum/

package main

import "fmt"

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		diff := target - nums[i]

		if value, exist := m[diff]; exist {
			return []int{value, i}
		}

		m[nums[i]] = i

	}
	return nil
}

func main() {
	var slice []int
	var length int
	var target int
	fmt.Println("Enter the length of the array:")
	fmt.Scan(&length)
	for i := 0; i < length; i++ {
		var num int
		fmt.Print("Enter the value of element:")
		fmt.Scan(&num)
		slice = append(slice, num)
	}
	fmt.Println("Enter the target:")
	fmt.Scan(&target)
	fmt.Println(twoSum(slice, target))

}
