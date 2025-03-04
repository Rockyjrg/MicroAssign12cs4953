package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	rl.InitWindow(800, 400, "Micro Assignment 12 - Playing Audio!")
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	//need to initialize sound device
	rl.InitAudioDevice()

	//set sound effects
	umbrellaOpen := rl.LoadSound("Audio/umbrella.wav")
	flute := rl.LoadSound("Audio/flute.wav")
	dontKnow := rl.LoadSound("Audio/dontKnowWhatToCall.wav")
	piano := rl.LoadSound("Audio/pianoSingleNote.wav")
	a4Mono := rl.LoadSound("Audio/a4Mono.wav")

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)

		if rl.IsKeyPressed(rl.KeyOne) {
			rl.PlaySound(umbrellaOpen)
		}
		if rl.IsKeyPressed(rl.KeyTwo) {
			rl.PlaySound(flute)
		}
		if rl.IsKeyPressed(rl.KeyThree) {
			rl.PlaySound(dontKnow)
		}
		if rl.IsKeyPressed(rl.KeyFour) {
			rl.PlaySound(piano)
		}
		if rl.IsKeyPressed(rl.KeyFive) {
			rl.PlaySound(a4Mono)
		}

		rl.EndDrawing()
	}
}
