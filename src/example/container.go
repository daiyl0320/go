package example

import "fmt"
import "framework"

type c struct {
	obj framework.Object
}

var c1 = c{framework.Object{"Container", "This is a pratice about containers"}}

func (e *c) Id() string {
	return c1.obj.Id
}

func (e *c) Description() string {
	return c1.obj.Description
}

func init() {
	framework.Register(&c1)
}

func (e *c) Run() {
	testArray()
	testSlice()
	testMaps()
	testRange()
}

func testArray() {
	fmt.Println("test array=========================")
	var nums = []int{1, 2, 3, 4, 5}
	fmt.Println("Before change, nums is ", nums)

	nums[0] = 5
	nums[4] = 0
	fmt.Println("After change, nums is ", nums)
}

func testSlice() {
	fmt.Println("test slice=========================")
	var slice1 = make([]string, 3)
	fmt.Println("slice1 is ", slice1)

	slice1[0] = "ab"
	slice1[1] = "cd"
	slice1[2] = "ef"
	fmt.Printf("slice1's size is %v, value is %v\n", len(slice1), slice1)

	slice1 = append(slice1, "gh")
	fmt.Printf("slice1's size is %v, value is %v", len(slice1), slice1)

	slice2 := make([]string, len(slice1))
	copy(slice2, slice1)
	fmt.Printf("slice2's size is %v, value is %v\n", len(slice2), slice2)

	fmt.Printf("slice2's [2 - 4] is %v\n", slice2[1:])
	fmt.Printf("slice2's [2 - 3] is %v\n", slice2[1:3])

	var slice3 = []string{"aa", "bb", "cc", "dd", "ee"}
	fmt.Printf("slice3's size is %v, value is %v\n", len(slice3), slice3)

	var slice4 = [][]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Printf("slice4's size is %v, value is %v\n", len(slice4), slice4)
	fmt.Printf("slice4[0]'s size is %v, value is %v\n", len(slice4[0]), slice4[0])
}

func testMaps() {
	fmt.Println("test maps=========================")
	map1 := make(map[string]int)

	map1["a"] = 1
	map1["b"] = 2
	map1["c"] = 3
	fmt.Printf("map1's size is %v, map1's value is %v\n", len(map1), map1)

	map1["b"] = 4
	fmt.Printf("map1's size is %v, map1's value is %v\n", len(map1), map1)

	delete(map1, "b")
	fmt.Printf("map1's size is %v, map1's value is %v\n", len(map1), map1)

	value, ok := map1["b"]
	fmt.Printf("value is %v, ok is %v\n", value, ok)

	map2 := map[string]int{"aa": 1, "bb": 2}
	fmt.Printf("map2's size is %v, map2's value is %v\n", len(map2), map2)
}

func testRange() {
	fmt.Println("test range=========================")
	nums := []int{1, 2, 3, 4, 5}

	for index, num := range nums {
		fmt.Printf("index is %v, value is %v\n", index, num)
	}

	dict := map[string]int{"xx": 1, "yy": 2, "zz": 3}
	for key, value := range dict {
		fmt.Printf("key is %v, value is %v\n", key, value)
	}

	for keys := range dict {
		fmt.Printf("key is %v\n", keys)
	}
}
