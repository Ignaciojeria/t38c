package e2e

import (
	"context"
	"testing"

	"github.com/Ignaciojeria/t38c"
	"github.com/paulmach/orb"
	"github.com/paulmach/orb/geojson"
	"github.com/stretchr/testify/require"
)

func testWithin(t *testing.T, client *t38c.Client) {
	err := client.Keys.Set("points", "point-1").Point(1, 1).Do(context.Background())
	require.NoError(t, err)

	outerRing := orb.Ring{
		orb.Point{0, 0},
		orb.Point{20, 0},
		orb.Point{20, 20},
		orb.Point{0, 20},
		orb.Point{0, 0}, // Cerrando el anillo
	}

	polygon := orb.Polygon{outerRing}

	geom := geojson.NewGeometry(polygon)

	err = client.Keys.Set("areas", "area-1").Geometry(geom).Do(context.Background())
	require.NoError(t, err)

	resp, err := client.Search.Within("points").
		Get("areas", "area-1").
		Format(t38c.FormatIDs).Do(context.Background())
	require.NoError(t, err)

	require.Equal(t, []string{"point-1"}, resp.IDs)
}
