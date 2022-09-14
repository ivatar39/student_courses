package beststring

import (
	"strconv"
)

type BestString string

func (a BestString) ConvertToInt() (int, error) {
	v, err := strconv.Atoi(string(a))
	if err != nil {
		return 0, err
	}
	return v, nil
}
