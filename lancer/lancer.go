package main

import (
	"mechs"
	"dataload"
)

func main(){
	dataload.Load_All()
	
	mech := mechs.GenerateMech()
	mechs.ToString(mech)

}