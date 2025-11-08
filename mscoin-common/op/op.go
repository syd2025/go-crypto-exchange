package op

import (
	"fmt"
	"strconv"
)

func DivN(x, y float64, n int) float64 {
	s := fmt.Sprintf("%d", n)
	parseFloat, err := strconv.ParseFloat(fmt.Sprintf("%."+s+"f", x/y), 64)
	if err != nil {
		return 0
	}
	return parseFloat
}

func MulN(x, y float64, n int) float64 {
	s := fmt.Sprintf("%d", n)
	parseFloat, err := strconv.ParseFloat(fmt.Sprintf("%."+s+"f", x*y), 64)
	if err != nil {
		return 0
	}
	return parseFloat
}

func AddN(x, y float64, n int) float64 {
	s := fmt.Sprintf("%d", n)
	parseFloat, err := strconv.ParseFloat(fmt.Sprintf("%."+s+"f", x+y), 64)
	if err != nil {
		return 0
	}
	return parseFloat
}

func SubN(x, y float64, n int) float64 {
	s := fmt.Sprintf("%d", n)
	parseFloat, err := strconv.ParseFloat(fmt.Sprintf("%."+s+"f", x-y), 64)
	if err != nil {
		return 0
	}
	return parseFloat
}
