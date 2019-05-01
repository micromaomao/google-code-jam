package snippets

func Round(f float64) float64 {
	intPart := int64(f)
	f -= float64(intPart)
	if f >= 0.5 {
		return float64(intPart + 1)
	} else {
		return float64(intPart)
	}
}

func Abs(f float64) float64 {
	if f < 0 {
		return -f
	} else {
		return f
	}
}
