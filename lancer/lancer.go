package main

import (
	"github.com/supersquashman/LancerGenerator/dataload"
	"github.com/supersquashman/LancerGenerator/mechs"
)

func main() {
	dataload.Load_All()

	mech := mechs.GenerateMech()
	mechs.ToString(mech)

}
