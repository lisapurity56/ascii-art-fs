package main

import "testing"

func TestGetLineIndex(t *testing.T) {
	//'A' is 65. (65-32)*9 + 1 = 298
	got := GetLineIndex('A', 1)
	want := 298
	if got != want {
		t.Errorf("GetLineIndex('A', 1) = %d; want %d", got, want)

	}
}

func TestIsOnlyNewLines(t *testing.T) {
	if !IsOnlyNewLines([]string{"", ""}) {
		t.Error("Expected true for empty newline slice")
	}
}
