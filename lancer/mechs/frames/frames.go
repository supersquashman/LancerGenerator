package frames

import(
	"math/rand"
)

func GetFrame() Frame{
	var tempFrame Frame
	frameList := []string{"GMS","IPS-Northstar","Smith-Shimano Corpo","Horus","Harrison Armory"}
	tempFrame.manufacturer = frameList[rand.Intn(len(frameList))]
	//tempFrame.mounts = []string{"Flex","Main","Heavy"}
	tempFrame.core_system = "Power Up"
	return tempFrame
}

func (frame Frame) Manufacturer() string{
	return frame.manufacturer
}

type Frame struct{
	manufacturer string
	mounts []Mount
	core_system string
}

type Mount struct{
	mountType string
}