package systems

import (
	"github.com/EngoEngine/ecs"
	"github.com/EngoEngine/engo"
	"github.com/EngoEngine/engo/common"
	"image/color"
	"log"
	"strings"
)

// Text is an entity containing text printed to the screen
type Text struct {
	ecs.BasicEntity
	common.SpaceComponent
	common.RenderComponent
}

// HUDTextEntity is an entity for the text system. This keeps track of the position
// size and text associated with that position.
type HUDTextEntity struct {
	*ecs.BasicEntity
	*common.SpaceComponent
	*common.MouseComponent
	Line1, Line2, Line3, Line4 string
}

// HUDTextSystem prints the text to our HUD based on the current state of the game
type HUDTextSystem struct {
	text1, text2, text3, text4, money Text

	entities []HUDTextEntity
}

// HUDTextMessage updates the HUD text based on messages sent from other systems
type HUDTextMessage struct {
	ecs.BasicEntity
	common.SpaceComponent
	common.MouseComponent
	Line1, Line2, Line3, Line4 string
}

const HUDTextMessageType string = "HUDTextMessage"

// Type implements the engo.Message Interface
func (HUDTextMessage) Type() string {
	return HUDTextMessageType
}

// centerString centers a string
func centerString(str string, width int) string {
	spaces := int(float64(width-len(str)) / 2)
	return strings.Repeat(" ", spaces) + str + strings.Repeat(" ", width-(spaces+len(str)))
}

// New is called when the system is added to the world.
// Adds text to our HUD that will update based on the state of the game.
func (h *HUDTextSystem) New(w *ecs.World) {
	fnt := &common.Font{
		URL:  "go.ttf",
		FG:   color.Black,
		Size: 24,
	}

	err := fnt.CreatePreloaded()
	if err != nil {
		log.Printf("failed to load preloaded font: %v", err)
		return
	}

	//txt := centerString("yo paul", 24)

	h.text1 = Text{BasicEntity: ecs.NewBasic()}
	h.text1.RenderComponent.Drawable = common.Text{
		Font: fnt,
		Text: "Nothing Selected!",
	}
	h.text1.SetShader(common.TextHUDShader)
	h.text1.RenderComponent.SetZIndex(1001)
	h.text1.SpaceComponent = common.SpaceComponent{
		Position: engo.Point{X: 0, Y: engo.WindowHeight() - 200},
	}
	for _, system := range w.Systems() {
		switch sys := system.(type) {
		case *common.RenderSystem:
			sys.Add(&h.text1.BasicEntity, &h.text1.RenderComponent, &h.text1.SpaceComponent)
		}
	}

	h.text2 = Text{BasicEntity: ecs.NewBasic()}
	h.text2.RenderComponent.Drawable = common.Text{
		Font: fnt,
		Text: "click on an element",
	}
	h.text2.SetShader(common.TextHUDShader)
	h.text2.RenderComponent.SetZIndex(1001)
	h.text2.SpaceComponent = common.SpaceComponent{
		Position: engo.Point{X: 0, Y: engo.WindowHeight() - 180},
	}
	for _, system := range w.Systems() {
		switch sys := system.(type) {
		case *common.RenderSystem:
			sys.Add(&h.text2.BasicEntity, &h.text2.RenderComponent, &h.text2.SpaceComponent)
		}
	}

	h.text3 = Text{BasicEntity: ecs.NewBasic()}
	h.text3.RenderComponent.Drawable = common.Text{
		Font: fnt,
		Text: "to get info",
	}
	h.text3.SetShader(common.TextHUDShader)
	h.text3.RenderComponent.SetZIndex(1001)
	h.text3.SpaceComponent = common.SpaceComponent{
		Position: engo.Point{X: 0, Y: engo.WindowHeight() - 160},
	}
	for _, system := range w.Systems() {
		switch sys := system.(type) {
		case *common.RenderSystem:
			sys.Add(&h.text3.BasicEntity, &h.text3.RenderComponent, &h.text3.SpaceComponent)
		}
	}

	h.text4 = Text{BasicEntity: ecs.NewBasic()}
	h.text4.RenderComponent.Drawable = common.Text{
		Font: fnt,
		Text: "about it.",
	}
	h.text4.SetShader(common.TextHUDShader)
	h.text4.RenderComponent.SetZIndex(1001)
	h.text4.SpaceComponent = common.SpaceComponent{
		Position: engo.Point{X: 0, Y: engo.WindowHeight() - 140},
	}
	for _, system := range w.Systems() {
		switch sys := system.(type) {
		case *common.RenderSystem:
			sys.Add(&h.text4.BasicEntity, &h.text4.RenderComponent, &h.text4.SpaceComponent)
		}
	}

	h.money = Text{BasicEntity: ecs.NewBasic()}
	h.money.RenderComponent.Drawable = common.Text{
		Font: fnt,
		Text: "$0",
	}
	h.money.SetShader(common.TextHUDShader)
	h.money.RenderComponent.SetZIndex(1001)
	h.money.SpaceComponent = common.SpaceComponent{
		Position: engo.Point{X: 0, Y: engo.WindowHeight() - 40},
	}
	for _, system := range w.Systems() {
		switch sys := system.(type) {
		case *common.RenderSystem:
			sys.Add(&h.money.BasicEntity, &h.money.RenderComponent, &h.money.SpaceComponent)
		}
	}

	engo.Mailbox.Listen(HUDTextMessageType, func(m engo.Message) {
		msg, ok := m.(HUDTextMessage)
		if !ok {
			return
		}
		for _, system := range w.Systems() {
			switch sys := system.(type) {
			case *common.MouseSystem:
				sys.Add(&msg.BasicEntity, &msg.MouseComponent, &msg.SpaceComponent, nil)
			case *HUDTextSystem:
				sys.Add(&msg.BasicEntity, &msg.SpaceComponent, &msg.MouseComponent, msg.Line1, msg.Line2, msg.Line3, msg.Line4)
			}
		}
	})

}

// Add adds an entity to the system
func (h *HUDTextSystem) Add(b *ecs.BasicEntity, s *common.SpaceComponent, m *common.MouseComponent, l1, l2, l3, l4 string) {
	h.entities = append(h.entities, HUDTextEntity{b, s, m, l1, l2, l3, l4})
}

// Update is called each frame to update the system.
func (h *HUDTextSystem) Update(dt float32) {}

// Remove takes an enitty out of the system.
// It does nothing as HUDTextSystem has no entities.
func (h *HUDTextSystem) Remove(entity ecs.BasicEntity) {}
