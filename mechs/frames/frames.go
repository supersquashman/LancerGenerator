package frames

import(
	"math/rand"
	"dataload"
	"mounts"
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

func InstantiateFrame(frameArch dataload.FrameArchetype) dataload.Frame{
	tempFrame := dataload.Frame {
		ID: frameArch.ID,
		LicenseLevel: frameArch.LicenseLevel,
		LicenseID: frameArch.LicenseID,
		Variant: frameArch.Variant,
		Source: frameArch.Source,
		Name: frameArch.Name,
		MechType: frameArch.MechType, // []string `json:"mechtype"`
		MountList: frameArch.Mounts,
		Description: frameArch.Description,
		Mounts: mounts.InitMountList(frameArch.Mounts), // []Mount //`json:"mounts"`
		Stats: frameArch.Stats,
		Traits: frameArch.Traits, // []FrameTraits `json:"traits"`
		CoreSystem: frameArch.CoreSystem,
		ImageURL: frameArch.ImageURL,
		YPOS: frameArch.YPOS,
	}

	return tempFrame
}

func GetRandomFrame() dataload.Frame{
	//var tempFrame Frame
	tempFrameArch := dataload.AllFramesList[rand.Intn(len(dataload.AllFramesList))]
	tempFrame := InstantiateFrame(tempFrameArch)
	return tempFrame
}

func GetRandomFrameManufacturer() string{
	tempFrame := dataload.AllFramesList[rand.Intn(len(dataload.AllFramesList))]
	manufacturer := tempFrame.Source
	return manufacturer
}