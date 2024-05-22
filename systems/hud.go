package systems

import (
	"image"
	"image/color"
	"log"

	"github.com/EngoEngine/ecs"
	"github.com/EngoEngine/engo"
	"github.com/EngoEngine/engo/common"
)

type HUDSystem struct {
	ecs.BasicEntity
	common.RenderComponent
	common.SpaceComponent
}

// New initializes the system.
func (h *HUDSystem) New(w *ecs.World) {
	log.Println("HUDSystem was added to the Scene")

	hud := HUDSystem{BasicEntity: ecs.NewBasic()}
	hud.SpaceComponent = common.SpaceComponent{
		Position: engo.Point{0, engo.WindowHeight() - 300},
		Width:    200,
		Height:   200,
	}

	hudImage := image.NewUniform(color.RGBA{255, 255, 255, 240})
	hudNRGBA := common.ImageToNRGBA(hudImage, 200, 200)
	hudImageObj := common.NewImageObject(hudNRGBA)
	hudTexture := common.NewTextureSingle(hudImageObj)

	hud.RenderComponent = common.RenderComponent{
		Repeat:   common.Repeat,
		Drawable: hudTexture,
		Scale:    engo.Point{1, 1},
	}
	hud.RenderComponent.SetShader(common.HUDShader)
	hud.RenderComponent.SetZIndex(1)

	for _, system := range w.Systems() {
		switch sys := system.(type) {
		case *common.RenderSystem:
			sys.Add(&hud.BasicEntity, &hud.RenderComponent, &hud.SpaceComponent)
		}
	}
}

func (h *HUDSystem) Remove(entity ecs.BasicEntity) {}
func (h *HUDSystem) Update(dt float32)             {}
