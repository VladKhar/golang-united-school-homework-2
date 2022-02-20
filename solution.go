package square

import (
	"fmt"
	"math"
)

// Define custom int type to hold sides number and update CalcSquare signature by replacing #yourTypeNameHere#

type intCustomType int

const SidesTriangle = 3
const SidesSquare = 4
const SidesCircle = 0

// Define constants to represent 0, 3 and 4 sides.  Test uses mnemos: SidesTriangle(==3), SidesSquare(==4), SidesCircle(==0)
// it's like:
// CalcSquare(10.0, SidesTriangle)
// CalcSquare(10.0, SidesSquare)
// CalcSquare(10.0, SidesCircle)

func CalcSquare(sideLen float64, sidesNum intCustomType) float64 {
	switch sidesNum {
	case SidesTriangle:
		return (math.Pow(sideLen, 2) * math.Sqrt(3)) / 4.0
	case SidesSquare:
		return math.Pow(sideLen, 2)
	case SidesCircle:
		return math.Pi * math.Pow(sideLen, 2)
	default:
		return 0
	}
}

func mainTest() {
	fmt.Println(CalcSquare(10.0, SidesTriangle))
	fmt.Println(CalcSquare(10.0, SidesSquare))
	fmt.Println(CalcSquare(10.0, SidesCircle))

}
