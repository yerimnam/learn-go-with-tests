package main

import (
	"reflect"
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

// reflect.DeepEqual : not type safe
// -> 문자열과 숫자가 슬라이스가 비교가 된다
// -그래서 reflect.DeepEqual은 주의해서 사용하고 슬라이스를 비교할때 편리하다.
func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2}, []int{0, 9})
	want := []int{3, 9}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}
