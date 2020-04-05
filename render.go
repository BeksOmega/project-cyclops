package main

import (
	"image"
	"image/color"
	"image/png"
	"os"
)

type Image interface {
	Set(x, y int, c color.Color)
}

func main() {
	// Init vars.
	cube := Cube{
		Center: Vector3{0, 0, -4},
		Scale: Vector3{1, 3, 2},
	}
	camera := Camera {
		CanvasDist: -1,
		Rect: image.Rect(0, 0, 2, 2),
	}
	imageRect := image.Rect(0, 0, 512, 512)

	// Do work.
	imageCoords := make([]image.Point, 8)
	for i, vector := range cube.Verts() {
		screenPoint := perspectiveDivide(&vector, &camera)
		normalize(&screenPoint, &camera)
		imageCoords[i] = normalToImage(screenPoint, imageRect)
	}

	// Create image.
	img := image.NewGray(imageRect)
	file, _ := os.Create("image.png")
	for _, edge := range cube.Edges() {
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
	vector.X = (vector.X + 1) / float32(camera.Rect.Dx());
	vector.Y = (vector.Y + 1) / float32(camera.Rect.Dy());
}

func normalToImage(v Vector2, rect image.Rectangle) image.Point {
	return image.Point{int(v.X * float32(rect.Dx())), int(v.Y * float32(rect.Dy()))}
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
