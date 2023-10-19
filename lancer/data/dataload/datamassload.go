package dataload

import (
	"path/filepath"
	"io/fs"
	"strings"
  )

type load_func func(string)

func Load_All() {
	Load_All_Frames()
	//Load_All_PilotGear()
	//Load_All_Reserves()
	//Load_All_Systems()
	//Load_All_Talents()
	Load_All_Weapons()
  }

func Load_All_Frames(){
	var compareVal = "frames"
	load_all_base(Load_Frames, compareVal)
}

func Load_All_PilotGear(){
	//var compareVal = "pilot_gear"
	//load_all_base(Load_PilotGear, compareVal)
}

func Load_All_Reserves(){
	//var compareVal = "reserves"
	//load_all_base(Load_Reserves, compareVal)
}

func Load_All_Systems(){
	//var compareVal = "systems"
	//load_all_base(Load_Systems, compareVal)
}

func Load_All_Talents(){
	//var compareVal = "talents"
	//load_all_base(Load_Talents, compareVal)
}

func Load_All_Weapons(){
	var compareVal = "weapons"
	load_all_base(Load_Weapons, compareVal)
}

func load_all_base(specFunc load_func, comparison string){
	filepath.WalkDir("./data/sources", func (Fpath string, di fs.DirEntry, err error) error {
		if !di.IsDir(){
			info, erri := di.Info()
			if strings.Contains(strings.ToLower(info.Name()),comparison){
				specFunc(Fpath)	
			}
			if erri != nil {err = erri}
		}
		return err
	})
}