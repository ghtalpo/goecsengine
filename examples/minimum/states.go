package main

import (
	"github.com/ghtalpo/goecsengine/loader"
	"github.com/ghtalpo/goecsengine/states"
	w "github.com/ghtalpo/goecsengine/world"

	"github.com/hajimehoshi/ebiten/v2"
)

// GameplayState is the main game state
type GameplayState struct{}

// OnPause method
func (st *GameplayState) OnPause(world w.World) {}

// OnResume method
func (st *GameplayState) OnResume(world w.World) {}

// OnStart method
func (st *GameplayState) OnStart(world w.World) {
	loader.LoadEntities("game.toml", world, nil)
}

// OnStop method
func (st *GameplayState) OnStop(world w.World) {
	world.Manager.DeleteAllEntities()
}

// Update method
func (st *GameplayState) Update(world w.World) states.Transition {
	return states.Transition{}
}

// Draw method
func (st *GameplayState) Draw(world w.World, screen *ebiten.Image) {
}
