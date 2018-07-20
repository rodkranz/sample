package a

import (
	"testing"
	"fmt"
)

func TestCount(t *testing.T) {
	tests := []struct {
		TextInput         string
		ExpectedVowels    int
		ExpectedConsonant int
	}{
		{
			TextInput:         "teste de consoantes e vogais",
			ExpectedVowels:    11,
			ExpectedConsonant: 13,
		},
		{
			TextInput:         "aaaaaa hdkjdhjjjjjdlska",
			ExpectedVowels:    7,
			ExpectedConsonant: 15,
		},
		{
			TextInput:         "aeiouaeiouaeiouaeiou qrt",
			ExpectedVowels:    20,
			ExpectedConsonant: 3,
		},
		{
			TextInput:         "Lorem Ipsum Dolor Amet",
			ExpectedVowels:    8,
			ExpectedConsonant: 11,
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			ActualConsonant, ActualVowels := Count(test.TextInput)
			if test.ExpectedConsonant != ActualConsonant {
				t.Errorf("For word [%s] is expected %d Consonant, but got %d Consonant", test.TextInput, test.ExpectedConsonant, ActualConsonant)
			}

			if test.ExpectedVowels != ActualVowels {
				t.Errorf("For word [%s] is expected %d Vowels, but got %d Vowels", test.TextInput, test.ExpectedVowels, ActualVowels)
			}
		})
	}
}
