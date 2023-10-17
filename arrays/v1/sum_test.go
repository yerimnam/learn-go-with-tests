package main

import "testing"

//arrays have a fixed capacity which you define when you delcare the vairable

// There are two wayd to initilize an array

// 1) [N]type(value1,value2, valueN) eg) numbers := [5]int{1,2,3,4,5}
// 2) [...]type(value1,value2,....valueN) eg)numbers := [...]{1,2,3,4,5}

// %v : the value in a default format when printing structs, the plus flag (%+v) add filed names
func TestSum(t *testing.T) {

	numbers := [5]int{1, 2, 3, 4, 5}

	got := Sum(numbers)
	want := 15

	if got != want {
		t.Errorf("got %d want %d given, %v", got, want, numbers)
	}

}
