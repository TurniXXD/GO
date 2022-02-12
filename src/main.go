package main

import (
	"github.com/TurniXXD/GO/api"
	// use variables to shorten exports names
	ct "github.com/TurniXXD/GO/convertTypes"
	dt "github.com/TurniXXD/GO/dataTypes"
	funcs "github.com/TurniXXD/GO/functions"
	"github.com/TurniXXD/GO/math"
)

// main function that go reads first
func main() {
	// to use a func from another file start with file's name and then choose exported func
	funcs.Functions()
	dt.DataTypes()
	ct.ConvertTypes()
	math.Math()
	api.HandleRequests()
}
