package problem648

import "strings"

/*
 * trie / dictionary tree
 */
func replaceWords(dictionary []string, sentence string) string {
	type trie map[rune]trie
	root := trie{}
	for _, s := range dictionary {
		cur := root
		for _, c := range s {
			if cur[c] == nil {
				cur[c] = trie{}
			}
			cur = cur[c]
		}
		cur['#'] = trie{}
	}

	words := strings.Split(sentence, " ")
	for i, word := range words {
		cur := root
		for j, c := range word {
			if cur['#'] != nil {
				words[i] = word[:j]
				break
			}
			if cur[c] == nil {
				break
			}
			cur = cur[c]
		}
	}
	return strings.Join(words, " ")
}

/*
 * Brute force solution
 */
// func replaceWords(dictionary []string, sentence string) string {
// 	existDictionary := make(map[string]bool, len(dictionary))
// 	for _, root := range dictionary {
// 		existDictionary[root] = true
// 	}
// 	words := strings.Split(sentence, " ")
// 	for i, word := range words {
// 		for j := 1; j < len(word); j++ {
// 			if existDictionary[word[:j]] {
// 				words[i] = word[:j]
// 				break
// 			}
// 		}
// 	}
// 	return strings.Join(words, " ")
// }
