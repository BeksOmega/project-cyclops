package main

import (
	"fmt"
	"math"
)

type Matrix4x4 struct {
	entries [4][4]float64
}

func (mat *Matrix4x4) Get(r, c int) float64 {
	return mat.entries[r][c]
}

func MakeScaleMatrix(s *Vector3) *Matrix4x4 {
	return &Matrix4x4{
		entries: [4][4]float64{
			{s.X,   0,   0, 0},
			{  0, s.Y,   0, 0},
			{  0,   0, s.Z, 0},
			{  0,   0,   0, 1},
		},
	}
}

func MakeTranslateMatrix(t *Vector3) *Matrix4x4 {
	return &Matrix4x4{
		entries: [4][4]float64{
			{  1,   0,   0, 0},
			{  0,   1,   0, 0},
			{  0,   0,   1, 0},
			{t.X, t.Y, t.Z, 1},
		},
	}
}

// Axis rotation, not point rotation.
func MakeRotationMatrix(r *Vector3) *Matrix4x4 {
	xC := math.Cos(r.X)
	xS := math.Sin(r.X)
	xMat := &Matrix4x4{
		entries: [4][4]float64{
			{  1,   0,   0, 0},
			{  0,  xC, -xS, 0},
			{  0,  xS,  xC, 0},
			{  0,   0,   0, 1},
		},
	}

	yC := math.Cos(r.Y)
	yS := math.Sin(r.Y)
	yMat := &Matrix4x4{
		entries: [4][4]float64{
			{ yC,   0,  yS, 0},
			{  0,   1,   0, 0},
			{-yS,   0,  yC, 0},
			{  0,   0,   0, 1},
		},
	}

	zC := math.Cos(r.Z)
	zS := math.Sin(r.Z)
	zMat := &Matrix4x4{
		entries: [4][4]float64{
			{ zC, -zS,   0, 0},
			{ zS,  zC,   0, 0},
			{  0,   0,   1, 0},
			{  0,   0,   0, 1},
		},
	}

	return xMat.Compose(yMat).Compose(zMat)
}

func (b *Matrix4x4) Compose(a *Matrix4x4) *Matrix4x4 {
	newMat := Matrix4x4{}
	for i:= 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			for k := 0; k < 4; k++ {
				newMat.entries[i][j] += b.Get(i, k) * a.Get(k, j)
			}
		}
	}
	return &newMat
}

func (mat Matrix4x4) ToString() string {
	return fmt.Sprintf(
		"[%v, %v, %v, %v]\n" +
		"[%v, %v, %v, %v]\n" +
		"[%v, %v, %v, %v]\n" +
		"[%v, %v, %v, %v]\n",
		mat.entries[0][0], mat.entries[0][1], mat.entries[0][2], mat.entries[0][3],
		mat.entries[1][0], mat.entries[1][1], mat.entries[1][2], mat.entries[1][3],
		mat.entries[2][0], mat.entries[2][1], mat.entries[2][2], mat.entries[2][3],
		mat.entries[3][0], mat.entries[3][1], mat.entries[3][2], mat.entries[3][3],
	)
}

func (mat *Matrix4x4) Multiply(vec *Vector3) *Vector3 {
	newVec := Vector3{}
	newVec.X = vec.X * mat.Get(0, 0) + vec.Y * mat.Get(1, 0) + vec.Z * mat.Get(2, 0) + mat.Get(3, 0)
	newVec.Y = vec.X * mat.Get(0, 1) + vec.Y * mat.Get(1, 1) + vec.Z * mat.Get(2, 1) + mat.Get(3, 1)
	newVec.Z = vec.X * mat.Get(0, 2) + vec.Y * mat.Get(1, 2) + vec.Z * mat.Get(2, 2) + mat.Get(3, 2)
	return &newVec
}

