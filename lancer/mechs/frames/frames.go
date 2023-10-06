package frames

import(
	"math/rand"
	"dataload"
)

type Frame dataload.Frame

func GetFrame() string{
	/*var tempFrame Frame
	frameList := []string{"GMS","IPS-Northstar","Smith-Shimano Corpo","Horus","Harrison Armory"}
	tempFrame.manufacturer = frameList[rand.Intn(len(frameList))]
	//tempFrame.mounts = []string{"Flex","Main","Heavy"}
	tempFrame.core_system = "Power Up"*/
	return ""
}

func (frame Frame) Manufacturer() string{
	return ""
}


func GenerateRandomFrame(){

}

func GetRandomFrame() dataload.Frame{
	//var tempFrame Frame
	tempFrame := dataload.AllFramesList[rand.Intn(len(dataload.AllFramesList))]
	return tempFrame
}

func GetRandomFrameManufacturer() string{
	tempFrame := dataload.AllFramesList[rand.Intn(len(dataload.AllFramesList))]
	manufacturer := tempFrame.Source
	return manufacturer
}