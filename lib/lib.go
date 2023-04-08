package lib

import "strconv"

type ColorCode int
type ToleranceColorCode float64

const (
	Black ColorCode = iota
	Brown
	Red
	Orange
	Yellow
	Green
	Blue
	Violet
	Gray
	White
)

const (
	Gold   ToleranceColorCode = 0.05
	Silver ToleranceColorCode = 0.1
)

func Multiplier(color ColorCode) int {
	if color <= Black {
		return 1
	} else {
		return 10 * Multiplier(color-1)
	}
}

type FourBandResistor struct {
	Band1     ColorCode
	Band2     ColorCode
	Band3     ColorCode
	Tolerance ToleranceColorCode
}

func (r *FourBandResistor) IdealValue() int {
	var prefix_value string = strconv.Itoa(int(r.Band1)) + strconv.Itoa(int(r.Band2))
	if lead_value, err := strconv.Atoi(prefix_value); err == nil {
		return lead_value * Multiplier(r.Band3)
	}
	return 0
}

func (r *FourBandResistor) ValueRange() (min float64, max float64) {
	idealValue := r.IdealValue()
	min = float64(idealValue) - float64(r.Tolerance)
	max = float64(idealValue) + float64(r.Tolerance)
	return
}
