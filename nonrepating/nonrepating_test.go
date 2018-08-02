package main

import "testing"

func TestSubstr(t *testing.T)  {
	test := []struct{
		a string
		ans int
	}{
		{"asdfasdfad", 4},
		{"aaaaaaaaaa", 1},
	}

	for _, tt := range test{
		if ans := lengthOfNonRepeatingSubStr(tt.a); ans != tt.ans {
			t.Errorf("Got %d for input %s;" +
				"expected %d", ans, tt.a, tt.ans)
		}
	}
}

func BenchmarkSubstr(b *testing.B) {
	s, ans := "SasdfasjkdfhajskdfhasdADsaffgddsgghhjhjj", 7

	for i := 0; i < b.N; i++ {
		actual := lengthOfNonRepeatingSubStr(s)
		if actual != ans {
			b.Errorf("Got %d for input %s;"+
				"expected %d", actual, s, ans)
		}
	}
}