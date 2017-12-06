package main

import (
	"fmt"
)

type Coordinates struct {
	Latitude  float64
	Longitude float64
}

type Point struct {
	Coordinates Coordinates
}

func (p *Point) ToGeoJSON() []byte {
	return []byte(`{"type":"Point","coordinates":[` + fmt.Sprint(p.Coordinates.Longitude) + `,` + fmt.Sprint(p.Coordinates.Latitude) + `]}`)
}

var pointNorthParis = &Point{
	Coordinates: Coordinates{
		Latitude:  48.88052184622007,
		Longitude: 2.3531341552734375,
	},
}

var pointSouthParis = &Point{
	Coordinates: Coordinates{
		Latitude:  48.8154549860561,
		Longitude: 2.379226684570312,
	},
}

var pointVillejuif = &Point{
	Coordinates: Coordinates{
		Latitude:  48.8047365487258,
		Longitude: 2.3622536659240723,
	},
}

var pointAulnaySous = &Point{
	Coordinates: Coordinates{
		Latitude:  48.91821286473131,
		Longitude: 2.4708938598632812,
	},
}
