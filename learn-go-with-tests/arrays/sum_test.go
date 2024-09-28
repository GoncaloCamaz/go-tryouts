package arrays

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("collection of 5 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		got := Sum(numbers)
		expected := 15

		if got != expected {
			t.Errorf("expected %d but got %d given, %v", expected, got, numbers)
		}
	})
	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		got := Sum(numbers)
		expected := 6

		if got != expected {
			t.Errorf("expected %d but got %d given, %v", expected, got, numbers)
		}
	})
}

func TestSumAll(t *testing.T) {
	t.Run("when a collection of arrays is passed", func(t *testing.T) {
		got := SumAll([]int{1, 2}, []int{0, 9})
		expected := []int{3, 9}

		if !reflect.DeepEqual(got, expected) {
			t.Errorf("expected %v but got %v", expected, got)
		}
	})

	t.Run("when an empty array is passed", func(t *testing.T) {
		got := SumAll([]int{})
		expected := []int{0}

		if !reflect.DeepEqual(got, expected) {
			t.Errorf("expected %v but got %v", expected, got)
		}
	})
}

func TestSumAllAppend(t *testing.T) {
	t.Run("when a collection of arrays is passed", func(t *testing.T) {
		got := SumAllAppend([]int{1, 2}, []int{0, 9})
		expected := []int{3, 9}

		if !reflect.DeepEqual(got, expected) {
			t.Errorf("expected %v but got %v", expected, got)
		}
	})

	t.Run("when an empty array is passed", func(t *testing.T) {
		got := SumAllAppend([]int{})
		expected := []int{0}

		if !reflect.DeepEqual(got, expected) {
			t.Errorf("expected %v but got %v", expected, got)
		}
	})
}

func TestSumAllTails(t *testing.T) {
	t.Run("when a collection of arrays is passed", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{0, 9})
		expected := []int{2, 9}

		if !reflect.DeepEqual(got, expected) {
			t.Errorf("expected %v but got %v", expected, got)
		}
	})

	t.Run("when an empty array is passed", func(t *testing.T) {
		got := SumAllTails([]int{})
		expected := []int{0}

		if !reflect.DeepEqual(got, expected) {
			t.Errorf("expected %v but got %v", expected, got)
		}
	})
}
