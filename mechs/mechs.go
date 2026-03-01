package mechs

import(
	"frames"
	"dataload"
	"weapons"
	"fmt"
	// "strings"
	"mounts"
	"colorscheme"
)

func GenerateMech() dataload.Mech{
	var newMech dataload.Mech
	newMech.Frame = frames.GetRandomFrame()
	// newMech.Frame.Mounts = initMountsGroup(newMech.Frame)
	newMech.Frame.Mounts = initMountsGroup(newMech)
	//newMech.Weapons = weapons.GenerateWeaponsList(newMech.Frame.Mounts)
	//newMech.Weapons, MountMap = weapons.FillMountsByWeaponType(newMech.Frame.Mounts)
	newMech.Weapons, newMech.Frame.Mounts = weapons.FillMountsByWeaponType(newMech.Frame.Mounts)
	// newMech.ColorScheme = "Flat grey of a newly printed mech"
	newMech.ColorScheme = colorscheme.GetRandomColorscheme()
	return newMech
}

func ToString(mecho dataload.Mech) string{
	frameString := ""
	//frame
		//manufacturer
		//color scheme
		//mounts
			//main
			//heavy
			//aux
		//core_system
	frameString = fmt.Sprintf("%s%s\n",frameString,"id: " +mecho.Frame.ID)
	frameString = fmt.Sprintf("%s%s\n",frameString,"name: " +mecho.Frame.Name)
	frameString = fmt.Sprintf("%s%s\n",frameString,"manufacturer: "+mecho.Frame.Source)
	frameString = fmt.Sprintf("%s%s\n",frameString,"color scheme: "+mecho.ColorScheme)
	// frameString = fmt.Sprintf("%s%s\n",frameString,"mounts: " + strings.Join(mecho.Frame.Mounts,", "))
	frameString = fmt.Sprintf("%s%s\n",frameString,"weapons: "+weapons.JoinWeapons(mecho.Weapons, ","))
	for _, mount := range mecho.Frame.Mounts{
		frameString = fmt.Sprintf("%s%s\n",frameString, mount.Type + " Mount: " + mounts.JoinWeapons(mount,","))
	}
	frameString = fmt.Sprintf("%s%s\n",frameString,"core system: ")
	//fmt.Sprintf("Description: " mecho.Frame.Description)
	return frameString
}

func PrintMech(mecho dataload.Mech){
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
	// fmt.Println("Mounts: " + strings.Join(mecho.Frame.Mounts,", "))
	fmt.Println("Weapons: "+weapons.JoinWeapons(mecho.Weapons, ","))
	for _, mount := range mecho.Frame.Mounts{
		fmt.Println(mount.Type + " Mount: " + mounts.JoinWeapons(mount,","))
	}
	fmt.Println("Core System: ")
	//fmt.Println("Description: " mecho.Frame.Description)
}

// func initMountsGroup(frame dataload.Frame) []mounts.MountObject{
// 	var mountList []mounts.MountObject

// 	for _, mType := range frame.Mounts{
// 		mountList = append(mountList, mounts.NewMount(mType))
// 	}

// 	return mountList
// }

func initMountsGroup(mech dataload.Mech) []dataload.Mount{
	var mountList []dataload.Mount

	for _, mType := range mech.Frame.MountList{
		mountList = append(mountList, mounts.NewMount(mType))
	}

	return mountList
}

type Mech struct{
	Frame dataload.Frame `json:"frame"`
	Weapons []dataload.Weapon `json:"weapons"`
	MountMap map[string][]dataload.Weapon `json:"mount_map"`
	Mounts []dataload.Mount `json:"mounts"`
	ColorScheme string `json:"color_scheme"`
}

