package main

import (
	"log"

	"errors"

	"github.com/davecgh/go-spew/spew"
	"github.com/garyburd/redigo/redis"
)

func main() {
	c, err := redis.Dial("tcp", ":32768")
	if err != nil {
		log.Fatalf("Could not connect: %v\n", err)
	}
	defer c.Close()

	// CURL
	// INTERSECTS cities OBJECT {"type":"Point","coordinates":[2.3531341552734375,48.88052184622007]}
	// INTERSECTS cities OBJECT {"type":"Point","coordinates":[2.3622536659240723,48.8047365487258]}

	_, err = c.Do("SET", "cities", "paris", "OBJECT", paris)
	if err != nil {
		log.Fatalf("Could not SET in store: %v\n", err)
	}

	_, err = c.Do("SET", "cities", "vj-vitry", "OBJECT", VJIvry)
	if err != nil {
		log.Fatalf("Could not SET in store: %v\n", err)
	}

	m, err := findEncompassingZones(c, pointNorthParis)
	// Should display 1 match and no err
	spew.Dump(m, err)

	m2, err := findEncompassingZones(c, pointVillejuif)
	// Should display 1 match and no err
	spew.Dump(m2, err)

	m3, err := findEncompassingZones(c, pointSouthParis)
	// Should display 2 match and no err
	spew.Dump(m3, err)

	m4, err := findEncompassingZones(c, pointAulnaySous)
	// Should display 0 match and no err
	spew.Dump(m4, err)

	//geoJSON := new(GeoJSONPolygonFeature)
	//err = json.Unmarshal(, geoJSON)
	//spew.Dump(geoJSON, err)

	spew.Dump("SHIRO")
}

func findEncompassingZones(c redis.Conn, p *Point) ([]*MatchResult, error) {
	res, err := c.Do("INTERSECTS", "cities", "OBJECT", p.ToGeoJSON())
	if err != nil {
		return nil, err
	}

	resSlice, ok := res.([]interface{})
	if !ok {
		return nil, errors.New("unexpected tile38 response format 1")
	}
	resSlice, ok = resSlice[1].([]interface{})
	if !ok {
		return nil, errors.New("unexpected tile38 response format 2")
	}

	matches := make([]*MatchResult, len(resSlice))
	for i, r := range resSlice {
		match, ok := r.([]interface{})
		if !ok {
			return nil, errors.New("unexpected tile38 response format 2")
		}

		matches[i] = &MatchResult{
			Name:    string(match[0].([]byte)),
			GeoJSON: match[1].([]byte),
		}
	}

	return matches, nil
}
