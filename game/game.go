package game

import (
	"image/color"
	collision "snake/collision"

	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

const (
	TileSize = 10
)

// Структура змейки
type Snake struct {
	Direction   collision.Vector
	Body        []collision.Vector
	GrowCounter int
}

func NewSnake() *Snake {
	return &Snake{
		Body: []collision.Vector{
			{X: collision.ScreenW / TileSize / 2, Y: collision.ScreenH / TileSize / 2},
		},
		Direction: collision.Vector{X: 1, Y: 0},
	}
}

// Движение змейки
func (s *Snake) Move() {
	newHead := collision.Vector{
		X: s.Body[0].X + s.Direction.X,
		Y: s.Body[0].Y + s.Direction.Y,
	}
	s.Body = append([]collision.Vector{newHead}, s.Body...)

	if s.GrowCounter > 0 {
		s.GrowCounter--
	} else {
		s.Body = s.Body[:len(s.Body)-1]
	}
}

// Отрисовка змейки
func (s *Snake) DrawSnake(screen *ebiten.Image, X, Y float32, color color.Color) {
	vector.DrawFilledRect(screen, X, Y, TileSize, TileSize, color, false)

}

// Структура еды
type Food struct {
	Position collision.Vector
}

func NewFood() *Food {
	return &Food{
		Position: collision.Vector{
			X: rand.Intn(collision.ScreenW / TileSize),
			Y: rand.Intn(collision.ScreenH / TileSize),
		},
	}
}

// Отрисовка еды
func (f *Food) DrawFood(screen *ebiten.Image, X, Y float32) {
	vector.DrawFilledRect(screen, float32(X*TileSize), float32(Y*TileSize), TileSize, TileSize, color.RGBA{255, 0, 0, 255}, false)
}

// Отрисовка заднего экрана (сетка)
func DrawBackGround(screen *ebiten.Image) {
	const w = collision.ScreenW
	const h = collision.ScreenH

	vector.DrawFilledRect(screen, 0, 0, w, h, color.Black, false)

	gridColor64 := &color.RGBA{G: 20}
	for y := 0.0; y < h; y += TileSize {
		vector.StrokeLine(screen, 0, float32(y), w, float32(y), 1, gridColor64, false)
	}
	for y := 0.0; y < h; y += TileSize * 2 {
		vector.StrokeLine(screen, 0, float32(y), w, float32(y), 1, gridColor64, false)
	}
	for x := 0.0; x < w; x += TileSize {
		vector.StrokeLine(screen, float32(x), 0, float32(x), h, 1, gridColor64, false)
	}
	for x := 0.0; x < w; x += TileSize * 2 {
		vector.StrokeLine(screen, float32(x), 0, float32(x), h, 1, gridColor64, false)
	}
}
