package day5

import "testing"

func TestMatch(t *testing.T) {
	t.Run("Min", func(t *testing.T) {
		want := 81
		seedRange := seedsRange{
			start: 79,
			end:   93,
		}
		mapRange := mapRange{
			start:       50,
			destination: 52,
			size:        48,
		}
		get := getMin(seedRange, mapRange)
		if want != get {
			t.Fatalf("want %v, got %v", want, get)
		}
	})

	t.Run("Min2", func(t *testing.T) {
		want := 79
		seedRange := seedsRange{
			start: 79,
			end:   93,
		}
		mapRange := mapRange{
			start:       50,
			destination: 98,
			size:        2,
		}
		get := getMin(seedRange, mapRange)
		if want != get {
			t.Fatalf("want %v, got %v", want, get)
		}
	})
}

func TestGetMax(t *testing.T) {
	want := 95
	seedRange := seedsRange{
		start: 79,
		end:   93,
	}
	mapRange := mapRange{
		start:       50,
		destination: 52,
		size:        48,
	}
	get := getMax(seedRange, mapRange)
	if want != get {
		t.Fatalf("want %v, got %v", want, get)
	}
}
