package systems

import (
	"log"

	"github.com/EngoEngine/ecs"
	"github.com/EngoEngine/engo"
	"github.com/EngoEngine/engo/common"
)

// Location describes everything mutable on the screen.
type Location struct {
	ecs.BasicEntity
	common.RenderComponent
	common.SpaceComponent
}

type LocationBuildingSystem struct {
	world *ecs.World
}

// New is the initialisation of the system
func (lb *LocationBuildingSystem) New(w *ecs.World) {
	lb.world = w
	log.Println("LocationBuildingSystem was added to the Scene")
}

// Update the system per frame
func (lb *LocationBuildingSystem) Update(dt float32) {
	if engo.Input.Button("AddLocation").JustPressed() {

		location := Location{BasicEntity: ecs.NewBasic()}

		texture, err := common.LoadedSprite("textures/main-char.png")
		if err != nil {
			log.Printf("failed to load texture: %v", err)
		}

		location.RenderComponent = common.RenderComponent{
			Drawable: texture,
			Scale:    engo.Point{0.1, 0.1},
		}

		for _, system := range lb.world.Systems() {
			switch sys := system.(type) {
			case *common.RenderSystem:
				sys.Add(&location.BasicEntity, &location.RenderComponent, &location.SpaceComponent)
			}
		}

		engo.Mailbox.Dispatch(HUDTextMessage{
			BasicEntity: ecs.NewBasic(),
			SpaceComponent: common.SpaceComponent{
				Position: engo.Point{0.1, 0.1},
				Width:    64,
				Height:   64,
			},
			text: []string{"line1, line2, line3, line4"},
		})
	}
}

func (*LocationBuildingSystem) Remove(entity ecs.BasicEntity) {}
