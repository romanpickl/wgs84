package wgs84

import "math"

func spheroid(gs ...GeodeticSpheroid) Spheroid {
	for _, s := range gs {
		if s == nil {
			continue
		}
		if s, ok := s.(Spheroid); ok {
			return s
		}
		return Spheroid{
			A:  s.MajorAxis(),
			Fi: s.InverseFlattening(),
		}
	}
	return Spheroid{}
}

func toWGS84(t Transformation, x, y, z float64) (x0, y0, z0 float64) {
	if t == nil {
		return x, y, z
	}
	return t.ToWGS84(x, y, z)
}

func fromWGS84(t Transformation, x0, y0, z0 float64) (x, y, z float64) {
	if t == nil {
		return x0, y0, z0
	}
	return t.FromWGS84(x0, y0, z0)
}

func sin2(east float64) float64 {
	return math.Pow(math.Sin(east), 2)
}

func cos2(east float64) float64 {
	return math.Pow(math.Cos(east), 2)
}

func tan2(east float64) float64 {
	return math.Pow(math.Tan(east), 2)
}

func degree(r float64) float64 {
	return r * 180 / math.Pi
}

func radian(d float64) float64 {
	return d * math.Pi / 180
}
