package systems

import "github.com/EngoEngine/engo/common"

const g float32 = 1.5
const hGlobal float32 = 1.1
const hMaxGlobal float32 = 4
const nGlobal int = 0
const nMaxGlobal int = 10

type Movable struct {
	h    float32
	hMax float32
	n    int
}

// Move changes the X and Y Position of a SpaceComponent by x and y
func Move(sc *common.SpaceComponent, x float32, y float32) {
	sc.Position.X += x
	sc.Position.Y += y
}
