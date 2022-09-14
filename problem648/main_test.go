package problem648

import (
	"reflect"
	"testing"
)

func TestReplaceWords(t *testing.T) {
	// test case 1
	dictionary := []string{"cat", "bat", "rat"}
	sentence := "the cattle was rattled by the battery"
	expected := "the cat was rat by the bat"
	actual := replaceWords(dictionary, sentence)
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("expected is %v, actual is %v", expected, actual)
	}

	// test case 2
	dictionary = []string{"a", "b", "c"}
	sentence = "aadsfasf absbs bbab cadsfafs"
	expected = "a a b c"
	actual = replaceWords(dictionary, sentence)
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("expected is %v, actual is %v", expected, actual)
	}
}
