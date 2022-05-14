package game

type Vector struct {
	X int32
	Y int32
}

func (v *Vector) Add(other *Vector) {
	v.X += other.X
	v.Y += other.Y
}
