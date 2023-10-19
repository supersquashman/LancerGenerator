package mounts

import(
	"dataload"
	"math/rand"
	"reflect"
	"strings"
	"slices"
	"fmt"
)

//constants for comparative values
const dMain = "main"
const dHeavy = "heavy"
const dFlex = "flex"
const dAux = "aux"
const dAuxiliary = "auxiliary"
const dAuxAux = "aux/aux"
const dMainAux = "main/aux"
const dIntegrated = "integrated"
const dSuperheavy = "superheavy"

type MountObject struct{
	Type string
	Weapons []dataload.Weapon
	AvailableSlots []string
}

//Check if a mount has space for the weapon passed in
func CanFit(mount MountObject, weapon dataload.Weapon) (bool, string){
	itFits :=  false //slices.Contains(mount.AvailableSlots, weapon.Type)
	slotList := mount.AvailableSlots
	whichSlotCanFit := ""
	weaponType := strings.ToLower(weapon.Mount)

	switch weaponType{
		case dMain: //one main or aux
			if (slices.Contains(slotList, dMain)){
				itFits = true
				whichSlotCanFit = dMain
			} else if (slices.Contains(slotList, dFlex)){
				itFits = true
				whichSlotCanFit = dFlex
			} else if (slices.Contains(slotList, dHeavy)){
				itFits = true
				whichSlotCanFit = dHeavy
			}else{}

		case dHeavy: //one heavy, main, or aux
			if (slices.Contains(slotList, dHeavy)){
				itFits = true
				whichSlotCanFit = dHeavy
			}

		case dAux: //aux weapon for, this should be the main use-case for this type.
			if (slices.Contains(slotList, dAux)){
				itFits = true
				whichSlotCanFit = dAux
			} else if (slices.Contains(slotList, dFlex)){
				itFits = true
				whichSlotCanFit = dFlex
			} else if (slices.Contains(slotList, dMain)){
				itFits = true
				whichSlotCanFit = dMain
			} else if (slices.Contains(slotList, dHeavy)){
				itFits = true
				whichSlotCanFit = dHeavy
			} else{}

		case dAuxiliary: //aux weapon for, this should be the main use-case for this type.
			if (slices.Contains(slotList, dAux)){
				itFits = true
				whichSlotCanFit = dAux
			} else if (slices.Contains(slotList, dFlex)){
				itFits = true
				whichSlotCanFit = dFlex
			} else if (slices.Contains(slotList, dMain)){
				itFits = true
				whichSlotCanFit = dMain
			} else if (slices.Contains(slotList, dHeavy)){
				itFits = true
				whichSlotCanFit = dHeavy
			} else{}

		default:
			fmt.Println("Unfound type: " + weaponType)
	}

	return itFits, whichSlotCanFit
}

//function to check if a mount is full not sure if needed, but still missing logic
func IsFull(mount MountObject) bool{
	isFull := reflect.DeepEqual(mount.AvailableSlots, []string{""})

	return isFull
}

//Function to verify if there is still space in a slot
func IsEmpty(mount MountObject) bool{
	isEmpty := reflect.DeepEqual(mount.AvailableSlots, GetMountSlotsByType(mount.Type))

	return isEmpty
}

//need to use this for general use to accomodate Superheavy's
func AttachIfCanFit(mountList []MountObject, tempWeapon dataload.Weapon) (bool, []MountObject){
	attached := false

	if (strings.ToLower(tempWeapon.Mount) == dSuperheavy){
		freeHeavy := false
		freeOther := false
		var heavyMount MountObject
		var otherMount MountObject

		for _, mount := range mountList{
			if (IsEmpty(mount)){
				if (!freeHeavy && mount.Type == dHeavy){
					freeHeavy = true
					heavyMount = mount
					continue
				}else{
					freeOther = true
					otherMount = mount
				}
			}else{attached=false}
		}
		if (freeHeavy && freeOther){
			heavyMount.Weapons = append(heavyMount.Weapons, tempWeapon)
			otherMount.Weapons = append(otherMount.Weapons, tempWeapon)

			heavyMount.AvailableSlots = []string{""}
			otherMount.AvailableSlots = []string{""}
			attached = true
		}
	} else {
		for i, mount := range mountList{
			if(attached){break}
			fits, whichSlot := CanFit(mount, tempWeapon)
			if fits{
				mount.AvailableSlots = removeSlot(mount.AvailableSlots, whichSlot)
				if (whichSlot == dFlex){
					if (tempWeapon.Type == dAux){
						if (rand.Intn(2)>0){mount.AvailableSlots = addSlot(mount.AvailableSlots, dAux)}
					}
				}
				mount.Weapons = append(mount.Weapons, tempWeapon)
				attached = true
				mountList[i] = mount
			}
		}
	}

	return attached, mountList
}

//function to standardize the slot values
func GetMountSlotsByType(mountType string, customSlots ...[]string) []string{
	slotsArray := []string{""}
	customCase := "custom"
	switch strings.ToLower(mountType){
		case dMain: //one main
			slotsArray = []string{dMain}
		case dHeavy: //one heavy, main, or aux
			slotsArray = []string{dHeavy}
		case dAux: //one aux, probably not a real case, but keeping for sake of homebrews
			slotsArray = []string{dAux}
		case dAuxAux: //up to two aux
			slotsArray = []string{dAux,dAux}
		case dMainAux: //one main and one aux, or two aux
			slotsArray = []string{dMain,dAux}
		case dFlex: //one main or up to two aux
			slotsArray = []string{dFlex}
		case dIntegrated: //special, including for now, may remove

		case dSuperheavy: //no mount type; weapons of this type require heavymount and a mount of any other size

		case customCase:
			slotsArray = customSlots[0]

		
	}
	return slotsArray
}

//function to initialize new mounts based on type
func NewMount(mType string) MountObject{
	var defMount MountObject

	defMount.Type = mType
	defMount.AvailableSlots = GetMountSlotsByType(mType)

	return defMount
}

//function to remove a slot from a mount's available slots
func removeSlot(slotList []string, removalItem string) []string{
	for i, currentItem := range slotList {
        if strings.ToLower(currentItem) == removalItem {
            return append(slotList[:i], slotList[i+1:]...)
        }
    }
    return slotList
}

//function to add a slot
func addSlot(slotList []string, newItem string) []string{
	slotList = append(slotList, newItem)
	return slotList
}

//function for formatting output of weapon names
func JoinWeapons(containerMount MountObject, delim string) string{
	stringifyed := ""
	weaponList := containerMount.Weapons

	for i, wep := range weaponList {
		if (i == 0){
			stringifyed = wep.Name
		}else{
			stringifyed = stringifyed+delim+wep.Name
		}
	}
	return stringifyed
}