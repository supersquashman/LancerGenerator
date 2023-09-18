package mechs

import(
	"frames"
	"fmt"
)

func GenerateMech() Mech{
	var newMech Mech
	newMech.frame = frames.GetFrame()
	newMech.colorScheme = "Flat grey of a newly printed mech"
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
	fmt.Println("Manufacturer: "+mecho.frame.Manufacturer())
	fmt.Println("Color Scheme: "+mecho.colorScheme)
	fmt.Println("Mounts: ")
	fmt.Println("Main Mount: ")
	fmt.Println("Heavy Mount: ")
	fmt.Println("Aux Mount: ")
	fmt.Println("Core System: ")
}

type Mech struct{
	frame frames.Frame
	weapons []string
	colorScheme string
}

