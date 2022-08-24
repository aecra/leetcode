package problem557

import "testing"

// 557. Reverse Words in a String III
// 557. 反转字符串中的单词 III
func TestReverseWords(t *testing.T) {
	// test case 1
	got := reverseWords("Let's take LeetCode contest")
	if got != "s'teL ekat edoCteeL tsetnoc" {
		t.Errorf("reverseWords() error %s.", got)
	}

	// test case 2
	got = reverseWords("God Ding")
	if got != "doG gniD" {
		t.Errorf("reverseWords() error %s.", got)
	}
}
