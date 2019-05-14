package util

import (
	"fmt"
	"strconv"
)

func ParseAndRemainFloat64(f float64, i int) float64{
	s := "%0." + strconv.FormatInt(int64(i), 10) +"f"
	fr, _ := strconv.ParseFloat(fmt.Sprintf(s, f), 64)
	return fr
}
