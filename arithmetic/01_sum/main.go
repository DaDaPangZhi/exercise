package main

import "fmt"

func sum(dataList []int, target int) {
	dataListLen := len(dataList)
	for aIndex, value := range dataList {
		for bIndex := aIndex + 1; bIndex < dataListLen; bIndex++ {
			if value+dataList[bIndex] == target {
				fmt.Println(aIndex, bIndex)
				return
			}
		}
	}
}

func sum1(dataList []int, target int) {
	resultMap := make(map[int]int, 0)
	for aIndex, value := range dataList {
		if resultValue, exist := resultMap[target-value]; exist {
			fmt.Println(resultValue, aIndex)
			return
		}
		resultMap[value] = aIndex
	}
}

func main() {
	dataList := []int{1, 2, 4, 5, 6}
	target := 5
	sum(dataList, target)
	sum1(dataList, target)
}
