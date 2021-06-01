package main

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {

	t.Run("arrays", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		actual := Sum(numbers)
		expected := 15

		if actual != expected {
			t.Errorf("actual %d , expected %d, given %v", actual, expected, numbers)
		}
	})

	t.Run("slice", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		actual := Sum(numbers)
		expected := 6

		if actual != expected {
			t.Errorf("actual %d , expected %d, given %v", actual, expected, numbers)
		}
	})

	t.Run("add multiple slices", func(t *testing.T) {
		numbers1 := []int{1, 2, 3}
		numbers2 := []int{4, 5, 6}

		actual := SumAll(numbers1, numbers2)
		expected := []int{6, 15}

		if !reflect.DeepEqual(actual, expected) {
			t.Errorf("actual %d , expected %d", actual, expected)
		}
	})

	t.Run("test last values only", func(t *testing.T) {
		numbers1 := []int{1, 2, 3}
		numbers2 := []int{4, 5}

		actual := SumAllTail(numbers1, numbers2)
		expected := []int{3, 5}

		if !reflect.DeepEqual(actual, expected) {
			t.Errorf("actual %d , expected %d", actual, expected)
		}
	})

	t.Run("delete from slice", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

		actual := removeFromSlice(numbers)
		expected := []int{1, 2, 5, 6, 7, 8, 9}

		if len(actual) != len(expected) {
			t.Errorf("got %d want %d", actual, expected)
		}

	})
}
