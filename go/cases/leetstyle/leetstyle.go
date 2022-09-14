package leetstyle

type LeetStyle string

// Урощенная версия
var alp = map[string]string{
	"E": "3",
	"L": "1",
	"T": "7",
}

func (ls LeetStyle) String() string {
	var result string
	for _, v := range ls {
		s, ok := alp[string(v)]
		if !ok {
			result += string(v)
		} else {
			result += s
		}
	}

	return result
}
