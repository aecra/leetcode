package problem854

type Item struct {
	Value string
	K     int
}

type Queue []Item

func (q *Queue) Push(i Item) {
	*q = append(*q, i)
}

func (q *Queue) Pop() Item {
	if q.Size() == 0 {
		return Item{}
	}
	res := (*q)[0]
	(*q) = (*q)[1:]
	return res
}

func (q *Queue) Size() int {
	return len(*q)
}

func kSimilarity(s1 string, s2 string) int {
	q := Queue{}
	// remove all the same characters
	sb1, sb2 := []byte{}, []byte{}
	for i := 0; i < len(s1); i++ {
		if s1[i] != s2[i] {
			sb1 = append(sb1, s1[i])
			sb2 = append(sb2, s2[i])
		}
	}
	ss1, ss2 := string(sb1), string(sb2)
	q.Push(Item{ss1, 0})
	for q.Size() != 0 {
		item := q.Pop()
		if item.Value == ss2 {
			return item.K
		}
		for i := 0; i < len(item.Value); i++ {
			if item.Value[i] != ss2[i] {
				for j := i + 1; j < len(item.Value); j++ {
					if item.Value[j] == ss2[i] {
						newValue := item.Value[:i] + string(item.Value[j]) + item.Value[i+1:j] + string(item.Value[i]) + item.Value[j+1:]
						q.Push(Item{newValue, item.K + 1})
					}
				}
				break
			}
		}
	}
	return -1
}
