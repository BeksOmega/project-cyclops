package main

type Cube struct {
	Center Vector3
	Scale Vector3
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
	halfX := cube.Scale.X / 2
	halfY := cube.Scale.Y / 2
	halfZ := cube.Scale.Z / 2
	center := cube.Center
	return [8]Vector3{
		{center.X + halfX, center.Y - halfY, center.Z - halfZ},
		{center.X + halfX, center.Y - halfY, center.Z + halfZ},
		{center.X + halfX, center.Y + halfY, center.Z - halfZ},
		{center.X + halfX, center.Y + halfY, center.Z + halfZ},
		{center.X - halfX, center.Y - halfY, center.Z - halfZ},
		{center.X - halfX, center.Y - halfY, center.Z + halfZ},
		{center.X - halfX, center.Y + halfY, center.Z - halfZ},
		{center.X - halfX, center.Y + halfY, center.Z + halfZ},
	}
}
