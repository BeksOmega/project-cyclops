package main

import (
	"math"
)

var cubeVerts = [8]Vector3 {
	{ 1, -1, -1},
	{ 1, -1,  1},
	{ 1,  1, -1},
	{ 1,  1,  1},
	{-1, -1, -1},
	{-1, -1,  1},
	{-1,  1, -1},
	{-1,  1,  1},
}

type Cube struct {
	Center Vector3
	Scale Vector3
	Rot Vector3
}

func (Cube) Edges() [12][2]int {
	return [12][2]int{
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
}

func (cube Cube) Verts() [8]Vector3 {
	rot := vectorToRadians(cube.Rot)
	rotM := MakeRotationMatrix(&rot)
	scaleM := MakeScaleMatrix(&cube.Scale);
	tranM := MakeTranslateMatrix(&cube.Center);
	mat := rotM.Compose(tranM).Compose(scaleM)
	return [8] Vector3 {
		*mat.Multiply(&cubeVerts[0]),
		*mat.Multiply(&cubeVerts[1]),
		*mat.Multiply(&cubeVerts[2]),
		*mat.Multiply(&cubeVerts[3]),
		*mat.Multiply(&cubeVerts[4]),
		*mat.Multiply(&cubeVerts[5]),
		*mat.Multiply(&cubeVerts[6]),
		*mat.Multiply(&cubeVerts[7]),
	}
}

func vectorToRadians(vec Vector3) Vector3 {
	return Vector3 {
		degreesToRadians(vec.X),
		degreesToRadians(vec.Y),
		degreesToRadians(vec.Z),
	}
}

func degreesToRadians(d float64) float64 {
	return float64(d) * math.Pi / float64(180);
}
