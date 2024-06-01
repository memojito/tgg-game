package systems

import (
	"github.com/EngoEngine/ecs"
	"github.com/EngoEngine/engo/common"
)

const g float32 = 1.5        // g represents the value for gravity
const hGlobal float32 = 1.1  // hGlobal represents the default factor for gravity which increases over time
const hMaxGlobal float32 = 4 // hMaxGlobal represents the max value for the gravity factor hGlobal
const nGlobal int = 10       // nGlobal is a factor that reduces the amount of updates per second

const PlayerSize float32 = 100 // PlayerSize represents the size of a player in pixels.
const TileSize = 32            // TileSize represents the size of a tile in pixels.

// JumpTextMessage gets fired if a Player jumps and is consumed by the physics system
type JumpTextMessage struct {
	ecs.BasicEntity
	common.SpaceComponent
}

const JumpTextMessageType string = "JumpTextMessage"

// Type implements the engo.Message Interface
func (JumpTextMessage) Type() string {
	return JumpTextMessageType
}

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
