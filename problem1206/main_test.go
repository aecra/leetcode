package problem1206

import "testing"

func TestSkiplist(t *testing.T) {
	skiplist := Constructor()
	skiplist.Add(1)
	skiplist.Add(2)
	skiplist.Add(3)
	if skiplist.Search(0) {
		t.Error("Search(0) should return false")
	}
	skiplist.Add(4)
	if !skiplist.Search(1) {
		t.Error("Search(1) should return true")
	}
	if skiplist.Erase(0) {
		t.Error("Erase(0) should return false")
	}
	if !skiplist.Erase(1) {
		t.Error("Erase(1) should return true")
	}
	if skiplist.Search(1) {

		t.Error("Search(1) should return false")
	}
}
