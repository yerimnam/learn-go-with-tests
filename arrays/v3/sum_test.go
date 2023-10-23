package main

import (
	"testing"
)

func TestSum(t *testing.T) {
	//배열을 크기가 정해져있지만 슬라이스는 크기가 정해져있지 않음
	t.Run("colleciton of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		got := Sum(numbers)
		want := 6

		if got != want {
			t.Errorf("got %d want %d given,%v", got, want, numbers)
		}
	})
}
