package main

import "testing"

func TestParseLine(t *testing.T) {
	lines := [][]string{
		{"1*1", "1"},
		{"2*1", "2"},
	}
	want := []Problems{
		{
			q: "1*1",
			a: "1",
		},
		{
			q: "2*1",
			a: "2",
		},
	}
	ret := ParseLines(lines)
	for i, value := range ret {
		if value != want[i] {
			t.Fatalf("want %v but got value %v", want[i], value)
		}
	}

}
