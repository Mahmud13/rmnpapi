package helpers

import "math"

func Rad2deg(r float64) float64{
	return r/math.Pi*180
}
func Deg2rad(d float64) float64{
	return d * math.Pi/180
}
