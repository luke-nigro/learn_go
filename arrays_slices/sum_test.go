package main

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("collection of 5 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		got := Sum(numbers)
		want := 15

		if got != want {
			t.Errorf("got %d, expected %d, given %v", got, want, numbers)
		}
	})
}

func TestSumAll(t *testing.T) {
	checkSums := func(t testing.TB, got, want []int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	}
	t.Run("two collections of numbers", func(t *testing.T) {
		got := SumAll([]int{1, 2}, []int{0, 9})
		want := []int{3, 9}
		checkSums(t, got, want)
	})
	t.Run("two collections of numbers, one empty", func(t *testing.T) {
		got := SumAll([]int{}, []int{0, 9})
		want := []int{0, 9}
		checkSums(t, got, want)
	})
}

func TestSumAllTails(t *testing.T) {
	checkSums := func(t testing.TB, got, want int) {
		t.Helper()
		if got != want {
			t.Errorf("got %d, wanted %d", got, want)
		}
	}
	t.Run("two collections of numbers", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{0, 9})
		want := 11
		checkSums(t, got, want)
	})
	t.Run("two collections of numbers, one empty", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{0, 9})
		want := 9
		checkSums(t, got, want)
	})
}
