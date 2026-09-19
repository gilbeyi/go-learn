package main

import (
	"fmt"
	"slices"
)

func literal() {
	integer2 := 0b01
	integer8 := 0o6
	integer16 := 0x8

	string1 := "あいうえお"
	string2 := `あいうえお\nかきくけこ`
	rune1 := '\141'

	println(integer2)
	println(integer8)
	println(integer16)
	println(string1)
	println(string2)
	println(rune1)
}

func compositeTypeArrayAndSlice() {
	var array = [3]int{10, 20, 30}        // 配列
	var slice = []int{10, 20, 30}         // slice
	var sliceIndex = []int{10, 5: 20, 30} // slice
	var sliceNil []int

	sliceMake := make([]int, 5)       // make
	sliceMakeCap := make([]int, 5, 6) // make
	sliceCopy := make([]int, 2)

	fmt.Println("array", array)
	fmt.Println("len(array)", len(array))

	fmt.Println("slice", slice)
	fmt.Println("sliceIndex", sliceIndex)
	fmt.Println("len(sliceIndex)", len(sliceIndex))
	sliceIndex[2] = 15
	fmt.Println("sliceIndex", sliceIndex)
	sliceIndex = append(sliceIndex, 35)
	fmt.Println("append sliceIndex 35", sliceIndex)
	sliceIndex = append(sliceIndex, 40)
	fmt.Println("append sliceIndex 40", sliceIndex)

	fmt.Println("sliceNil", sliceNil)
	fmt.Println("isNil", sliceNil == nil)
	fmt.Println("slice : sliceIndex", slices.Equal(slice, sliceIndex))

	fmt.Println("make", sliceMake)
	fmt.Println("make capacity", sliceMakeCap)

	// capacityを超えてもappendできる
	sliceMakeCap = append(sliceMakeCap, 40)
	fmt.Println("append make capacity 40", sliceMakeCap)
	sliceMakeCap = append(sliceMakeCap, 50)
	fmt.Println("append make capacity 50", sliceMakeCap)

	// capacity を超える要素を指定すると panic
	// sliceMakeCap[10] = 100
	// fmt.Println("make capacity 100", sliceMakeCap)

	clear(sliceMakeCap)
	fmt.Println("clear", sliceMakeCap)

	fmt.Println("sliceCopy", sliceCopy)
	fmt.Println("slice", slice)
	copy(sliceCopy, slice)
	fmt.Println("copy slice to sliceCopy", sliceCopy)
}

func main() {
	// literal()
	compositeTypeArrayAndSlice()
}
