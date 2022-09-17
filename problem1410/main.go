package problem1410

func entityParser(text string) string {
	chars := []byte{}
	for i := 0; i < len(text); i++ {
		if text[i] == '&' {
			if i+5 < len(text) && text[i+1:i+6] == "quot;" {
				chars = append(chars, '"')
				i += 5
				continue
			}
			if i+5 < len(text) && text[i+1:i+6] == "apos;" {
				chars = append(chars, '\'')
				i += 5
				continue
			}
			if i+4 < len(text) && text[i+1:i+5] == "amp;" {
				chars = append(chars, '&')
				i += 4
				continue
			}
			if i+3 < len(text) && text[i+1:i+4] == "gt;" {
				chars = append(chars, '>')
				i += 3
				continue
			}
			if i+3 < len(text) && text[i+1:i+4] == "lt;" {
				chars = append(chars, '<')
				i += 3
				continue
			}
			if i+6 < len(text) && text[i+1:i+7] == "frasl;" {
				chars = append(chars, '/')
				i += 6
				continue
			}
		}
		chars = append(chars, text[i])
	}
	return string(chars)
}
