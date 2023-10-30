package main

import (
	"reflect"
	"testing"
)

func TestSumAllTails(t *testing.T) {
	t.Run("make the sums of some slices", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{0, 9})
		want := []int{2, 9}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}

	})

	t.Run("safety sum empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{3, 4, 5})
		want := []int{0, 9}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}

	})

}

//reflect 사용이유
// 상대방이 보낸 데이터의 타입 및 종류에 대해 구분이 필요한 경우 사용
