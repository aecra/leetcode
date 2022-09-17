package problem65

type State int
type CharType int

const (
	STATE_START State = iota
	STATE_INT_SIGN
	STATE_INT
	STATE_DECIMAL_WITHOUT_FRACTION
	STATE_POINT
	STATE_DECIMAL
	STATE_EXP
	STATE_EXP_SIGN
	STATE_EXP_NUMBER
)

const (
	CHAR_NUMBER CharType = iota
	CHAR_EXP
	CHAR_POINT
	CHAR_SIGN
	CHAR_ILLEGAL
)

func toCharType(ch byte) CharType {
	switch ch {
	case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
		return CHAR_NUMBER
	case 'e', 'E':
		return CHAR_EXP
	case '.':
		return CHAR_POINT
	case '+', '-':
		return CHAR_SIGN
	default:
		return CHAR_ILLEGAL
	}
}

func isNumber(s string) bool {
	transfer := map[State]map[CharType]State{
		STATE_START: {
			CHAR_NUMBER: STATE_INT,
			CHAR_POINT:  STATE_POINT,
			CHAR_SIGN:   STATE_INT_SIGN,
		},
		STATE_INT_SIGN: {
			CHAR_NUMBER: STATE_INT,
			CHAR_POINT:  STATE_POINT,
		},
		STATE_INT: {
			CHAR_NUMBER: STATE_INT,
			CHAR_EXP:    STATE_EXP,
			CHAR_POINT:  STATE_DECIMAL_WITHOUT_FRACTION,
		},
		STATE_DECIMAL_WITHOUT_FRACTION: {
			CHAR_NUMBER: STATE_DECIMAL,
			CHAR_EXP:    STATE_EXP,
		},
		STATE_POINT: {
			CHAR_NUMBER: STATE_DECIMAL,
		},
		STATE_DECIMAL: {
			CHAR_NUMBER: STATE_DECIMAL,
			CHAR_EXP:    STATE_EXP,
		},
		STATE_EXP: {
			CHAR_NUMBER: STATE_EXP_NUMBER,
			CHAR_SIGN:   STATE_EXP_SIGN,
		},
		STATE_EXP_SIGN: {
			CHAR_NUMBER: STATE_EXP_NUMBER,
		},
		STATE_EXP_NUMBER: {
			CHAR_NUMBER: STATE_EXP_NUMBER,
		},
	}
	state := STATE_START
	for i := 0; i < len(s); i++ {
		charType := toCharType(s[i])
		if _, ok := transfer[state][charType]; !ok {
			return false
		} else {
			state = transfer[state][charType]
		}
	}
	return state == STATE_INT || state == STATE_DECIMAL_WITHOUT_FRACTION || state == STATE_DECIMAL || state == STATE_EXP_NUMBER
}
