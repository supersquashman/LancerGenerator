package weapons

import(
	"dataload"
	"math/rand"
	"strings"
	"fmt"
	"mounts"
)

const dMain = "main"
const dHeavy = "heavy"
const dFlex = "flex"
const dAux = "aux"
const dAuxiliary = "auxiliary"
const dAuxAux = "aux/aux"
const dMainAux = "main/aux"
const dIntegrated = "integrated"
const dSuperheavy = "superheavy"

func GenerateWeaponsList(mountList []string) []dataload.Weapon{
	
	tempWeaponsList := make([]dataload.Weapon, 0, len(mountList))
	for _, mount := range mountList{
		tempWeapon := SelectRandomWeaponOfType(mount)
		
		tempWeaponsList = append(tempWeaponsList, tempWeapon)

	}
	return tempWeaponsList
}

//function to select a random weapon
func SelectRandomWeapon(){

}

//function to select random weapon of a specific type
func SelectRandomWeaponOfType(wepType string) dataload.Weapon{
	var tempWeapon dataload.Weapon
	for {
		tempWeapon = dataload.AllWeaponsList[rand.Intn(len(dataload.AllWeaponsList))]
		fmt.Println("WepType: " + wepType)
		fmt.Println("Weapon Type: " + tempWeapon.Mount)
		if (tempWeapon.Mount == wepType){
			break
		}
	}
	return tempWeapon
}

//function to fill specific mount type
func FillMountByType(mountType string) []dataload.Weapon{
	var tempWeapons []dataload.Weapon
	
	switch mountType{
		case dMain: //one main or aux

		case dHeavy: //one heavy, main, or aux

		case dAux: //one aux, probably not a real case, but keeping for sake of homebrews

		case dAuxAux: //up to two aux

		case dMainAux: //one main and one aux, or two aux

		case dFlex: //one main or up to two aux

		case dIntegrated: //special, including for now, may remove

		case dSuperheavy: //no mount type; weapons of this type require heavymount and a mount of any other size
	}
	return tempWeapons
}

//function to fill mounts based on a weapons-first approach
func FillMountsByWeaponType(mountList []mounts.MountObject) ([]dataload.Weapon, []mounts.MountObject){
	var tempWeapons []dataload.Weapon
	tempMountList := mountList
	
	attached := false
	
	i := 0

	for{
		i = i+1
		emptySpace := false
		
		for _, mount := range mountList{
			if (!mounts.IsFull(mount)){
				emptySpace = true
			}
		}

		if (i > 3000 || !emptySpace){
			break
		}
		tempWeapon := dataload.AllWeaponsList[rand.Intn(len(dataload.AllWeaponsList))]

		attached, tempMountList = mounts.AttachIfCanFit(tempMountList, tempWeapon)


		if(attached){tempWeapons = append(tempWeapons, tempWeapon)}

	}
	return tempWeapons, tempMountList
}

/*func CalculateAvailableMountSpace(openMounts []string) map[string]int {
	mountSpace := map[string]int {"main"}
}*/

//function to remove a mount from the availability list
func removeMount(mountList []string, removalItem string) []string{
	for i, currentItem := range mountList {
        if strings.ToLower(currentItem) == removalItem {
            return append(mountList[:i], mountList[i+1:]...)
        }
    }
    return mountList
}

//function to add a mount to the availability list
func addMount(mountList []string, newItem string) []string{
	mountList = append(mountList, newItem)
	return mountList
}

//function to join weapon names per mount into a string for output
func JoinWeapons(weaponList []dataload.Weapon, delim string) string{
	stringifyed := ""

	for i, wep := range weaponList {
		if (i == 0){
			stringifyed = wep.Name
		}else{
			stringifyed = stringifyed+delim+wep.Name
		}
	}
	return stringifyed
}