package main

type GeoJSONPolygonFeature struct {
	Type     string         `json:"type"`
	Geometry GeoJSONPolygon `json:"geometry"`
}

type GeoJSONPolygon struct {
	Type        string        `json:"type"`
	Coordinates [][][]float64 `json:"coordinates"`
}

type MatchResult struct {
	Name    string
	GeoJSON []byte
}

var paris = `{
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

var VJIvry = `{
	"type": "Feature",
	"properties": {},
	"geometry": {
		"type": "Polygon",
		"coordinates": [
			[
				[
					2.4104690551757812,
					48.79827046389008
				],
				[
					2.407379150390625,
					48.80935063287144
				],
				[
					2.39501953125,
					48.82020230247925
				],
				[
					2.3785400390625,
					48.82359296752074
				],
				[
					2.3610305786132812,
					48.82359296752074
				],
				[
					2.3476409912109375,
					48.81794173168324
				],
				[
					2.3418045043945312,
					48.807541792897744
				],
				[
					2.3352813720703125,
					48.79578274255762
				],
				[
					2.3380279541015625,
					48.7794964290288
				],
				[
					2.349014282226562,
					48.76795709209035
				],
				[
					2.3692703247070312,
					48.762978520066326
				],
				[
					2.40325927734375,
					48.77044619302067
				],
				[
					2.4152755737304683,
					48.78085382371311
				],
				[
					2.4104690551757812,
					48.79827046389008
				]
			]
		]
	}
}`