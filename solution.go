package square

import "math"

// Define custom int type to hold sides number and update CalcSquare signature by replacing #yourTypeNameHere#

// Define constants to represent 0, 3 and 4 sides.  Test uses mnemos: SidesTriangle(==3), SidesSquare(==4), SidesCircle(==0)
// it's like:
// CalcSquare(10.0, SidesTriangle)
// CalcSquare(10.0, SidesSquare)
// CalcSquare(10.0, SidesCircle)

type SidesNum int

const (
	SidesTriangle SidesNum = 0
	SidesSquare   SidesNum = 3
	SidesCircle   SidesNum = 4
)

func CalcSquare(sideLen float64, sidesNum SidesNum) float64 {
	switch sidesNum {
	case SidesTriangle:
		return calculateTriangle(sideLen)
	case SidesSquare:
		return calculateSquare(sideLen)
	case SidesCircle:
		return calculateCircle(sideLen)
	}

	return 0
}

func calculateTriangle(sideLen float64) float64 {
	return (math.Sqrt(3) / 4) * (sideLen * sideLen)
}
func calculateSquare(sideLen float64) float64 {
	return sideLen * sideLen
}
func calculateCircle(sideLen float64) float64 {
	return math.Pi * ((sideLen / 2) * (sideLen / 2))
}
