package main

import (
	"fmt"
	"math"
)

type Vector3 struct {
	x, y, z float64
}

func (v *Vector3) Add(input *Vector3) *Vector3 {
	return &Vector3{
		x: v.x + input.x,
		y: v.y + input.y,
		z: v.z + input.z,
	}
}

func (v *Vector3) Multiply(t float64) *Vector3 {
	return &Vector3{
		x: v.x * t,
		y: v.y * t,
		z: v.z * t,
	}
}

func (v *Vector3) Divide(t float64) *Vector3 {
	return &Vector3{
		x: v.x / t,
		y: v.y / t,
		z: v.z / t,
	}
}

func (v *Vector3) Length() float64 {
	return math.Sqrt(v.x*v.x + v.y*v.y + v.z*v.z)
}

func (v *Vector3) UnitVector() *Vector3 {
	return v.Divide(v.Length())
}

func (v *Vector3) ToString() string {
	return fmt.Sprintf("(%f, %f, %f)", v.x, v.y, v.z)
}

type Point3 Vector3
type Color Vector3

func (c *Color) Write() string {
	r := c.x * 255.999
	g := c.y * 255.999
	b := c.z * 255.999
	return fmt.Sprintf("%d %d %d\n", int(r), int(g), int(b))
}

type Ray struct {
	origin    *Point3
	direction *Vector3
}

func (r *Ray) At(t float64) *Point3 {
	tmp := r.direction.Multiply(t)
	return &Point3{
		r.origin.x + tmp.x,
		r.origin.y + tmp.y,
		r.origin.z + tmp.z,
	}
}

func RayColor(r *Ray) *Color {
	unitVector := r.direction.UnitVector()
	// fmt.Print(unitVector.ToString())
	t := 0.5 * (unitVector.y + 1.0)

	color1 := &Color{1.0, 1.0, 1.0}
	color1.x = color1.x * (1.0 - t)
	color1.y = color1.y * (1.0 - t)
	color1.z = color1.z * (1.0 - t)
	color2 := &Color{0.5, 0., 1.0}
	color2.x = color2.x * t
	color2.y = color2.y * t
	color2.z = color2.z * t

	result := &Color{
		x: color1.x - color2.x,
		y: color1.y - color2.y,
		z: color1.z - color2.z,
	}
	return result
}

func DivideVectorByT(v *Vector3, t float64) *Vector3 {
	return &Vector3{
		x: v.x / t,
		y: v.y / t,
		z: v.z / t,
	}
}

func MutilplyVectorByT(v *Vector3, t float64) *Vector3 {
	return &Vector3{
		x: v.x * t,
		y: v.y * t,
		z: v.z * t,
	}
}

func MinusVectors(v *Vector3, params ...*Vector3) *Vector3 {
	result := &Vector3{v.x, v.y, v.z}
	for _, p := range params {
		result.x = result.x - p.x
		result.y = result.y - p.y
		result.z = result.y - p.z
	}

	return result
}

func AddVectors(v *Vector3, params ...*Vector3) *Vector3 {
	result := &Vector3{v.x, v.y, v.z}
	for _, p := range params {
		result.x = result.x + p.x
		result.y = result.y + p.y
		result.z = result.y + p.z
	}

	return result
}

func main() {

	// Image
	aspectRatio := 16.0 / 9.0
	imageWidth := 400
	imageHeight := int(float64(imageWidth) / aspectRatio)

	// Camera
	viewportHeight := 2.0
	viewportWidth := aspectRatio * viewportHeight
	focalLength := 1.0

	origin := &Point3{0, 0, 0}
	horizontal := Vector3{viewportWidth, 0, 0}
	vertical := Vector3{0, viewportHeight, 0}
	lowerLeftCorner := MinusVectors((*Vector3)(origin), DivideVectorByT(&horizontal, 2), DivideVectorByT(&vertical, 2), &Vector3{0, 0, focalLength})

	fmt.Printf("P3\n%d %d\n255\n", imageHeight, imageWidth)

	for h := imageHeight - 1; h >= 0; h-- {
		for w := 0; w < imageWidth; w++ {
			u := float64(w) / float64(imageWidth-1)
			v := float64(h) / float64(imageHeight-1)
			tmp := AddVectors(lowerLeftCorner, MutilplyVectorByT(&horizontal, u), MutilplyVectorByT(&vertical, v))
			ray := &Ray{
				origin:    origin,
				direction: MinusVectors(tmp, (*Vector3)(origin)),
			}
			color := RayColor(ray)
			fmt.Print(color.Write())
		}
	}
}
