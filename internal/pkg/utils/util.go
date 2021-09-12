package utils

func HashCode(s string) int32 {
	var h int32
	for _, aChar := range s {
		h = h*31 + aChar
	}
	return h
}
