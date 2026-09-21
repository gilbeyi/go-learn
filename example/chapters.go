package main

import (
	"fmt"
	"maps"
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

func arraySlice() {
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

func arrayMap() {
	var nilMap map[string]int
	fmt.Println("nilMap", nilMap)
	fmt.Println("nilMap isNil", nilMap == nil)

	initMap := map[string]int{}
	fmt.Println("initMap", initMap)
	fmt.Println("initMap isNil", initMap == nil)

	users := map[string][]string{
		"admin": {"admin user", "power user"},
		"write": {"write user"},
		"read":  {"read user"},
	}
	fmt.Println("users", users)

	wins := map[string]int{}
	wins["dogers"] = 25
	wins["padres"] = 20
	wins["angels"] = 10
	fmt.Println("wins", wins)

	valueDogers, matchDogers := wins["dogers"]
	fmt.Println("valueDogers, matchDogers", valueDogers, matchDogers)
	valueYankees, matchYankees := wins["yankees"]
	fmt.Println("valueYankees, matchYankees", valueYankees, matchYankees)

	delete(wins, "angels")
	fmt.Println("delete wins", wins)

	clear(wins)
	fmt.Println("clear wins", wins)
	fmt.Println("wins isNil", wins == nil)

	mapA := map[string]int{
		"a": 1, "b": 2,
	}
	mapB := map[string]int{
		"a": 1, "b": 2,
	}
	mapC := map[string]int{
		"a": 2, "b": 3,
	}
	fmt.Println("mapA == mapB", maps.Equal(mapA, mapB))
	fmt.Println("mapA == mapC", maps.Equal(mapA, mapC))

}

func learnStruct() {
	type typePerson struct {
		name string
		age  int
	}

	m := typePerson{
		"man",
		30,
	}
	fmt.Println("m: person", m)
	fmt.Println("m.name: person", m.name)

	var person struct {
		name string
		age  int
	}
	person.name = "man"
	person.age = 20
	fmt.Println("person", person)
}

func main() {
	// literal()
	// arraySlice()
	// arrayMap()
	learnStruct()
}
