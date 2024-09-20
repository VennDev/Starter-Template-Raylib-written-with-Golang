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

	ball := ball.NewBall()
	ball.Init() // This is initialization of ball

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)

		ball.Update()
		ball.Draw()

		rl.EndDrawing()
	}
	rl.CloseWindow()
}
