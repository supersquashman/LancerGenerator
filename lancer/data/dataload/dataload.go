package dataload

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
  )

  var AllFramesList []Frame
  var AllPilotGearList []PilotGear
  var AllWeaponsList []Weapon
  var AllSystemsList []System
  var AllTalentsList []Talent
  var AllReservesList []Reserves

  const dFRAME = "frame"
  const dSYSTEM = "system"
  const dPILOTGEAR = "pilotgear"
  const dWEAPON = "weapon"
  const dTALENT = "talent"
  const dRESERVE = "reserve"
  const dSITREP = "sitrep" //to-do
  const dSKILLTRIGGER = "skilltrigger" //to-do
  const dSTATUS = "status" //to-do

  func Load_Frames(path string){
	file_load(path, dFRAME)

	currentFrame := AllFramesList[0]
	fmt.Println("ID: "+currentFrame.ID)
	fmt.Println("Name: "+currentFrame.Name)
	fmt.Println("Description: "+currentFrame.Description)

  }

  func Load_PilotGear(path string){
	file_load(path, dPILOTGEAR)
  }

  func Load_Systems(path string){
	file_load(path, dSYSTEM)
  }

  func Load_Talents(path string){
	file_load(path, dTALENT)
  }

  func Load_Weapons(path string){
	file_load(path, dWEAPON)
  }

  func Load_Reserves(path string){
	file_load(path, dRESERVE)
  }

  func file_load(path string, dtype string) {

	file, err := os.Open(path)
	if err != nil {
	  log.Fatalf("Error opening file: %v", err)
	}
	defer file.Close()

	// Read the file content
	content, err := ioutil.ReadAll(file)
	if err != nil {
	  log.Fatalf("Error reading file: %v", err)
	}

	switch dtype{
		case dFRAME:
			var tempFrameList []Frame

			err = json.Unmarshal(content, &tempFrameList)
			if err != nil{
				log.Fatalf("Error unmarshalling data:  %v", err)
			}

			AllFramesList = append(AllFramesList[:len(AllFramesList):len(AllFramesList)],tempFrameList...)
		case dPILOTGEAR:
			var tempPilotGearList []PilotGear

			err = json.Unmarshal(content, &tempPilotGearList)
			if err != nil{
				log.Fatalf("Error unmarshalling data:  %v", err)
			}

			AllPilotGearList = append(AllPilotGearList[:len(AllPilotGearList):len(AllPilotGearList)],tempPilotGearList...)
		case dWEAPON:
			var tempWeaponsList []Weapon

			err = json.Unmarshal(content, &tempWeaponsList)
			if err != nil{
				log.Fatalf("Error unmarshalling data:  %v", err)
			}

			AllWeaponsList = append(AllWeaponsList[:len(AllWeaponsList):len(AllWeaponsList)],tempWeaponsList...)
		case dSYSTEM:
			var tempSystemsList []System

			err = json.Unmarshal(content, &tempSystemsList)
			if err != nil{
				log.Fatalf("Error unmarshalling data:  %v", err)
			}

			AllSystemsList = append(AllSystemsList[:len(AllSystemsList):len(AllSystemsList)],tempSystemsList...)
		case dRESERVE:
			var tempReservesList []Reserves

			err = json.Unmarshal(content, &tempReservesList)
			if err != nil{
				log.Fatalf("Error unmarshalling data:  %v", err)
			}

			AllReservesList = append(AllReservesList[:len(AllReservesList):len(AllReservesList)],tempReservesList...)
		case dTALENT:
			var tempTalentsList []Talent

			err = json.Unmarshal(content, &tempTalentsList)
			if err != nil{
				log.Fatalf("Error unmarshalling data:  %v", err)
			}

			AllTalentsList = append(AllTalentsList[:len(AllTalentsList):len(AllTalentsList)],tempTalentsList...)
	}

  }