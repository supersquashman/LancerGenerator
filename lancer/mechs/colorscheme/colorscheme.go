package colorscheme

import (
	"fmt"
	"math/rand"
	// "strings"
)


//init a color list for length
//have a list of schemes that include number of colors needed to select
//scheme should also probably include guide to logic for interpolation
func GetRandomColorscheme() string{
	colorsList := []string{"black","brown","gray","green","red","blue","white","yellow","purple","teal","mauve","turquoise","sapphire","jade","emerald","sky blue","maroon","beige","orange"}
	schemeList := []string{"solid %s","faded %s","%s with %s accents"}
	randScheme := schemeList[rand.Intn(len(schemeList))]
	firstColor := colorsList[rand.Intn(len(colorsList))]
	secondColor := colorsList[rand.Intn(len(colorsList))]
	return fmt.Sprintf(randScheme, firstColor,secondColor)
}