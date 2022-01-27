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
	return math.Sqrt(v.LengthSquare())
}

func (v *Vector3) LengthSquare() float64 {
	return v.x*v.x + v.y*v.y + v.z*v.z
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

func HitSphere(center *Point3, radius float64, r *Ray) float64 {
	oc := MinusVectors((*Vector3)(r.origin), (*Vector3)(center))
	a := r.direction.LengthSquare()
	halfB := Dot(oc, r.direction)
	c := oc.LengthSquare() - radius*radius
	discriminant := halfB*halfB - a*c
	if discriminant < 0 {
		return -1.0
	}

	return (-halfB - math.Sqrt(discriminant)) / a
}

func RayColor(r *Ray) *Color {
	t := HitSphere(&Point3{0, 0, -1}, 0.5, r)
	if t > 0.0 {
		n := MinusVectors((*Vector3)(r.At(t)), &Vector3{0, 0, -1})
		tmp := &Vector3{n.x + 1, n.y + 1, n.z + 1}
		tmp = tmp.Multiply(0.5)
		return (*Color)(tmp)
	}
	unitDirection := r.direction.UnitVector()
	t = 0.5 * (unitDirection.y + 1.0)
	color1 := &Color{1.0, 1.0, 1.0}
	color2 := &Color{0.5, 0.7, 1.0}

	a := MultiplyVectorByT((*Vector3)(color1), (1.0 - t))
	b := MultiplyVectorByT((*Vector3)(color2), t)
	result := AddVectors(a, b)
	return (*Color)(result)
}

func DivideVectorByT(v *Vector3, t float64) *Vector3 {
	return MultiplyVectorByT(v, 1/t)
}

func MultiplyVectorByT(v *Vector3, t float64) *Vector3 {
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
		result.z = result.z - p.z
	}

	return result
}

func AddVectors(v *Vector3, params ...*Vector3) *Vector3 {
	result := &Vector3{v.x, v.y, v.z}
	for _, p := range params {
		result.x = result.x + p.x
		result.y = result.y + p.y
		result.z = result.z + p.z
	}

	return result
}

func Dot(u, v *Vector3) float64 {
	return u.x*v.x + u.y*v.y + u.z*v.z
}

func Cross(u, v *Vector3) *Vector3 {
	return &Vector3{
		x: u.y*v.z - u.z*v.y,
		y: u.z*v.x - u.x*v.z,
		z: u.x*v.y - u.y*v.x,
	}
}

type HitRecord struct {
	p           *Point3
	normal      *Vector3
	t           float64
	isFrontFace bool
}

func (h *HitRecord) SetFaceNormal(r *Ray, outwardNormal *Vector3) {
	h.isFrontFace = Dot(r.direction, outwardNormal) < 0
	h.normal = outwardNormal
	if !h.isFrontFace {
		h.normal = MultiplyVectorByT(outwardNormal, -1)
	}
}

type Sphere struct {
	p, center *Point3
	normal    *Vector3
	t, radius float64
}

func (s *Sphere) Hit(r *Ray, tMin, tMax float64, rec *HitRecord) bool {
	oc := MinusVectors((*Vector3)(r.origin), (*Vector3)(s.center))
	a := r.direction.LengthSquare()
	halfB := Dot(oc, r.direction)
	c := oc.LengthSquare() - s.radius*s.radius

	discriminant := halfB*halfB - a*c
	if discriminant < 0 {
		return false
	}
	sqrtd := math.Sqrt(discriminant)

	root := (-halfB - sqrtd) / a
	if root < tMin || tMax < root {
		root = (-halfB + sqrtd) / a
		if root < tMin || tMax < root {
			return false
		}
	}

	rec.t = root
	rec.p = r.At(rec.t)
	outwardNormal := DivideVectorByT(MinusVectors((*Vector3)(rec.p), (*Vector3)(s.center)), s.radius)
	rec.SetFaceNormal(r, outwardNormal)
	rec.normal = DivideVectorByT(MinusVectors((*Vector3)(rec.p), (*Vector3)(s.center)), s.radius)
	return true
}

type HitableList struct {
	objects []*Sphere
}

func (h *HitableList) Hit(r *Ray, tMin, tMax float64, rec *HitRecord) bool {
	tmp := &HitRecord{}
	isHitAnything := false
	closestSoFar := tMax

	for _, item := range h.objects {
		if item.Hit(r, tMin, closestSoFar, tmp) {
			isHitAnything = true
			closestSoFar = tmp.t
			rec = tmp
		}
	}

	return isHitAnything
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

	fmt.Printf("P3\n%d %d\n255\n", imageWidth, imageHeight)

	for h := imageHeight - 1; h >= 0; h-- {
		for w := 0; w < imageWidth; w++ {
			u := float64(w) / float64(imageWidth-1)
			v := float64(h) / float64(imageHeight-1)
			tmp := AddVectors(lowerLeftCorner, MultiplyVectorByT(&horizontal, u), MultiplyVectorByT(&vertical, v))
			direction := MinusVectors(tmp, (*Vector3)(origin))
			ray := &Ray{
				origin:    origin,
				direction: direction,
			}
			color := RayColor(ray)
			fmt.Print(color.Write())
		}
	}
}
