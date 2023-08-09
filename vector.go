/*
# Vector

This is a simple package for working with 3D vectors
*/
package vector

/*
Represents 3D vector data type
*/
type vector64 struct {
	X float64
	Y float64
	Z float64
}

/*
Initializes 3D vector
*/
func Vector(x float64, y float64, z float64) vector64 {
	return vector64{
		X: x,
		Y: y,
		Z: z,
	}
}
