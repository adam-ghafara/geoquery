package config

import "github.com/adam-ghafara/geoquery/helper"

var Mongocon = helper.SetConnection(Mongostring, "geojson")
