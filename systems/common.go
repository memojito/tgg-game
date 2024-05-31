package systems

import "github.com/EngoEngine/engo/common"

// Move changes the X and Y Position of a SpaceComponent by x and y
func Move(sc *common.SpaceComponent, x float32, y float32) {
	sc.Position.X += x
	sc.Position.Y += y
}
