package problem676

import "testing"

func TestMagicDictionary(t *testing.T) {
	// test case 1
	obj := Constructor()
	obj.BuildDict([]string{"hello", "leetcode"})
	if obj.Search("hello") != false {
		t.Errorf("TestMagicDictionary: search \"hello\", expected false, actual true")
	}
	if obj.Search("hhllo") != true {
		t.Errorf("TestMagicDictionary: search \"hhllo\", expected true, actual false")
	}
	if obj.Search("hell") != false {
		t.Errorf("TestMagicDictionary: search \"hell\", expected false, actual true")
	}
	if obj.Search("leetcoded") != false {
		t.Errorf("TestMagicDictionary: search \"leetcoded\", expected false, actual true")
	}
	// test case 2
	obj = Constructor()
	obj.BuildDict([]string{"hello", "hallo", "leetcode", "judge", "judgg"})
	// ["hello"], ["hallo"], ["hell"], ["leetcodd"], ["judge"], ["judgg"]
	if obj.Search("hello") != true {
		t.Errorf("TestMagicDictionary: search \"hello\", expected true, actual false")
	}
	if obj.Search("hallo") != true {
		t.Errorf("TestMagicDictionary: search \"hallo\", expected true, actual false")
	}
	if obj.Search("hell") != false {
		t.Errorf("TestMagicDictionary: search \"hell\", expected false, actual true")
	}
	if obj.Search("leetcodd") != true {
		t.Errorf("TestMagicDictionary: search \"leetcodd\", expected false, actual true")
	}
	if obj.Search("judge") != true {
		t.Errorf("TestMagicDictionary: search \"judge\", expected true, actual false")
	}
	if obj.Search("judgg") != true {
		t.Errorf("TestMagicDictionary: search \"judgg\", expected true, actual false")
	}
}
