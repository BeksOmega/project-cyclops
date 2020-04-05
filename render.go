package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
)

type Vector3 struct {
	X float32
	Y float32
	Z float32
}

func (vector Vector3) String() string {
	return fmt.Sprintf("{X: %v, Y: %v, Z: %v}", vector.X, vector.Y, vector.Z)
}

type Vector2 struct {
	X float32
	Y float32
}

func (vector Vector2) String() string {
	return fmt.Sprintf("{x: %v, Y: %v}", vector.X, vector.Y)
}

type Camera struct {
	CanvasDist float32
}

type Image interface {
	Set(x, y int, c color.Color)
}


func main() {
	// Init vars.
	worldCoordPoints := []Vector3{
		{ 1, -1, -5},
		{ 1, -1, -3},
		{ 1,  1, -5},
		{ 1,  1, -3},
		{-1, -1, -5},
		{-1, -1, -3},
		{-1,  1, -5},
		{-1,  1, -3},
	}
	edges := [][]int{
		{0, 1},
		{1, 3},
		{3, 2},
		{2, 0},
		{4, 5},
		{5, 7},
		{7, 6},
		{6, 4},
		{0, 4},
		{1, 5},
		{3, 7},
		{2, 6},
	}
	camera := Camera {
		CanvasDist: -1,
	}

	// Do work.
	imageCoords := make([]image.Point, 8)
	for i, vector := range worldCoordPoints {
		screenPoint := perspectiveDivide(&vector, &camera)
		normalize(&screenPoint, &camera)
		imageCoords[i] = normalToImage(screenPoint, 512, 512)
	}

	// Create image.
	img := image.NewGray(image.Rect(0, 0, 512, 512))
	file, _ := os.Create("image.png")
	for _, edge := range edges {
		Line(imageCoords[edge[0]], imageCoords[edge[1]], img)
	}
	png.Encode(file, img)
}

// perspectiveDivide projects the point onto the canvas by multiplying the
// X and Y coordinates of the point by the depth (distance of the canvas from
// the POV) and then dividing that by the distance of the point from the POV
// (in the z dimension)
func perspectiveDivide(vector *Vector3, camera *Camera) Vector2 {
	return Vector2{
		X: vector.X * camera.CanvasDist / vector.Z,
		Y: vector.Y * camera.CanvasDist / vector.Z,
	}
}

func normalize(vector *Vector2, camera *Camera) {
	// TODO: Change '2' to instead use camera.
	vector.X = (vector.X + 1) / 2;
	vector.Y = (vector.Y + 1) / 2;
}

func normalToImage(v Vector2, x, y float32) image.Point {
	return image.Point{int(v.X * x), int(v.Y * y)}
}

func Line(a, b image.Point, img Image) {
	dx := a.X - b.X
	dy := a.Y - b.Y
	if dx < 0 {
		dx = -dx
	}
	if dy < 0 {
		dy = -dy
	}

	incX, incY := -1, -1
	if a.X < b.X {
		incX = 1
	}
	if a.Y < b.Y {
		incY = 1
	}

	err := dx - dy
	for {
		img.Set(a.X, a.Y, color.Gray{uint8(255)})
		if a.X == b.X && a.Y == b.Y {
			break
		}
		e2 := 2 * err
		if e2 > -dx {
			err -= dy
			a.X += incX
		}
		if e2 < dx {
			err += dx
			a.Y += incY
		}
	}
}
