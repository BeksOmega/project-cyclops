package main

import (
	"image"
	"image/color"
	"image/png"
	"os"
	"fmt"
)

type Image interface {
	Set(x, y int, c color.Color)
}

func main() {

	matrixA := Matrix4x4{
		entries: [4][4]float64{
			{2, 0, 0, 0},
			{0, 1, 0, 0},
			{0, 0, 1, 0},
			{4, 4, 4, 1},
		},
	}

	/*matrixB := Matrix4x4{
		entries: [4][4]float64{
			{2, 0, 0, 0},
			{0, 1, 0, 0},
			{0, 0, 1, 0},
			{0, 0, 0, 1},
		},
	}*/

	vectorA := Vector3{ 1, 0, 0 }

	//matrixB = *matrixB.Compose(&matrixA);
	//fmt.Print(matrixB.ToString());

	vectorA = *matrixA.Multiply(&vectorA);
	fmt.Print(vectorA.String());

	// Init vars.
	cube := Cube{
		Center: Vector3{0, 0, -4},
		Scale: Vector3{2, 2, 2},
		Rot: Vector3{0, 0, 0},  // Degrees.
	}
	camera := Camera {
		CanvasDist: -1,
		Rect: image.Rect(0, 0, 2, 2),
	}
	imageRect := image.Rect(0, 0, 512, 512)

	// Do work.
	rasterCoords := make([]image.Point, 8)
	for i, vector := range cube.Verts() {
		screenPoint := perspectiveDivide(&vector, &camera)
		normalize(&screenPoint, &camera)
		rasterCoords[i] = normalToImage(screenPoint, imageRect)
	}

	// Create image.
	img := image.NewGray(imageRect)
	file, _ := os.Create("image.png")
	for _, edge := range cube.Edges() {
		Line(rasterCoords[edge[0]], rasterCoords[edge[1]], img)
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
	vector.X = (vector.X + 1) / float64(camera.Rect.Dx());
	vector.Y = (vector.Y + 1) / float64(camera.Rect.Dy());
}

func normalToImage(v Vector2, rect image.Rectangle) image.Point {
	return image.Point{int(v.X * float64(rect.Dx())), int(v.Y * float64(rect.Dy()))}
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
