package colorscheme

import (
	"fmt"
	"math/rand"
	// "strings"
)

type ColorScheme struct{
	SchemeString string `json:"scheme_string"`
	ColorCount int `json:"color_count"`
}

func NewColorScheme(schemeString string, colorCount int) ColorScheme{
	if(colorCount < 1){colorCount=1}
	scheme := ColorScheme{SchemeString:schemeString, ColorCount:colorCount}

	return scheme
}

//init a color list for length
//have a list of schemes that include number of colors needed to select
//scheme should also probably include guide to logic for interpolation
func GetRandomColorscheme() string{
	colorsList := []string{"black","brown","gray","green","red","blue","white","yellow","purple","teal","mauve","turquoise","sapphire","jade","emerald","sky blue","maroon","beige","orange"}
	schemeList := []ColorScheme{
		NewColorScheme("solid %s",1),
		NewColorScheme("faded %s",1),
		NewColorScheme("metallic %s",1),
		NewColorScheme("faded %s with metallic %s highlights",2),
		NewColorScheme("matte %s with metallic %s highlights",2),
		NewColorScheme("matte %s and metallic %s",2),
		NewColorScheme("%s with %s accents", 2),
		NewColorScheme("%s and %s",2),
		NewColorScheme("%s and %s with %s highlights",3),
		NewColorScheme("%s and %s with %s accents",3)}
	randScheme := schemeList[rand.Intn(len(schemeList))]
	selectedColors := make([]any, randScheme.ColorCount)
	for i:=0; i < randScheme.ColorCount; i++{
		selectedColors[i] = colorsList[rand.Intn(len(colorsList))]
	}
	// firstColor := colorsList[rand.Intn(len(colorsList))]
	// secondColor := colorsList[rand.Intn(len(colorsList))]
	// testThing := []any{firstColor,secondColor}
	return fmt.Sprintf(randScheme.SchemeString, selectedColors...)
}