package main

import "testing"

func TestHowMany9(t *testing.T) {
	want := 9
	got := howMany([]int{3, 7}, "51589")

	if want != got {
		t.Errorf("Expected %v, got %v", want, got)
	}
}

func TestHowMany5(t *testing.T) {
	want := 5
	got := howMany([]int{3, 7}, "01245")

	if want != got {
		t.Errorf("Expected %v, got %v", want, got)
	}
}

func TestHowMany18(t *testing.T) {
	want := 18
	got := howMany([]int{3, 7}, "92510")

	if want != got {
		t.Errorf("Expected %v, got %v", want, got)
	}
}

func TestHowMany2018(t *testing.T) {
	want := 2018
	got := howMany([]int{3, 7}, "59414")

	if want != got {
		t.Errorf("Expected %v, got %v", want, got)
	}
}
