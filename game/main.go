package game

import (
	"fmt"
	"github.com/eriknikulski/2048/config"
	"github.com/eriknikulski/2048/logic"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
)

type Game struct {
	window   *pixelgl.Window
	state    *logic.GameState
	imd      *imdraw.IMDraw
	mainRect pixel.Rect
}

func NewGame() *Game {
	win, err := pixelgl.NewWindow(config.Cfg)
	if err != nil {
		panic(err)
	}

	return &Game{
		window: win,
		state:  logic.NewGameState(),
		imd:    imdraw.New(nil),
		mainRect: pixel.R(
			(config.WindowSize.X-config.MainRectSize.X)/2,
			(config.WindowSize.Y-config.MainRectSize.Y)/2,
			(config.WindowSize.X-config.MainRectSize.X)/2+config.MainRectSize.X,
			(config.WindowSize.Y-config.MainRectSize.Y)/2+config.MainRectSize.Y),
	}
}

func (game *Game) moveOnInput() (finished bool, err error) {
	if game.window.JustPressed(pixelgl.KeyA) || game.window.JustPressed(pixelgl.KeyLeft) {
		finished, err = game.state.Move(logic.LEFT)
		if err != nil {
			return
		}
	}

	if game.window.JustPressed(pixelgl.KeyD) || game.window.JustPressed(pixelgl.KeyRight) {
		finished, err = game.state.Move(logic.RIGHT)
		if err != nil {
			return
		}
	}

	if game.window.JustPressed(pixelgl.KeyW) || game.window.JustPressed(pixelgl.KeyUp) {
		finished, err = game.state.Move(logic.UP)
		if err != nil {
			return
		}
	}

	if game.window.JustPressed(pixelgl.KeyS) || game.window.JustPressed(pixelgl.KeyDown) {
		finished, err = game.state.Move(logic.DOWN)
		if err != nil {
			return
		}
	}
	return
}

func (game *Game) gameLoop() error {
	// TODO: draw only on input
	for !game.window.Closed() {
		finished, err := game.moveOnInput()
		if err != nil {
			return err
		}
		if finished {
			break
		}
		game.draw()
		game.window.SetTitle(fmt.Sprintf("%s | Score: %d", config.Cfg.Title, game.state.Score))
	}

	for !game.window.Closed() {
		game.drawFinish()
		game.window.Update()
	}
	return nil
}

func Run() {
	game := NewGame()
	game.drawBoard()
	err := game.gameLoop()
	if err != nil {
		panic(err)
	}
}
