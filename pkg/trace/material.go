package trace

import "github.com/hunterloftis/oneweekend/pkg/geom"

// Material represents a material that scatters light.
type Material interface {
	Scatter(in geom.Unit, n geom.Unit, p geom.Vec) (out geom.Unit, attenuation Color, ok bool)
}