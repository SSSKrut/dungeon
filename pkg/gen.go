package dungeon

import (
	"github.com/butvinm/dungeon/pkg/math"
)

type RoomSpec struct {
	Name string
	Size math.Vector2
}

type Room struct {
	Spec  RoomSpec
	Pos   math.Vector2
	Nbors []Room
}

type Generator struct {
	roomSpecs         []RoomSpec
	roomProbabilities map[string]float32
	// ForbidPatterns    []Pattern
	generated map[string]Room
}

func (g *Generator) getRoomById(id string) Room {
	room, ok := g.generated[id]
	if ok {
		return room
	}
	return Room{
		Spec: g.roomSpecs[0],
		Pos:  math.Vector2{X: 0, Y: 0},
	}
}
