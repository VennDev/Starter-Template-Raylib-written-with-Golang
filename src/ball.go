package ball

import rl "github.com/gen2brain/raylib-go/raylib"

const (
	X      int32   = 100
	Y      int32   = 100
	SpeedX int32   = 5
	SpeedY int32   = 5
	Radius float32 = 15.0
)

type Ball struct {
	X      int32
	Y      int32
	SpeedX int32
	SpeedY int32
	Radius float32
}

func (b *Ball) Init() {
	// TODO: Implement this function
}

func (b *Ball) GetInt32Radius() int32 {
	return int32(b.Radius)
}

func (b *Ball) Update() {
	b.X += b.SpeedX
	b.Y += b.SpeedY
	prX := b.X + b.GetInt32Radius()
	prY := b.Y + b.GetInt32Radius()
	drX := b.X - b.GetInt32Radius()
	drY := b.Y - b.GetInt32Radius()
	if prX >= int32(rl.GetScreenWidth()) || drX <= 0 {
		b.SpeedX *= -1
	}
	if prY >= int32(rl.GetScreenHeight()) || drY <= 0 {
		b.SpeedY *= -1
	}
}

func (b *Ball) Draw() {
	rl.DrawCircle(b.X, b.Y, b.Radius, rl.Red)
}

func NewBall() *Ball {
	return &Ball{
		X:      X,
		Y:      Y,
		SpeedX: SpeedX,
		SpeedY: SpeedY,
		Radius: Radius,
	}
}
