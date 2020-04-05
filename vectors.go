package main

import (
	"image"
	"fmt"
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

func (vector Vector2) toPoint() image.Point {
	return image.Point{int(vector.X), int(vector.Y)}
}
