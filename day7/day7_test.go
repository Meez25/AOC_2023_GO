package day7

import "testing"

func TestSort(t *testing.T) {
	t.Run("case 1", func(t *testing.T) {
		input1 := "32T3K"
		input2 := "KK677"
		want := -1
		got := pokerCmp(input1, input2)
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	})
}

func TestFiveOfAKind(t *testing.T) {
	t.Run("test five of a king", func(t *testing.T) {
		input := "AAAAA"
		want := true
		got := isFiveOfAKind(input)
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	})
	t.Run("test not five of a king", func(t *testing.T) {
		input := "AAAAB"
		want := false
		got := isFiveOfAKind(input)
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	})
}

func TestFourOfAKind(t *testing.T) {
	t.Run("test four of a king", func(t *testing.T) {
		input := "AAAAB"
		want := true
		got := isFourOfAKind(input)
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	})
	t.Run("test not four of a king", func(t *testing.T) {
		input := "AACBA"
		want := false
		got := isFourOfAKind(input)
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	})
}

func TestFullHouse(t *testing.T) {
	t.Run("Test fullHouse", func(t *testing.T) {
		input := "ABAAB"
		want := true
		got := isFullHouse(input)
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	})

	t.Run("Test not fullHouse", func(t *testing.T) {
		input := "BBBAB"
		want := false
		got := isFullHouse(input)
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	})
}

func TestThreeOfAKind(t *testing.T) {
	t.Run("Test threeofakind", func(t *testing.T) {
		input := "ABAAC"
		want := true
		got := isThreeOfAKind(input)
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	})

	t.Run("Test not threeofakind", func(t *testing.T) {
		input := "BBBAB"
		want := false
		got := isThreeOfAKind(input)
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	})
}

func TestIsTwoPair(t *testing.T) {
	t.Run("Test twopairs", func(t *testing.T) {
		input := "BBAAC"
		want := true
		got := isTwoPair(input)
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	})

	t.Run("Test not twopair", func(t *testing.T) {
		input := "BBBAA"
		want := false
		got := isTwoPair(input)
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	})
}

func TestPair(t *testing.T) {
	t.Run("Test pairs", func(t *testing.T) {
		input := "BDAAC"
		want := true
		got := isPair(input)
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	})

	t.Run("Test not pair", func(t *testing.T) {
		input := "BBBAA"
		want := false
		got := isPair(input)
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	})
}

func TestHighCard(t *testing.T) {
	t.Run("Test HighCard", func(t *testing.T) {
		input := "ABCDE"
		want := true
		got := isHighCard(input)
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	})

	t.Run("Test not highcard", func(t *testing.T) {
		input := "BBBAA"
		want := false
		got := isHighCard(input)
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	})
}

func TestEqualValue(t *testing.T) {
	t.Run("Equal card", func(t *testing.T) {
		inputA := "2222A"
		inputB := "22228"
		got := pokerCmp(inputA, inputB)
		want := 1
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	})
}

func TestGoodCase(t *testing.T) {
	t.Run("Test", func(t *testing.T) {
		input := []string{"32T3K 765", "T55J5 684", "KK677 28", "KTJJT 220", "QQQJA 483"}
		got := computeInput(input)
		want := 6440
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	})
}

func TestConvertToBest(t *testing.T) {
	t.Run("Convert to best", func(t *testing.T) {
		input := "T55J5 684"
		got := FindBest(input)
		want := 6
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	})
	t.Run("Convert to best2", func(t *testing.T) {
		input := "KTJJT 220"
		got := FindBest(input)
		want := 6
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	})
}

func TestGoodCaseJoker(t *testing.T) {
	t.Run("Test", func(t *testing.T) {
		input := []string{"32T3K 765", "T55J5 684", "KK677 28", "KTJJT 220", "QQQJA 483"}
		got := computeInputJoker(input)
		want := 5905
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	})
}
