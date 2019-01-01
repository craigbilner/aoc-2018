package main

import "testing"

func TestScore9(t *testing.T) {
	want := "5158916779"
	got := score([]int{3, 7}, 9)

	if want != got {
		t.Errorf("Expected %v, got %v", want, got)
	}
}

func TestScore5(t *testing.T) {
	want := "0124515891"
	got := score([]int{3, 7}, 5)

	if want != got {
		t.Errorf("Expected %v, got %v", want, got)
	}
}

func TestScore18(t *testing.T) {
	want := "9251071085"
	got := score([]int{3, 7}, 18)

	if want != got {
		t.Errorf("Expected %v, got %v", want, got)
	}
}

func TestScore2018(t *testing.T) {
	want := "5941429882"
	got := score([]int{3, 7}, 2018)

	if want != got {
		t.Errorf("Expected %v, got %v", want, got)
	}
}
