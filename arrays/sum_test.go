package main

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {

	t.Run("sum of any size", func(t *testing.T) {
		mySlice := []int{1, 2, 3}
		got := Sum(mySlice)
		want := 6
		if got != want {
			t.Errorf("got '%d', want '%d' given input %v", got, want, mySlice)
		}
	})

}

func TestSumAllTails(t *testing.T) {
	checkSums := func(t testing.TB, got []int, want []int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v given input", got, want)
		}
	}
	t.Run("sum of slices", func(t *testing.T) {
		slice1 := []int{1, 2, 3}
		slice2 := []int{0, 9, 3}
		got := SumAll(slice1, slice2)
		want := []int{5, 12}

		checkSums(t, got, want)
	})

	t.Run("safely sum empty slices", func(t *testing.T) {
		slice1 := []int{1, 2, 3}
		slice2 := []int{}
		got := SumAll(slice1, slice2)
		want := []int{5, 0}
		checkSums(t, got, want)

	})

}
