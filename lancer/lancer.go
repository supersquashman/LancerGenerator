package main

import (
	"mechs"
	"dataload"
)

func main(){
	dataload.Load_All_Frames()
	
	mech := mechs.GenerateMech()
	mechs.ToString(mech)

}