package gcf

import (
	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
	"github.com/adam-ghafara/geoquery"
)

func init() {
	functions.HTTP("Init", geoquery.PostGeoIntersects)
}
