package main

import "fmt"

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


func main() {
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
	camera := Camera {
		CanvasDist: -1,
	}


	screenCoordPoints := make([]Vector2, 8)
	for i, vector := range worldCoordPoints {
		screenCoordPoints[i] = perspectiveDivide(&vector, &camera)
		normalize(&screenCoordPoints[i], &camera)
	}
	fmt.Println(screenCoordPoints)
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

