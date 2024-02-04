package util

// direction is an enum for the cardinal directions.
type direction int

// Constants for the cardinal directions.
const (
	Up direction = iota
	Down
	Left
	Right
)

// String returns the string representation of the direction.
func (d direction) String() string {
	return [...]string{"up", "down", "left", "right"}[d]
}

// DirectionFromString returns the direction from a string.
func DirectionFromString(s string) direction {
	switch s {
	case "up":
		return Up
	case "down":
		return Down
	case "left":
		return Left
	case "right":
		return Right
	}
	return -1
}

// DirectionFromInt returns the direction from an int.
func DirectionFromInt(i int) direction {
	switch i {
	case 0:
		return Up
	case 1:
		return Down
	case 2:
		return Left
	case 3:
		return Right
	}
	return -1
}

// Int returns the int representation of the direction.
func (d direction) Int() int {
	return int(d)
}

// Opposite returns the opposite direction.
func (d direction) Opposite() direction {
	switch d {
	case Up:
		return Down
	case Down:
		return Up
	case Left:
		return Right
	case Right:
		return Left
	}
	return -1
}

// Rotate returns the direction rotated by 90 degrees clockwise.
func (d direction) Rotate() direction {
	switch d {
	case Up:
		return Right
	case Down:
		return Left
	case Left:
		return Up
	case Right:
		return Down
	}
	return -1
}

// RotateCounter returns the direction rotated by 90 degrees counter-clockwise.
func (d direction) RotateCounter() direction {
	switch d {
	case Up:
		return Left
	case Down:
		return Right
	case Left:
		return Down
	case Right:
		return Up
	}
	return -1
}
