package problem1268

import (
	"reflect"
	"testing"
)

func TestSuggestedProducts(t *testing.T) {
	// test case 1
	products := []string{"mobile", "mouse", "moneypot", "monitor", "mousepad"}
	searchWord := "mouse"
	expected := [][]string{
		{"mobile", "moneypot", "monitor"},
		{"mobile", "moneypot", "monitor"},
		{"mouse", "mousepad"},
		{"mouse", "mousepad"},
		{"mouse", "mousepad"},
	}
	actual := suggestedProducts(products, searchWord)
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("expected: %v, actual: %v", expected, actual)
	}

	// test case 2
	products = []string{"havana"}
	searchWord = "havana"
	expected = [][]string{
		{"havana"},
		{"havana"},
		{"havana"},
		{"havana"},
		{"havana"},
		{"havana"},
	}
	actual = suggestedProducts(products, searchWord)
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("expected: %v, actual: %v", expected, actual)
	}

	// test case 3
	products = []string{"bags", "baggage", "banner", "box", "cloths"}
	searchWord = "bags"
	expected = [][]string{
		{"baggage", "bags", "banner"},
		{"baggage", "bags", "banner"},
		{"baggage", "bags"},
		{"bags"},
	}
	actual = suggestedProducts(products, searchWord)
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("expected: %v, actual: %v", expected, actual)
	}

	// test case 4
	products = []string{"havana"}
	searchWord = "tatiana"
	expected = [][]string{{}, {}, {}, {}, {}, {}, {}}
	actual = suggestedProducts(products, searchWord)
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("expected: %v, actual: %v", expected, actual)
	}
}
