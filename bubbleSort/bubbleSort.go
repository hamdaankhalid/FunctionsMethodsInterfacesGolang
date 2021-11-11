package main

import (
	"os"
	"fmt"
	"bufio"
	"strings"
	"strconv"
)
/*
Write a Bubble Sort program in Go. The program
should prompt t9`rters. The program
should print the integers out on one line, in sorted order, from least to
greatest. Use your favorite search tool to find a description of how the bubble
sort algorithm works.

As part of this program, you should write a
function called BubbleSort() which
takes a slice of integers as an argument and returns nothing. The BubbleSort() function should modify the slice so that the elements are in sorted
order.

A recurring operation in the bubble sort algorithm is
the Swap operation which swaps the position of two adjacent elements in the
slice. You should write a Swap() function which performs this operation. Your Swap()
function should take two arguments, a slice of integers and an index value i which
indicates a position in the slice. The Swap() function should return nothing, but it should swap
the contents of the slice in position i with the contents in position i+1.
*/

func main(){
	var numberSlice []int

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Please input upto 10 integers space delmitted that you would like to have sorted:")

	userInput, _ := reader.ReadString('\n')

	userInput = strings.Replace(userInput, "\n", "", -1)
	
	numberSliceString := strings.Split(userInput, " ")

	for _, val := range numberSliceString{
		new_int, err := strconv.Atoi(val)
		if err == nil{
			numberSlice = append(numberSlice, new_int)
		}
	}

	BubbleSort(numberSlice)

	fmt.Println("Sorted via Bubble Sort Algorithm => ", numberSlice)
	
	
}

func BubbleSort(toSort []int){

	var sortedTill int = len(toSort)-1

	// move from left to right swapping if needed then keep iterating while sorted_till is not at 0
	for sortedTill !=0 {
		// a loop that iterates from 0 till sorted_till
		var outerLoop int = 0
		for outerLoop < sortedTill{
			if toSort[outerLoop] > toSort[outerLoop+1]{
				Swap(toSort, outerLoop)
			}
			outerLoop+=1
		}
		sortedTill-=1
	}
}

func Swap(toSort []int, idx int){
	// swap idx with idx+1
	toSort[idx], toSort[idx+1] = toSort[idx+1], toSort[idx]
}
