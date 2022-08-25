package problem676

type MagicDictionary struct {
	arr []string
}

func Constructor() MagicDictionary {
	return MagicDictionary{}
}

func (magicDictionary *MagicDictionary) BuildDict(dictionary []string) {
	magicDictionary.arr = dictionary
}

func (magicDictionary *MagicDictionary) Search(searchWord string) bool {
	for _, s := range magicDictionary.arr {
		if len(s) != len(searchWord) {
			continue
		}
		if isOneDiff(s, searchWord) {
			return true
		}
	}
	return false
}

func isOneDiff(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	diff := 0
	for i := 0; i < len(s1); i++ {
		if s1[i] != s2[i] {
			diff++
			if diff > 1 {
				return false
			}
		}
	}
	return diff == 1
}
