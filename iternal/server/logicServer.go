package server

import "math"

type Coor struct {
	x float32
	y float32
}

var shiftXb float32 = 10.5
var shiftXt float32 = 32

var shiftYb float32 = 10.5
var shiftYt float32 = 32

func CheckHit(bx float32, by float32, tx float32, ty float32) bool {
	bulletC := Coor{bx, by}
	tankC := Coor{tx, ty}

	oxCma := math.Max(float64(bulletC.x), float64(tankC.x))
	oyCma := math.Max(float64(bulletC.y), float64(tankC.y))

	oxCmi := math.Min(float64(bulletC.x+shiftXb), float64(tankC.x+shiftXt))
	oyCmi := math.Max(float64(bulletC.y+shiftYb), float64(tankC.y+shiftYt))

	if oxCma < oxCmi && oyCma < oyCmi {
		return true
	} else {
		return false
	}
}
