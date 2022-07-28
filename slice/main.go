package main

import (
	"fmt"
)

func main() {
	var nums [3]int = [3]int{1, 2, 3}
	fmt.Println(nums[0])
	fmt.Println(len(nums))

	numbers := [3]int{1, 2, 3}
	fmt.Println(numbers)

	// 1, 2, 3の要素を持つスライスを作成して代入
	nums2 := []int{1, 2, 3}

	// あるいは既存の配列やスライスからも範囲アクセスでスライス作成
	nums3 := nums[0:2]
	nums4 := nums2[1:3]

	// 範囲外アクセスはパニック
	fmt.Println(nums2[1])

	// 要素の割り当ても可能
	nums2[0] = 100

	fmt.Println(nums2)

	nums2 = append(nums2, 4)

	fmt.Println(nums2)
	fmt.Println(nums3)
	fmt.Println(nums4)
}
