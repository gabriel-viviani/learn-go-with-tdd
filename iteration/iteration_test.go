package iterators

import "testing"

func TestRepeat(t testing.T) {
	repeated := Repeat("a")
	expected := "aaaaa"

	if repeated != expected {
		t.Errorf("Expect %q but got %q", repeated, expected)
	}
}

func TestRepeatNTimes(t testing.T) {
	repeated := RepeatNTimes("g", 3)
	expected := "ggg"

	if repeated != expected {
		t.Errorf("Expect %q but got %q", repeated, expected)
	}
}
