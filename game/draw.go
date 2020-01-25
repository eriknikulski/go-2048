package game

import (
	"fmt"
	"github.com/eriknikulski/2048/config"
	"github.com/eriknikulski/2048/shape"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/text"
	"strconv"
)

func (game *Game) draw() {
	game.window.Clear(config.Background)
	game.imd.Draw(game.window)
	game.drawTiles()
	game.window.Update()
}

func (game *Game) drawBoard() {
	game.imd.Color = config.BigRect
	shape.DrawRoundedRectangle(game.mainRect.Min, game.mainRect.Max, 6, game.imd)
	game.imd.Draw(game.window)
}

func (game *Game) drawTiles() {
	for i := len(game.state.Board) - 1; i >= 0; i-- {
		w := game.state.Board[i]
		for j, v := range w {
			game.drawTile(i, j, v)
		}
	}
}

func (game *Game) drawTile(row int, column int, value int) {
	game.imd.Color = config.Color[strconv.Itoa(value)]
	minVec := pixel.V(
		game.mainRect.Min.X+(float64(column)+1)*config.GridPadding+float64(column)*config.TileRectVec.X,
		game.mainRect.Min.Y+(float64(3-row)+1)*config.GridPadding+float64(3-row)*config.TileRectVec.Y)
	maxVec := pixel.V(minVec.X+config.TileRectVec.X, minVec.Y+config.TileRectVec.Y)
	shape.DrawRoundedRectangle(minVec, maxVec, 3.0, game.imd)

	if value != 0 {
		game.drawText(value, minVec)
	}
}

func (game *Game) drawText(value int, tilePos pixel.Vec) {
	stringValue := strconv.Itoa(value)
	atlas := config.Atlas[stringValue]

	txt := text.New(pixel.ZV, atlas)
	txt.Dot.X = tilePos.X + (config.TileRectVec.X-txt.BoundsOf(stringValue).W())/2
	txt.Dot.Y = tilePos.Y + (config.TileRectVec.Y-txt.BoundsOf(stringValue).H())/1.4
	txt.Color = config.FontColor[strconv.Itoa(value)]

	_, err := fmt.Fprintln(txt, value)
	if err != nil {
		panic(err)
	}

	txt.Draw(game.window, pixel.IM)
}

func (game *Game) drawFinish() {
	// TODO: implement
}
