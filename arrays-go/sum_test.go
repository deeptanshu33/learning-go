package main

import (
	"reflect"
	"slices"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}
		got := Sum(numbers)
		expected := 6

		if got != expected {
			t.Errorf("got %d, expected %d, given %v", got, expected, numbers)
		}
	})
}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2, 3}, []int{3, 4, 5, 6})
	expected := []int{6, 18}

	if !slices.Equal(got, expected) {
		t.Errorf("got %v, expected %v", got, expected)
	}
}

func TestSumAllTails(t *testing.T) {
	got := SumAllTails([]int{0, 2}, []int{0, 9})
	expected := []int{2, 9}

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("got %v, expected %v", got, expected)
	}
}
