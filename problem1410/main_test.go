package problem1410

import "testing"

func TestEntityParase(t *testing.T) {
	// test case 1
	text := "&amp; is an HTML entity but &ambassador; is not."
	expected := "& is an HTML entity but &ambassador; is not."
	actual := entityParser(text)
	if actual != expected {
		t.Errorf("expected %v, but got %v", expected, actual)
	}

	// test case 2
	text = "and I quote: &quot;...&quot;"
	expected = "and I quote: \"...\""
	actual = entityParser(text)
	if actual != expected {
		t.Errorf("expected %v, but got %v", expected, actual)
	}

	// test case 3
	text = "Stay home! Practice on Leetcode :)"
	expected = "Stay home! Practice on Leetcode :)"
	actual = entityParser(text)
	if actual != expected {
		t.Errorf("expected %v, but got %v", expected, actual)
	}

	// test case 4
	text = "x &gt; y &amp;&amp; x &lt; y is always false"
	expected = "x > y && x < y is always false"
	actual = entityParser(text)
	if actual != expected {
		t.Errorf("expected %v, but got %v", expected, actual)
	}

	// test case 5
	text = "leetcode.com&frasl;problemset&frasl;all"
	expected = "leetcode.com/problemset/all"
	actual = entityParser(text)
	if actual != expected {
		t.Errorf("expected %v, but got %v", expected, actual)
	}
}
