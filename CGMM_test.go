package main

import (
	"fmt"
	rl "github.com/chunqian/go-raylib/raylib"
	"math"
	"runtime"
	"testing"
)

func init() {
	runtime.LockOSThread()
}

func TestDrawPoints(t *testing.T) {
	screenWidth := int32(xMax)
	screenHeight := int32(yMax)
	rl.InitWindow(screenWidth, screenHeight, "DDA algorithm | Computer Graphics Demo algorithms")
	defer rl.CloseWindow()

	winImage := rl.LoadImage("res/golang-48.png")
	rl.SetWindowIcon(winImage)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)
		rl.DrawPixel(3, 3, rl.Black)
		rl.DrawPixel(10, 10, rl.Blue)
		rl.DrawPixel(15, 15, rl.Brown)
		rl.DrawPixel(20, 20, rl.Beige)
		rl.DrawPixel(25, 25, rl.Pink)
		rl.DrawPixel(30, 30, rl.Green)
		rl.DrawPixel(35, 35, rl.DarkGreen)
		rl.EndDrawing()
	}
	rl.CloseWindow()
}

func TestDDA(t *testing.T) {
	var x1 float32 = 2
	var y1 float32 = 3
	var x2 float32 = 231
	var y2 float32 = 344

	screenWidth := int32(xMax)
	screenHeight := int32(yMax)
	rl.InitWindow(screenWidth, screenHeight, "DDA algorithm | Computer Graphics Demo algorithms")
	defer rl.CloseWindow()
	winImage := rl.LoadImage("res/golang-48.png")
	rl.SetWindowIcon(winImage)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)
		DDA(x1, y1, x2, y2)
		if rl.IsMouseButtonDown(0) {
			x1, y1 = float32(rl.GetMouseX()), float32(rl.GetMouseY())
		}
		if rl.IsMouseButtonPressed(0) {
			x2, y2 = float32(rl.GetMouseX()), float32(rl.GetMouseY())
		}
		rl.EndDrawing()
	}
	rl.CloseWindow()
}


func TestBresenham(t *testing.T) {
	var x1 int32 = 2
	var y1 int32 = 3
	var x2 int32 = 231
	var y2 int32 = 500

	screenWidth := int32(xMax)
	screenHeight := int32(yMax)
	rl.InitWindow(screenWidth, screenHeight, "Bresenham algorithm | Computer Graphics Demo algorithms")
	defer rl.CloseWindow()
	winImage := rl.LoadImage("res/golang-48.png")
	rl.SetWindowIcon(winImage)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)
		Bresenham(x1, y1, x2, y2)
		if rl.IsMouseButtonDown(0) {
			x1, y1 = rl.GetMouseX(), rl.GetMouseY()
		}
		if rl.IsMouseButtonPressed(0) {
			x2, y2 = rl.GetMouseX(), rl.GetMouseY()
		}
		rl.EndDrawing()
	}
	rl.CloseWindow()
}


func TestCircleBresenham(t *testing.T) {
	var (
		xc int32 = 50
		yc int32 = 50
		r int32 = 30
	)
	x1, x2, y1, y2 := rl.GetMouseX(), rl.GetMouseX(), rl.GetMouseY(), rl.GetMouseY()

	screenWidth := int32(xMax)
	screenHeight := int32(yMax)
	rl.InitWindow(screenWidth, screenHeight, "Bresenham algorithm | Computer Graphics Demo algorithms")
	defer rl.CloseWindow()
	winImage := rl.LoadImage("res/golang-48.png")
	rl.SetWindowIcon(winImage)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)
		CircleBresenham(xc, yc, r)
		if rl.IsMouseButtonDown(0) {
			x1, y1 = rl.GetMouseX(), rl.GetMouseY()
			xc, yc = (x1 + x2) / 2, (y1 + y2) / 2
			r = int32(math.Sqrt(math.Pow(float64(x2-x1), 2) + math.Pow(float64(y2-y1), 2)))
		}
		if rl.IsMouseButtonPressed(0) {
			x2, y2 = rl.GetMouseX(), rl.GetMouseY()
		}
		rl.EndDrawing()
	}
	rl.CloseWindow()
}
func TestCohenSutherlandClip(t *testing.T) {
	l := LineInt{-3, -3, 5000, 4440}
	CohenSutherlandClip(l)
	fmt.Printf("Region code for Point 1 (%d, %d) = %04b \n", l.x1, l.y1, computeRegionCode(l.x1, l.y1))
	fmt.Printf("Region code for Point 2 (%d, %d) = %04b \n", l.x2, l.y2, computeRegionCode(l.x2, l.y2))
}
