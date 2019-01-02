package trace

import "github.com/hunterloftis/oneweekend/pkg/geom"

// Metal describes a reflective material
type Metal struct {
	Albedo Color
	Rough  float64
}

// NewMetal creates a new Metal material with a given color and roughness.
func NewMetal(albedo Color, roughness float64) Metal {
	return Metal{Albedo: albedo, Rough: roughness}
}

// Scatter reflects incoming light rays about the normal.
func (m Metal) Scatter(in geom.Ray, p geom.Vec, n geom.Unit) (out geom.Ray, attenuation Color, ok bool) {
	r := in.Dir.Reflect(n)
	dir := r.Plus(geom.RandVecInSphere().Scaled(m.Rough)).ToUnit()
	out = geom.NewRay(p, dir)
	return out, m.Albedo, out.Dir.Dot(n) > 0
}