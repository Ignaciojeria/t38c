//go:build ignore
// +build ignore

package main

import (
	"context"
	"log"

	"github.com/Ignaciojeria/t38c"
	"github.com/paulmach/orb"
	"github.com/paulmach/orb/geojson"
)

func main() {
	tile38, err := t38c.New(t38c.Config{
		Address: "localhost:9851",
		Debug:   true,
	})
	if err != nil {
		log.Fatal(err)
	}
	defer tile38.Close()

	tile38.Keys.Set("fleet", "truck1").Point(33.5123, -112.2693).Do(context.TODO())
	tile38.Keys.Set("fleet", "truck1").PointZ(33.5123, -112.2693, 225).Do(context.TODO())
	tile38.Keys.Set("fleet", "truck1").Bounds(30, -110, 40, -100).Do(context.TODO())
	tile38.Keys.Set("fleet", "truck1").Hash("9tbnthxzr").Do(context.TODO())

	ring := orb.Ring{
		{0, 0},
		{10, 10},
		{10, 0},
		{0, 0},
	}
	polygon := orb.Polygon{ring}

	geometry := geojson.NewGeometry(polygon)
	tile38.Keys.Set("city", "tempe").Geometry(geometry).Do(context.TODO())
}
