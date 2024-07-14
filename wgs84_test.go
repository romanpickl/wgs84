package wgs84

import (
	"testing"
)

func TestTransform(t *testing.T) {
	tests := []struct {
		name     string
		fromEpsg int
		inA      float64
		inB      float64
		inC      float64
		toEpsg   int
		wantA    float64
		wantB    float64
		wantC    float64
	}{
		{"Web Mercator: EPSG(4326) to EPSG(3857)", 4326, 10, 50, 0, 3857, 1.113194908e+06, 6.446275841e+06, 0},
		{"OSGB: EPSG(4326) to EPSG(27700)", 4326, -2.25, 52.25, 0, 27700, 383029.296, 261341.615, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fromEPSG := EPSG(tt.fromEpsg)
			toEPSG := EPSG(tt.toEpsg)

			transform := Transform(fromEPSG, toEPSG).Round(3)

			gotA, gotB, gotC := transform(tt.inA, tt.inB, tt.inC)

			if gotA != tt.wantA {
				t.Errorf("Transform() A = %v, want %v", gotA, tt.wantA)
			}
			if gotB != tt.wantB {
				t.Errorf("Transform() B = %v, want %v", gotB, tt.wantB)
			}
			if gotC != tt.wantC {
				t.Errorf("Transform() C = %v, want %v", gotC, tt.wantC)
			}
		})
	}
}
