package main

import (
	"log"

	"github.com/davecgh/go-spew/spew"
	"github.com/garyburd/redigo/redis"
)

func main() {
	c, err := redis.Dial("tcp", ":32768")
	if err != nil {
		log.Fatalf("Could not connect: %v\n", err)
	}
	defer c.Close()

	paris := `
	{
		"type": "Feature",
		"properties": {},
		"geometry": {
			"type": "Polygon",
			"coordinates": [
				[
					[
						2.3548507690429683,
						48.90647978628254
					],
					[
						2.29888916015625,
						48.90332040974438
					],
					[
						2.252197265625,
						48.87465122643438
					],
					[
						2.230224609375,
						48.847773057644694
					],
					[
						2.2899627685546875,
						48.81296811706886
					],
					[
						2.3349380493164062,
						48.807089572708925
					],
					[
						2.4097824096679688,
						48.81115940759497
					],
					[
						2.4334716796875,
						48.83353759505566
					],
					[
						2.4461746215820312,
						48.85974578950385
					],
					[
						2.4365615844726562,
						48.890455171696374
					],
					[
						2.3850631713867183,
						48.90625412315376
					],
					[
						2.3548507690429683,
						48.90647978628254
					]
				]
			]
		}
	}`

	pointInside := `
	{
		"type": "Point",
		"coordinates": [
		  2.3531341552734375,
		  48.88052184622007
		]
	}`

	pointOutside := `
	{
		"type": "Point",
		"coordinates": [
			2.3622536659240723,
			48.8047365487258
		]
	}`

	res, _ := c.Do("SET", "cities", "paris", "OBJECT", paris)
	spew.Dump(res)

	res2, _ := c.Do("GET", "cities", "paris")
	resStr, err := redis.String(res2, nil)
	spew.Dump(resStr, err)

	res3, _ := c.Do("INTERSECTS", "cities", "OBJECT", pointInside)
	blek, err := redis.StringMap(res3, nil)
	spew.Dump(blek, err)
	spew.Dump(res3)

	res4, _ := c.Do("INTERSECTS", "cities", "OBJECT", pointOutside)
	spew.Dump(res4)
	spew.Dump("SHIRO")
}
