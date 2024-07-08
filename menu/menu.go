package menu

import (
	"bytes"
	"fmt"
	"image"
	"image/color"
	"log"
	"os"
	collision "snake/collision"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

var (
	BgImage *ebiten.Image

	ButtonWidth  = 100
	ButtonHeight = 40
	ButtonX      = (collision.ScreenW / 2) - (ButtonWidth / 2)
	ButtonY      = (collision.ScreenH / 2) - (ButtonHeight / 2)

	arcadeFaceSource *text.GoTextFaceSource

	SnakeGameIcon *ebiten.Image
)

// Инициализация шрифта
func init() {
	s, err := text.NewGoTextFaceSource(bytes.NewReader(fonts.PressStart2P_ttf))
	if err != nil {
		log.Fatal(err)
	}
	arcadeFaceSource = s
}

// Создание движущегося изображения
func init() {
	openImg, err := os.Open("image/snake_scales.png")
	if err != nil {
		fmt.Print(err)
	}
	defer openImg.Close()

	img, _, err := image.Decode(openImg)
	if err != nil {
		log.Fatal(err)
	}
	BgImage = ebiten.NewImageFromImage(img)
}

type Viewport struct {
	x16 int
	y16 int
}

func (p *Viewport) Move() {
	s := BgImage.Bounds().Size()
	maxY16 := s.Y * 16

	p.y16 += s.Y / 50
	p.y16 %= maxY16
}

func (p *Viewport) Position() (int, int) {
	return p.x16, p.y16
}

// Структура кнопки
type Button struct {
	IsClicked   bool
	Color       color.Color
	TimeClicked time.Time
}

// Функция отрисовки кнопки
func (b *Button) DrawButton(screen *ebiten.Image) {
	mx, my := ebiten.CursorPosition()

	// Если ЛКМ нажата и курсор находится в пределах кнопки, то кнопка срабатывает
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) && ButtonX <= mx && mx <= (ButtonX+ButtonWidth) && ButtonY <= my && my <= (ButtonY+ButtonHeight) {
		b.IsClicked = true
		b.Color = color.RGBA{0, 102, 0, 255}
	} else {
		b.Color = color.RGBA{0, 153, 0, 255}
	}
	vector.DrawFilledRect(screen, float32(ButtonX), float32(ButtonY), float32(ButtonWidth), float32(ButtonHeight), b.Color, false)

	scoreText := "Start"
	op := &text.DrawOptions{}
	op.GeoM.Translate(float64(ButtonX+14), float64(ButtonY+13))
	text.Draw(screen, scoreText, &text.GoTextFace{
		Source: arcadeFaceSource,
		Size:   float64(15),
	}, op)
}

// Добавление изображения
type SnakeIcon struct{}

func init() {
	var err error
	SnakeGameIcon, _, err = ebitenutil.NewImageFromFile("image/snake.png")
	if err != nil {
		log.Fatal(err)
	}
}

// Функция отрисовки изображения
func (s SnakeIcon) DrawSnakeIcon(screen *ebiten.Image, iconPositionX, iconPositionY float64) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(iconPositionX-15, iconPositionY-12)
	screen.DrawImage(SnakeGameIcon, op)
}
