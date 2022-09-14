package converter

type Converter interface {
	ConvertToInt() (int, error)
}
