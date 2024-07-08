package main

import (
	"fmt"
	"image/color"
	"log"
	collision "snake/collision"
	game "snake/game"
	picture "snake/menu"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font/basicfont"
)

type Game struct {
	Food          *game.Food
	viewport      picture.Viewport
	Button        picture.Button
	Snake         *game.Snake
	score         int
	gameOver      bool
	ticks         int
	updateCounter int
	speed         float64
}

func (g *Game) Layout(w, h int) (int, int) {
	return collision.ScreenW, collision.ScreenH
}

func (g *Game) Update() error {
	if !g.Button.IsClicked {
		g.viewport.Move()
	}
	if g.gameOver {
		if inpututil.IsKeyJustPressed(ebiten.KeyR) {
			g.restart()
		}
		return nil
	}

	if ebiten.IsKeyPressed(ebiten.KeyLeft) && g.Snake.Direction.X == 0 {
		g.Snake.Direction = collision.Vector{X: -1, Y: 0}
	} else if ebiten.IsKeyPressed(ebiten.KeyRight) && g.Snake.Direction.X == 0 {
		g.Snake.Direction = collision.Vector{X: 1, Y: 0}
	} else if ebiten.IsKeyPressed(ebiten.KeyUp) && g.Snake.Direction.Y == 0 {
		g.Snake.Direction = collision.Vector{X: 0, Y: -1}
	} else if ebiten.IsKeyPressed(ebiten.KeyDown) && g.Snake.Direction.Y == 0 {
		g.Snake.Direction = collision.Vector{X: 0, Y: 1}
	}
	g.updateCounter++
	if g.updateCounter < int(g.speed) {
		return nil
	}
	g.updateCounter = 0

	// Обновление состояния змейки
	g.Snake.Move()

	// Столкновение змейки с границами окна
	head := g.Snake.Body[0]
	if head.X < 0 || head.Y < 0 || head.X >= collision.ScreenW/game.TileSize || head.Y >= collision.ScreenH/game.TileSize {
		g.gameOver = true
		g.speed = 35
	}

	// Столкновение головы змейки с ее телом
	for _, part := range g.Snake.Body[1:] {
		if head.X == part.X && head.Y == part.Y {
			g.gameOver = true
			g.speed = 35
		}
	}

	// Когда змейка съедает еду, очки и скорость прибавлются. Появляется новая еда
	if head.X == g.Food.Position.X && head.Y == g.Food.Position.Y {
		g.score++
		g.Snake.GrowCounter += 1
		g.Food = game.NewFood()

		// Уменьшение скорости
		if g.speed > 2 {
			g.speed--
		}
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	if !g.Button.IsClicked {
		// Отрисовка главного меню
		x16, y16 := g.viewport.Position()
		offsetX, offsetY := float64(-x16)/16, float64(-y16)/16

		const repeat = 3
		w, h := picture.BgImage.Bounds().Dx(), picture.BgImage.Bounds().Dy()
		for j := 0; j < repeat; j++ {
			for i := 0; i < repeat; i++ {
				op := &ebiten.DrawImageOptions{}
				op.GeoM.Translate(float64(w*i), float64(h*j))
				op.GeoM.Translate(offsetX, offsetY)
				screen.DrawImage(picture.BgImage, op)
			}
		}

		ebitenutil.DebugPrint(screen, "")

		g.Button.DrawButton(screen)
		snake := &picture.SnakeIcon{}
		snake.DrawSnakeIcon(screen, collision.ScreenW/5, collision.ScreenH/10)
	} else {
		game.DrawBackGround(screen)
		var colors = []color.RGBA{
			{0, 255, 0, 255}, // Зелёный
			{0, 150, 0, 255}, // Красный
		}
		for i, p := range g.Snake.Body {
			color := colors[i%2]
			// Используем остаток от деления для чередования цветов
			game.NewSnake().DrawSnake(screen, float32(p.X*game.TileSize), float32(p.Y*game.TileSize), color)
		}
		game.NewFood().DrawFood(screen, float32(g.Food.Position.X), float32(g.Food.Position.Y))
		face := basicfont.Face7x13

		if g.gameOver {
			text.Draw(screen, "Game Over", face, collision.ScreenW/2-40, collision.ScreenH/2, color.White)
			text.Draw(screen, "Press 'R' to restart", face, collision.ScreenW/2-60, collision.ScreenH/2+16, color.White)
		}

		scoreText := fmt.Sprintf("Score: %d", g.score)
		text.Draw(screen, scoreText, face, 5, collision.ScreenH-5, color.White)
	}
}

// Функция рестарта
func (g *Game) restart() {
	g.Snake = game.NewSnake()
	g.score = 0
	g.gameOver = false
	g.Food = game.NewFood()
	g.speed = 35
}

func main() {
	game := &Game{
		Snake:    game.NewSnake(),
		Food:     game.NewFood(),
		gameOver: false,
		ticks:    0,
		speed:    35,
	}
	ebiten.SetWindowSize(collision.ScreenW*2, collision.ScreenH*2)
	ebiten.SetWindowTitle("Test Snake Game")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
