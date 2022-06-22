package main

import "fmt"

type listNote struct {
	Value int
	Next  *listNote
}

func sum(listOne, listTwo *listNote) (result *listNote) {

	for listOne.Next != nil {
		fmt.Println(listOne.Value)
		listOne = listOne.Next
	}
	return
}

func main() {
	listOne1 := listNote{Value: 1}
	listOne2 := listNote{Value: 2, Next: &listOne1}
	listOne3 := listNote{Value: 3, Next: &listOne2}

	listTwo1 := listNote{Value: 3}
	listTwo2 := listNote{Value: 2, Next: &listTwo1}
	listTwo3 := listNote{Value: 1, Next: &listTwo2}

	sum(&listOne3, &listTwo3)

}
