package systems

import (
	"github.com/EngoEngine/ecs"
	"github.com/EngoEngine/engo"
	"github.com/EngoEngine/engo/common"
	"log"
)

type MouseTracker struct {
	ecs.BasicEntity
	common.MouseComponent
}

type Location struct {
	ecs.BasicEntity
	common.RenderComponent
	common.SpaceComponent
}

type LocationBuildingSystem struct {
	world        *ecs.World
	mouseTracker MouseTracker
}

// Remove is called whenever an Entity is removed from the World, in order to remove it from this system as well
func (*LocationBuildingSystem) Remove(entity ecs.BasicEntity) {
}

// Update is ran every frame, with `dt` being the time
// in seconds since the last frame
func (lb *LocationBuildingSystem) Update(dt float32) {
	if engo.Input.Button("AddLocation").JustPressed() {

		location := Location{BasicEntity: ecs.NewBasic()}
		location.SpaceComponent = common.SpaceComponent{
			Position: engo.Point{lb.mouseTracker.MouseComponent.MouseX, lb.mouseTracker.MouseComponent.MouseY},
			Width:    32,
			Height:   32,
		}

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

// New is the initialisation of the System
func (lb *LocationBuildingSystem) New(w *ecs.World) {
	lb.world = w
	log.Println("LocationBuildingSystem was added to the Scene")

	lb.mouseTracker.BasicEntity = ecs.NewBasic()
	lb.mouseTracker.MouseComponent = common.MouseComponent{Track: true}

	for _, system := range w.Systems() {
		switch sys := system.(type) {
		case *common.MouseSystem:
			sys.Add(&lb.mouseTracker.BasicEntity, &lb.mouseTracker.MouseComponent, nil, nil)
		}
	}
}
