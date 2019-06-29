package utils

import "strconv"

// StrToFloat32 : converts string to float
func StrToFloat32(f string) (floatString float32, err error) {
	s, err := strconv.ParseFloat(f, 32)
	if err == nil {
		floatString = float32(s)
	}
	return
}

// StrToInt returns integer value of str
func StrToInt(str string) (value int, err error) {
	value, err = strconv.Atoi(str)
	return
}

// Abs is used to get the positive absolute value of a float32 value
func Abs(x float32) float32 {
	if x < 0 {
		return -x
	}
	return x
}
