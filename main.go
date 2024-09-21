package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	ball "github.com/venndev/VRaylibStarterTemp/src"
)

const (
	ScreenWidth  = 800
	ScreenHeight = 600
	TitleGame    = "Balling"
)

func main() {
	rl.InitWindow(ScreenWidth, ScreenHeight, TitleGame)
	rl.SetTargetFPS(60)

	b := ball.NewBall()
	b.Init() // This is initialization of ball

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)

		b.Update()
		b.Draw()

		rl.EndDrawing()
	}
	rl.CloseWindow()
}
