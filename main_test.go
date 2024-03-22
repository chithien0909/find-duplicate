package main

import (
	"testing"
)

var testCases = []struct {
	name     string
	str      string
	expected []string
}{

	{
		name:     "Test 1",
		str:      "When we started Slite as a notes app for teams, we didn't know that we'd become fanatical about wikis. As our team grew from a single home base in Paris, to a constellation of offices around the world, we realized that we needed to share more than notes. We needed a home base for gathering, organizing, and distributing team knowledge. Now, our Slite contains ideas, projects, discussions, sketches, guidebooks, and glossaries. It's basically our collective brain.",
		expected: []string{"e", "a"},
	},
	{
		name:     "Test 2",
		str:      "comment on the main question, he was not worried about escape characters. Nor were they mentioned in the question. This would likely have been the best solution for the problem.",
		expected: []string{"e", "t"},
	},
	{
		name:     "Test 3",
		str:      "Of course, I can use strings.ReplaceAll to manually convert some of these, but I'd also like arbitrary unicode escapes to be translated to the unicode glyph, and so on. Is there a standard library or 3rd party library for that? Input is not trusted, it needs to be safe.",
		expected: []string{"e", "t"},
	},
}

func Test_FindTopDuplicates1(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := FindTopDuplicates1(tc.str)
			if len(tc.expected) != len(result) {
				t.Errorf("%s: expected %v, got %v", tc.name, tc.expected, result)
			}
			for i := 0; i < len(tc.expected); i++ {
				if tc.expected[i] != result[i] {
					t.Errorf("%s: expected %v, got %v", tc.name, tc.expected, result)
				}
			}
		})
	}
}

func Test_FindTopDuplicates2(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := FindTopDuplicates2(tc.str)
			if len(tc.expected) != len(result) {
				t.Errorf("%s: expected %v, got %v", tc.name, tc.expected, result)
			}
			for i := 0; i < len(tc.expected); i++ {
				if tc.expected[i] != result[i] {
					t.Errorf("%s: expected %v, got %v", tc.name, tc.expected, result)
				}
			}
		})
	}
}
