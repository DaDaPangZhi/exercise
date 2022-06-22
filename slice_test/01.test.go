package main

import "fmt"

func main() {
	test1()
	test3()
}

func test1() {
	testSlice1 := []int{1, 2, 3}
	testSlice2 := testSlice1
	for i := 0; i < len(testSlice2); i++ {
		testSlice2[i] = 1
	}
	fmt.Println(testSlice1, testSlice2)
	fmt.Println(&testSlice1[0], &testSlice2[0])
}

func test2() {
	testSlice1 := []int{1, 2, 3}
	fmt.Println(&testSlice1[0], &testSlice1[1], &testSlice1[2])
	for _, v := range testSlice1 {
		v = 1
		fmt.Println(&v)
	}
	fmt.Println(testSlice1)
}

func test3() {
	testSlice1 := []int{1, 2, 3}
	testSlice2 := testSlice1
	fmt.Println(&testSlice1[0], &testSlice2[0])
	testSlice2 = append(testSlice2, 4)
	fmt.Println(testSlice1, testSlice2)
	fmt.Println(&testSlice1[0], &testSlice2[0])
}
