package mechs

import(
	"frames"
	"dataload"
	"weapons"
	"fmt"
	"strings"
	"mounts"
)

func GenerateMech() Mech{
	var newMech Mech
	newMech.Frame = frames.GetRandomFrame()
	newMech.Mounts = initMountsGroup(newMech)
	//newMech.Weapons = weapons.GenerateWeaponsList(newMech.Frame.Mounts)
	//newMech.Weapons, MountMap = weapons.FillMountsByWeaponType(newMech.Frame.Mounts)
	newMech.Weapons, newMech.Mounts = weapons.FillMountsByWeaponType(newMech.Mounts)
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
	fmt.Println("Weapons: "+weapons.JoinWeapons(mecho.Weapons, ","))
	for _, mount := range mecho.Mounts{
		fmt.Println(mount.Type + " Mount: " + mounts.JoinWeapons(mount,","))
	}
	fmt.Println("Core System: ")
	//fmt.Println("Description: " mecho.Frame.Description)
}

func initMountsGroup(mech Mech) []mounts.MountObject{
	var mountList []mounts.MountObject

	for _, mType := range mech.Frame.Mounts{
		mountList = append(mountList, mounts.NewMount(mType))
	}

	return mountList
}

type Mech struct{
	Frame dataload.Frame
	Weapons []dataload.Weapon
	MountMap map[string][]dataload.Weapon
	Mounts []mounts.MountObject
	ColorScheme string
}

