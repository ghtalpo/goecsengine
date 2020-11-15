package main

import (
	r "github.com/ghtalpo/goecsengine/resources"
	s "github.com/ghtalpo/goecsengine/states"
	"github.com/ghtalpo/goecsengine/utils"
	w "github.com/ghtalpo/goecsengine/world"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	windowWidth  = 600
	windowHeight = 600
)

type mainGame struct {
	world        w.World
	stateMachine s.StateMachine
}

func (game *mainGame) Layout(outsideWidth, outsideHeight int) (int, int) {
	return windowWidth, windowHeight
}

func (game *mainGame) Update() error {
	game.stateMachine.Update(game.world)
	return nil
}

func (game *mainGame) Draw(screen *ebiten.Image) {
	game.stateMachine.Draw(game.world, screen)
}

func main() {
	world := w.InitWorld(nil)

	// Init screen dimensions
	world.Resources.ScreenDimensions = &r.ScreenDimensions{Width: windowWidth, Height: windowHeight}

	ebiten.SetWindowResizable(true)
	ebiten.SetWindowSize(windowWidth, windowHeight)
	ebiten.SetWindowTitle("")

	utils.LogError(ebiten.RunGame(&mainGame{world, s.Init(&GameplayState{}, world)}))
}
