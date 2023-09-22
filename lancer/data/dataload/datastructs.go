package dataload

 // Define a struct to hold the JSON data
 type Frame struct {
	ID string `json:"id"`
	LicenseLevel int `json:"licence_level"`
	LiscenceID string `json:"license_id"`
	Variant string `json:"variant"`
	Source string `"jsonc:source"`
	Name string `json:"name"`
	MechType []string `json:"mechtype"`
	Specialty bool `json:"specialty"` //| IPrerequisite // revisit adding IPrerequisite struct for outputting homebrew
	Description string `json:"description"`
	Mounts []string `json:"mounts"`
	Stats FrameStats `json:"stats"`
	Traits []FrameTraits `json:"traits"`
	CoreSystem CoreSystemData `json:"core_system"`
	ImageURL string `json:"image_url"?` //shouldn't be necessary currently
	YPOS float32 `json:"y_pos"?` // used for vertical alignment of the mech in banner views (like in the new mech selector)
  }

  type FrameStats struct{
	Size int `json:"size"`
	Structure int `json:"structure"`
	Stress int `json:"stress"`
	Armor int `json:"armor"`
	HP int `json:"hp"`
	Evaison int `json:"evasion"`
	Edef int `json:"edef"`
	HeatCap int `json:"heatcap"`
	RepairCap int `json:"repcap"`
	SensorRange int `json:"sensor_range"`
	TechAttack int `json:"tech_attack"`
	Save int `json:"save"`
	Speed int `json:"speed"`
	SP int `json:"sp"`
  }

  type FrameTraits struct{
	Name string `json:"name"`
	Description string `json:"description"` // v-html
	UseFreq string `json:"use"?` //options include 'Turn' | 'Next Turn' | 'Round' | 'Next Round' | 'Scene' | 'Encounter' | 'Mission',
	IActions []IActionData `json:"actions"?`  //IActionData[],
	Bonuses []IBonusData `json:"bonuses"?`
	Synergies []ISynergyData `json:"synergies"?`
	Deployables []IDeployableData `json:"deployables"?`
	Counters []ICounterData `json:"counters"?`
	Integrated []string `json:"integrated"?` //It is possible to "chain" equipment through integrated arrays, and is therefore possible to loop infinitely, which will cause the app to crash. Keep this in mind when building LCPs.
	SpecialEquipment []string `json:"special_equipment"?`
  }

  type CoreSystemData struct{
	Name string `json:"name"`
	ActiveName string `json:"active_name"`
	ActiveEffect string `json:"active_effect"` // v-html
	Activation string `json:"activation"`
	Description string `json:"description"?` // v-html
	Deactivation string `json:"deactivation"?`
	UseFreq string `json:"use"?` // 'Round' | 'Next Round' | 'Scene' | 'Encounter' | 'Mission',
	ActiveActions []IActionData `json:"active_actions"?`
	ActiveBonuses []IBonusData `jason:"active_bonuses"?`
	ActiveSynergies []ISynergyData `json:"active_synergies"?`
	PassiveName string `json:"passive_name"?`
	PassiveEffect string `json:"passive_effect"?` // v-html,
	PassiveActions []IActionData `json:"passive_actions"?`
	PassiveBonuses []IBonusData `json:"passive_bonuses"?`
	PassiveSynergies []ISynergyData `json:"passive_synergies"`
	Deployables []IDeployableData `json:"deployables"?`
	Counter []ICounterData `json:"counters"?`
	Integrated []string `json:"integrated"?`
	SpecialEquipment []string `json:"special_equipment"?`
	Tags []ITagData `json:"tags"?`
  }

  type IActionData struct{
	Name string `json:"name"?`
	Activation string `json:"activation"` //ActivationType,
	Detail string `json:"detail"` // v-html
	Cost int `json:"cost"?`
	Pilot bool `json:"pilot"?`
	SynergyLocations []string `json:"synergy_locations"?`
	TechAttack bool `json:"tech_attack"?`
	Log []string `json:"log"?`
  }

  type IBonusData struct{
	ID string `json:"id"`
	Val interface{} `json:"val"`
	DamageTypes []string `json:"damage_types"?` //damage type values from array?
	RangeTypes []string `json:"range_types"?` //same as DamageTypes
	WeaponTypes []string `json:"weapon_types"?` //same as above
	WeaponSizes []string `json:"weapon_sizes"?`
	Overwrite bool `json:"overwrite"?`
	Replace bool `json:"replace"?`
	Traits []FrameTraits `json:"traits"`
  }

  type ISynergyData struct{
	Locations []string `json:"locations"` // see below,
	Detail string `json:"detail"` // v-html
	WeaponTypes []string `json:"weapon_types"?`
	SystemTypes []string `json:"system_types"?`
	WeaponSize []string `json:"weapon_sizes"?`
  }

  type IDeployableData struct{
	Name string `json:"name"`
	UIType string `json:"type"` // this is for UI furnishing only
	Detail string `json:"detail"`
	Size int `json:"size"` // not required for Mines
	Activation string `json:"activation"?`
	Deactivation string `json:"deactivation"?`
	Recall string `json:"recall"?`
	Redeploy string `json:"redeploy"?`
	Instances int `json:"instances"?`
	Cost int `json:"cost"?`
	Armor int `json:"armor"?`
	HP int `json:"hp"?`
	Evasion int `json:"evasion"?`
	Edef int `json:"edef"?`
	HeatCap int `json:"heatcap"?`
	RepairCap int `json:"repcap"?`
	SensorRange int `json:"sensor_range"?`
	TechAttack int `json:"tech_attack"?`
	Save int `json:"save"?`
	Speed int `json:"speed"?`
	Pilot bool `json:"pilot"?`
	Mech bool `json:"mech"?`
	Actions []IActionData `json:"actions"?`
	Bonuses []IBonusData `json:"bonuses"?`
	Synergies []ISynergyData `json:"synergies"`
	Counters []ICounterData `json:"counters"?`
	Tags []ITagData `json:"tags"?`
  }

  type ICounterData struct{
	ID string `json:"id"`
	Name string `json:"name"`
	DefaultValue int `json:"default_value"?`
	Min int `json:"min"?`
	Max int `json:"max"?`
  }

  type ActivationType struct{
	//Enum array maybe?
  }

  type ITagData struct{
	ID string `json:"id"`
  	Val string `json:"val"`
  }

  type IRangeData struct{
	Type string `json:"type"`
	Val int `json:"val"`
  }

  type IDamageData struct{
	Type string `json:"type"`
	Val int `json:"val"`
  }

  type Weapon struct{
	ID string `json:"id"`
	Name string `json:"name"`
	Source string `json:"source"` // must be the same as the Manufacturer ID to sort correctly
	License string `json:"license"` // reference to the Frame name of the associated license
	LicenseID string `json:"license_id"` // reference to the Frame id of the associated license
	LicenseLevel int `json:"license_level"` // set to zero for this item to be available to a LL0 character
	Mount string `json:"mount"` // MountType ?? enum array eventually maybe
	Type string `json:"type"` //  WeaponType ?? enum array eventually maybe?
	Cost int `json:"cost"?`
	Barrage bool `json:"barrage"?`
	Skirmish bool `json:"skirmish"?`
	NoAttack bool `json:"no_attack"?`
	NoMods bool `json:"no_mods"?`
	NoCoreBonus bool `json:"no_core_bonus"?`
	Damage []IDamageData `json:"damage"?`
	Range IRangeData `json:"range"?`
	Tags ITagData `json:"tags"?`
	SP int `json:"sp"?`
	Description string `json:"description"` // v-html
	Effect string `json:"effect"?` // v-html
	OnAttack string `json:"on_attack"?`
	OnHit string `json:"on_hit"?`
	OnCrit string `json:on_crit"?`
	Actions []IActionData `json:"actions"?`
	Bonuses []IBonusData `json:"bonuses"?`
	Synergies []ISynergyData `json:"synergies"?`
	Deployables []IDeployableData `json:"deployables"?`
	Counters []ICounterData `json:"counters"?`
	Integrated []string `json:"integrated"?`
	SpecialEquipment []string `json:"special_equipment"?`
	Profiles []IWeaponProfile `json:"profiles"?` 
  }

  type IWeaponProfile struct{
	Name string `json:"name"`
	Effect string `json:"effect"?`
	Skirmish bool `json:"skirmish"?`
	Barrage bool `json:"barrage"?`
	Cost int `json:"cost"?`
	OnAttack string `json:"on_attack"?`
	OnHit string `json:"on_hit"?`
	OnCrit string `json:on_crit"?`
	Damage []IDamageData `json:"damage"?`
	Range []IRangeData `json:"range"?`
	Actions []IActionData `json:"actions"?`
	Bonuses []IBonusData `json:"bonuses"?`
	Synergies []ISynergyData `json:"synergies"?`
	Deployables []IDeployableData `json:"deployables"?`
	Counters []ICounterData `json:"counters"?`
	Integrated []string `json:"integrated"?`
	SpecialEquipment []string `json:"special_equipment"?`
  }

  type WeaponTypeData struct{
	//enum array
  }

  type WeaponSizeData struct{
	//enum array eventually
  }

  type SystemTypeData struct{
	//enum array eventually
  }

  type PilotGear struct{
	ID string `json:"id"`
	Name string `json:"name"` // v-html
	Type string `json:"type"` // "Weapon" || "Armor" || "Gear"
	Description string `json:"description"?`
	Tags []ITagData `json:"tags"?`
	Range []IRangeData `json:"range"?`
	Damage []IDamageData `json:"damage"?` 
	Actions []IActionData `json:"actions"?` // these are only available to UNMOUNTED pilots
	Bonuses []IBonusData `json:"bonuses"?` // these bonuses are applied to the pilot, not parent system
	Synergies []ISynergyData `json:"synergies"?`
	Deployables []IDeployableData `json:"deployables"?` // these are only available to UNMOUNTED pilots
  }



  type FramesList struct{
	AllFrames Frame `json:"frames"`
  }