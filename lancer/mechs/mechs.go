package mechs

import(
	"frames"
	"dataload"
	"fmt"
	"strings"
)

func GenerateMech() Mech{
	var newMech Mech
	newMech.Frame = frames.GetRandomFrame()
	//newMech.Frame
	newMech.ColorScheme = "Flat grey of a newly printed mech"
	return newMech
}

func ToString(mecho Mech){
	//frame
		//manufacturer
		//color scheme
		//mounts
			//main
			//heavy
			//aux
		//core_system
	fmt.Println("ID: " +mecho.Frame.ID)
	fmt.Println("Name: " +mecho.Frame.Name)
	fmt.Println("Manufacturer: "+mecho.Frame.Source)
	fmt.Println("Color Scheme: "+mecho.ColorScheme)
	fmt.Println("Mounts: " + strings.Join(mecho.Frame.Mounts,", "))
	fmt.Println("Main Mount: ")
	fmt.Println("Heavy Mount: ")
	fmt.Println("Aux Mount: ")
	fmt.Println("Core System: ")
	//fmt.Println("Description: " mecho.Frame.Description)
}

type Mech struct{
	Frame dataload.Frame
	Weapons []dataload.Weapon
	ColorScheme string
}

