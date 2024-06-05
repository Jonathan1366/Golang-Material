// package main

// import "fmt"

// func twoSum(nums[]int, target int)[]int {

// 	seen:= make(map[int]int)

// 	for i, num:= range nums{
// 		needs:= target-num
// 		if j, exists:=seen[needs];exists{
// 			return []int{j, i}
// 		}
// 		seen[num]=i
// 	}
// 	return [] int{}

// }

// func main(){
// 	nums := []int{2, 7, 11, 15}
// 	target := 9
// 	result := twoSum(nums, target)

// 	// Cetak hasilnya
// 	fmt.Println(result) // Output: [0, 1]
// }