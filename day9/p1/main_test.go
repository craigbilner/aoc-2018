package main

import (
	"testing"
)

func TestHighScore32(t *testing.T) {
	input := "9 players; last marble is worth 25 points"
	result := highScoreFromText(input)

	if result != 32 {
		t.Error("Expected 32 got ", result)
	}
}

func TestHighScore8317(t *testing.T) {
	input := "10 players; last marble is worth 1618 points"
	result := highScoreFromText(input)

	if result != 8317 {
		t.Error("Expected 8317 got ", result)
	}
}

func TestHighScore146373(t *testing.T) {
	input := "13 players; last marble is worth 7999 points"
	result := highScoreFromText(input)

	if result != 146373 {
		t.Error("Expected 146373 got ", result)
	}
}

func TestHighScore2764(t *testing.T) {
	input := "17 players; last marble is worth 1104 points"
	result := highScoreFromText(input)

	if result != 2764 {
		t.Error("Expected 2764 got ", result)
	}
}

func TestHighScore54718(t *testing.T) {
	input := "21 players; last marble is worth 6111 points"
	result := highScoreFromText(input)

	if result != 54718 {
		t.Error("Expected 54718 got ", result)
	}
}

func TestHighScore37305(t *testing.T) {
	input := "30 players; last marble is worth 5807 points"
	result := highScoreFromText(input)

	if result != 37305 {
		t.Error("Expected 37305 got ", result)
	}
}
