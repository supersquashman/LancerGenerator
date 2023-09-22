package main

import (
	"mechs"
	"dataload"
)

func main(){
	mech := mechs.GenerateMech()
	mechs.ToString(mech)

	dataload.Load_All_Frames()
}