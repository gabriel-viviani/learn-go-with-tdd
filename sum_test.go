package main

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("Test any size collection", func(t *testing.T) {
		numbers := []int{1, 2, 4}

		got := Sum(numbers)
		expect := 7

		if got != expect {
			t.Errorf("got %d  want %d given, %v", got, expect, numbers)
		}
	})

}

func TestSumAll(t *testing.T) {
	t.Run("Test sum all in collection", func(t *testing.T) {
		got := SumAll([]int{1, 2}, []int{0, 9})
		want := []int{3, 9}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})
}
