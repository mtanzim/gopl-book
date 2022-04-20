package word2

import "testing"

func Test(t *testing.T) {
	testCases := []struct {
		input string
		want  bool
	}{
		{"", true},
		{"a", true},
		{"ab", false},
		{"kayak", true},
		{"detartrated", true},
		{"a man, a plan, a canal: Panama", true},
		{"palindrome", false},
		{"desserts", false},
	}
	for _, test := range testCases {
		if got := IsPalindrome(test.input); got != test.want {
			t.Errorf("IsPalindrome(%q) = %v", test.input, got)
		}
	}
}
