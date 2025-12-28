package src

import "fmt"

func ArraysAndSlices() {
	// arrays
	var ar [10]int

	// slices
	var sl []int

	for i := range 10 {
		ar[i] = i
		sl = append(sl, i)
	}

	// slice with length and capacity
	s := make([]int, 5, 10)
	fmt.Println(s)

	s = append(s, 10)
	fmt.Println(s)

	// copy and references

	a := make([]int, 10)
	b := a                     // b will be refering to the same data
	c := append([]int{}, a...) // deep copy

	fmt.Println(b)
	fmt.Println(c)

	fmt.Println(sl)
	ReverseElements(sl)
	fmt.Println(sl)

	RotateArray(sl, 3)
	fmt.Println(sl)
}

func ReverseElements(nums []int) {
	i, j := 0, len(nums)-1
	for i < j {
		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}
}

// rotate to right by k
func RotateArray(nums []int, k int) {
	n := len(nums)
	k %= n

	ReverseElements(nums)
	ReverseElements(nums[:k])
	ReverseElements(nums[k:])
}
