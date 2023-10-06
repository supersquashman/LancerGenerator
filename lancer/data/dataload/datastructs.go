package dataload

 // Define a struct to hold the JSON data
 type Frame struct {
	ID string `json:"id"`
	LicenseLevel int `json:"licence_level"`
	LiscenceID string `json:"license_id"`
	Variant string `json:"variant"`
	Source string `"jsonc:source"` //manufacturer
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
	PilotUsable bool `json:"pilot"?`
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
	PilotUsable bool `json:"pilot"?`
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

  type System struct{
	ID string `json:"id"`
    Name string `json:"name"`
    Source string `json:"source"` // must be the same as the Manufacturer ID to sort correctly
    License string `json:"license"` // reference to the Frame name of the associated license
    LicenseID string `json:"license_id"` // reference to the Frame id of the associated license
    LicenseLevel int `json:"license_level"` // set to zero for this item to be available to a LL0 character
    Effect string `json:"effect"?` // v-html
    Type string `json:"type"?` //SystemType....enum eventually?
    SP int `json:"sp"?`
    Description string `json:"description"?` // v-html
    Tags ITagData `json:"tags"?`
    Actions []IActionData `json:"actions"?`
    Bonuses []IBonusData `json:"bonuses"?`
    Synergies []ISynergyData `json:"synergies"?`
    Deployables []IDeployableData `json:"deployables"?`
    Counters []ICounterData `json:"counters"?`
    Integrated []string `json:"integrated"?`
    SpecialEquipment []string `json:"special_equipment"?`
  }

  type ICoreBonus struct{
	ID string `json:"id"`
	Name string `json:"name"`
	Source string `json:"source"` // must be the same as the Manufacturer ID to sort correctly
	Effect string `json:"effect"` // v-html
	Description string `json:"description"` // v-html
	MountedEffect string `json:"mounted_effect"?`
	Actions []IActionData `json:"actions"?`
	Bonuses []IBonusData `json:"bonuses"?`
	Synergies []ISynergyData `json:"synergies"?`
	Deployables []IDeployableData `json:"deployables"?`
	Counters []ICounterData `json:"counters"?`
	Integrated []string `json:"integrated"?`
	SpecialEquipment []string `json:"special_equipment"?`
  }

  type Mech struct{
	ID string `json:"id"`
	Name string `json:"name"`
	Notes string `json:"notes"`
	GMNotes string `json:"gm_note"`
	Frame string `json:"frame"`
	Active bool `json:"active"`
	CurrentStructure int `json:"current_structure"`
	CurrentMove int `json:"current_move"`
	Boost int `json:"boost"`
	CurrentHP int `json:"current_hp"`
	Overshield int `json:"overshield"`
	CurrentStress int `json:"current_stress"`
	CurrentHeat int `json:"current_heat"`
	CurrentRepairs int `json:"current_repairs"`
	CurrentOvercharge int `json:"current_overcharge"`
	CurrentCoreEnergy int `json:"current_core_energy"`
	Statuses []string `json:"statuses"`
	Conditions []string `json:"conditions"`
	Resistances []string `json:"resistances"`
	Reactions []string `json:"reactions"`
	Burn int `json:"burn"`
	Destroyed bool `json:"destroyed"`
	Defeat string `json:"defeat"`
	Activations int `json:"activations"`
	MeltdownImminent bool `json:"meltdown_imminent"`
	ReactorDestroyed bool `json:"reactor_destroyed"`
	CoreActive bool `json:"core_active"`
	CCVer string `json:"cc_ver"`
	LastModified string `json:"lastModified"`
	IsDeleted bool `json:"isDeleted"`
	ExpireTime string `json:"expireTime"`
	DeleteTime string `json:"delteTime"`
	Portrait string `json:"portrait"`
	CloudPortrait string `json:"cloud_portrait"`
	Loadouts []Loadout `json:"loadouts"`
	ActiveLoadoutIndex int `json:"active_loadout_index"`
  }

  type Loadout struct{
	ID string `json:"id"`
	Name string `json:"name"`
	Systems []SystemInstance `json:"systems"`
	IntegratedSystems []SystemInstance `json:"integratedSystems"`
	Mounts []Mount `json:"mounts"`
	IntegratedMounts []Mount `json:"integratedMounts"`
	ImprovedArmament []Mount `json:"improved_armament"`
	SuperHeavyMounting []Mount `json:"superheavy_mounting"`
	IntegratedWeapon []Mount `json:"integratedWeapon"`
  }

  type Mount struct{
	MountType string `json:"mount_type"`
	Lock bool `json:"lock"`
	Slots []MountSlot `json:"slots"`
  }

  type MountSlot struct{
	Size string `json:"size"`
	Weapon WeaponInstance `json:"weapon"`
	Extra []string `json:"extra"`
	BonusEffects []string `json:"bonus_effects"`
  }

  type SystemInstance struct{
	ID string `json:"id"`
	Uses int `json:"uses"`
	Destroyed bool `json:"destroyed"`
	Cascading bool `json:"cascading"`
	FlavorName string `json:"flavorName"`
	FlavorDescription string `json:"flavorDescription"`
  }
  
  type WeaponInstance struct{
	ID string `json:"id"`
	Destroyed bool `json:"destroyed"`
	Cascading bool `json:"cascading"`
	Loaded bool `json:"loaded"`
	Mod Mod `json:"mod"`
	FlavorName string `json:"flavorName"`
	FlavorDescription string `json:"flavorDescription"`
	CustomDamageType string `json:"customDamageType"`
	MaxUseOverride int `json:"maxUseOverride"`
	Uses int `json:"uses"`
	SelectedProfile int `json:"selectedProfile"`
  }

  type Mod struct{
	ID string `json:"id"`
	Name string `json:"name"`
	Source string `json:"source"` // Manufacturer ID
	License string `json:"license"` // Frame Name
	LicenseID string `json:"license_id"` // reference to the Frame id of the associated license
	LicenseLevel int `json:"license_level"` // set to 0 to be available to all Pilots
	SP int `json:"sp"?`
	Description string `json:"description"?` // v-html
	Effect string `json:"effect"?` // v-html
	Tags []ITagData `json:"tags"?` // tags related to the mod itself
	AllowedTypes []string `json:"allowed_types"?` // weapon types the mod CAN be applied to
	AllowedSizes []string `json:"allowed_sizes"?` // weapon sizes the mod CAN be applied to
	RestrictedTypes []string `json:"restricted_types"?` // weapon types the mod CAN NOT be applied to
	RestrictedSizes []string `json:"restricted_sizes"?` // weapon sizes the mod CAN NOT be applied to
	AddedTags []ITagData `json:"added_tags"?` // tags propagated to the weapon the mod is installed on
	AddedDamage []IDamageData `json:"added_damage"?` // damage added to the weapon the mod is installed on, see note
	AddedRange []IRangeData `json:"added_range"?` // damage added to the weapon the mod is installed on, see note
	Actions []IActionData `json:"actions"?`
	Bonuses []IBonusData `json:"bonuses"?` // these bonuses are applied to the pilot, not parent weapon
	Synergies []ISynergyData `json:"synergies"?`
	Deployables []IDeployableData `json:"deployables"?`
	Counter []ICounterData `json:"counters"?`
	Integrated []string `json:"integrated"?`
	SpecialEquipment []string `json:"special_equipment"?`
  }

  type Reserves struct{
	ID string `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"` //"Mech" | "Tactical" | "Resource" | "Bonus",
	Label string `json:"label"`
	Description string `json:"description"?` // v-html
	Consumable bool `json:"consumable"?`// defaults to false
	Actions []IActionData `json:"actions"?`
	Bonuses []IBonusData `json:"bonuses"?`
	Synergies []ISynergyData `json:"synergies"?`
	Deployables []IDeployableData `json:"deployables"?`
	Counter []ICounterData `json:"counters"?`
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

  type Talent struct{
	ID string `json:"id"`
	Name string `json:"name"`
	Description string `json:"description"` // v-html
	Ranks []IRankData `json:"ranks"`
	IconURL string `json:"icon_url"?`// Must be .svg
	Terse string `json:"terse"?` // terse text used in short descriptions. The fewer characters the better
  }

  type IRankData struct{
	Name string `json:"name"`
	Description string `json:"description"` // v-html
	Exclusing bool `json:"exclusive"?` // see below
	Actions []IActionData `json:"actions"?`
	Bonuses []IBonusData `json:"bonuses"?`
	Synergies []ISynergyData `json:"synergies"?`
	Deployables []IDeployableData `json:"deployables"?`
	Counter []ICounterData `json:"counters"?`
	Integrated []string `json:"integrated"?`
  }



  type FramesList struct{
	AllFrames Frame `json:"frames"`
  }