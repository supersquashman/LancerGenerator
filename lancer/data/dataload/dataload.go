package dataload

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
  )
  
  /*func load() {
	// Open the JSON file
	file, err := os.Open("data.json")
	if err != nil {
	  log.Fatalf("Error opening file: %v", err)
	}
	defer file.Close()
  
	// Read the file content
	content, err := ioutil.ReadAll(file)
	if err != nil {
	  log.Fatalf("Error reading file: %v", err)
	}
  
	// Parse the JSON content
	var student Student
	err = json.Unmarshal(content, &student)
	if err != nil {
	  log.Fatalf("Error unmarshalling JSON: %v", err)
	}
  
	// Use the parsed JSON data
	fmt.Printf("Name: %s\nAge: %d\nEmail: %s\n", 
			   student.Name, student.Age, student.Email)
  }*/

  func Load_All() {
	Load_Frames()
  }

  func Load_Frames() {
	//fmt.Println(os.Getwd())

	file, err := os.Open("./data/sources/frames.json")
	if err != nil {
	  log.Fatalf("Error opening file: %v", err)
	}
	defer file.Close()

	// Read the file content
	content, err := ioutil.ReadAll(file)
	if err != nil {
	  log.Fatalf("Error reading file: %v", err)
	}

	var AllFramesTest []Frame

	err = json.Unmarshal(content, &AllFramesTest)

	if err != nil{
		log.Fatalf("Error unmarshalling data:  %v", err)
	}

	currentFrame := AllFramesTest[0]
	fmt.Println("ID: "+currentFrame.ID)
	fmt.Println("Name: "+currentFrame.Name)
	fmt.Println("Description: "+currentFrame.Description)
  }