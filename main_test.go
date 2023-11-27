package geoquery

import (
	"fmt"
	"testing"

	"github.com/adam-ghafara/geoquery/gq"
	"github.com/adam-ghafara/geoquery/helper"
)

func TestUpdateGetData(t *testing.T) {
	mconn := helper.SetConnection("mongodb+srv://admin:402390@kukidata.jtgvziw.mongodb.net/", "geojsonList")
	datagedung := gq.GeoIntersects(mconn, 107.60010652297046, -6.925622320467241)
	fmt.Println(datagedung)
}
