package trace

import (
	"math"

	"github.com/hunterloftis/oneweekend/pkg/geom"
)

// Sphere represents a spherical Surface
type Sphere struct {
	Center0, Center1 geom.Vec
	T0, T1           float64
	Rad              float64
	Mat              Material
}

// NewSphere creates a new Sphere with the given center and radius.
func NewSphere(center geom.Vec, radius float64, m Material) *Sphere {
	return NewMovingSphere(center, center, 0, 1, radius, m)
}

// NewMovingSphere creates a new Sphere with two centers separated by times t0 and t1
func NewMovingSphere(center0, center1 geom.Vec, t0, t1, radius float64, m Material) *Sphere {
	return &Sphere{
		Center0: center0,
		Center1: center1,
		T0:      t0,
		T1:      t1,
		Rad:     radius,
		Mat:     m,
	}
}

// Hit finds the distance to the first intersection (if any) between Ray r and the Sphere's surface.
// If no intersection is found, d = 0.
func (s *Sphere) Hit(r Ray, dMin, dMax float64) (d float64, bo Bouncer) {
	oc := r.Or.Minus(s.Center(r.t))
	a := r.Dir.Dot(r.Dir)
	b := oc.Dot(r.Dir.Vec)
	c := oc.Dot(oc) - s.Rad*s.Rad
	disc := b*b - a*c
	if disc <= 0 {
		return 0, s
	}
	sqrt := math.Sqrt(b*b - a*c)
	d = (-b - sqrt) / a
	if d > dMin && d < dMax {
		return d, s
	}
	d = (-b + sqrt) / a
	if d > dMin && d < dMax {
		return d, s
	}
	return 0, s
}

// Bounce returns the normal and material at point p on the Sphere
func (s *Sphere) Bounce(in Ray, dist float64) (out Ray, attenuation Color, ok bool) {
	p := in.At(dist)
	norm := p.Minus(s.Center(in.t)).Scaled(s.Rad).Unit()
	dir, attenuation, ok := s.Mat.Scatter(in.Dir, norm)
	return NewRay(p, dir, in.t), attenuation, ok
}

// Center returns the center of the sphere at a given time t
func (s *Sphere) Center(t float64) geom.Vec {
	p := (t - s.T0) / (s.T1 - s.T0)
	offset := s.Center1.Minus(s.Center0).Scaled(p)
	return s.Center0.Plus(offset)
}

// Box returns the Axis Aligned Bounding Box of the sphere encompassing all times between t0 and t1
func (s *Sphere) Box(t0, t1 float64) (box *AABB) {
	box0 := NewAABB(
		s.Center(t0).Minus(geom.NewVec(s.Rad, s.Rad, s.Rad)),
		s.Center(t0).Plus(geom.NewVec(s.Rad, s.Rad, s.Rad)),
	)
	box1 := NewAABB(
		s.Center(t1).Minus(geom.NewVec(s.Rad, s.Rad, s.Rad)),
		s.Center(t1).Plus(geom.NewVec(s.Rad, s.Rad, s.Rad)),
	)
	return box0.Plus(box1)
}
