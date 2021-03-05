package main

import "math"

// Vector2D struct represents the vector
type Vector2D struct {
	x float64
	y float64
}

// Add will add the points of two vectors
func (v1 Vector2D) Add(v2 Vector2D) Vector2D {
	return Vector2D{v1.x + v2.x, v1.y + v2.y}
}

// Substract will subtract the points of two vectors
func (v1 Vector2D) Substract(v2 Vector2D) Vector2D {
	return Vector2D{v1.x - v2.x, v1.y - v2.y}
}

// Multiply will multiply the points of two vectors
func (v1 Vector2D) Multiply(v2 Vector2D) Vector2D {
	return Vector2D{v1.x * v2.x, v1.y * v2.y}
}

// AddV will add a given value to the points of vector
func (v1 Vector2D) AddV(d float64) Vector2D {
	return Vector2D{v1.x + d, v1.y + d}
}

// MultiblyV will multiply a given value to the points of vector
func (v1 Vector2D) MultiblyV(d float64) Vector2D {
	return Vector2D{v1.x * d, v1.y * d}
}

// DivisionV will subtract a given value to the points of vector
func (v1 Vector2D) DivisionV(d float64) Vector2D {
	return Vector2D{v1.x / d, v1.y / d}
}

// limit will limit the vector values to the given parameters
func (v1 Vector2D) limit(lower, upper float64) Vector2D {
	return Vector2D{math.Min(math.Max(v1.x, lower), upper), math.Min(math.Max(v1.y, lower), upper)}
}

// Distance func will take a vector and tell us how far we
// are from that vector
//						#(c, d)
//					#	.
//				#		.
//			#			. sqr((a - c)^2 + (b - d)^2)
//		#				.
//	#(a, b)				.
func (v1 Vector2D) Distance(v2 Vector2D) float64 {
	return math.Sqrt(math.Pow(v1.x-v2.x, 2) + math.Pow(v1.y-v2.y, 2))
}
