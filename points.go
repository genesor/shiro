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
		Latitude:  48.822194345985714,
		Longitude: 2.363004684448242,
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

var pointParcDesPrinces = &Point{
	Coordinates: Coordinates{
		Latitude:  48.841404318083505,
		Longitude: 2.253119945526123,
	},
}

var pointPorteStCloud = &Point{
	Coordinates: Coordinates{
		Latitude:  48.83788774899389,
		Longitude: 2.2564244270324703,
	},
}

var ptdonut = &Point{
	Coordinates: Coordinates{
		Longitude: 2.234344482421875,
		Latitude:  49.13006052488493,
	},
}

var ptdonuthole = &Point{
	Coordinates: Coordinates{
		Latitude:  49.139045760119835,
		Longitude: 2.036590576171875,
	},
}
