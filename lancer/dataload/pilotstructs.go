package dataload

type Pilot struct{
	ID string `json:"id"`
	Name string `json:"name"`
	Level int `json:"level"`
	Callsign string `json:"callsign"`
	PlayerName string `json:"player_name"`
	Status string `json:"status"`
	Dead bool `json:"dead"`
	TextAppearance string `json:"text_appearance"`
	Notes string `json:"notes"`
	History string `json:"history"`
	Quirks []string `json:"quirks"`
	CurrentHP int `json:"current_hp"`
	Background string `json:"background"`
	Resistances []string `json:"resistances"`
	Mechs []Mech `json:"mechs"`
	SpecialEquipments SpecialEquipment `json:"special_equipment"`
	CombHistory CombatStats `json:"combat_history"`
	State PilotState `json:"state"`
	LastModified string `json:"lastModified"`
	IsDeleted string `json:"isDeleted"`
	ExpireTime string `json:"expireTime"`
	DeleteTime string `json:"deleteTime"`
	LasUpdateCloud string `json:"lasUpdate_cloud"`
	LastSync string `json:"lastSync"`
	Skills []SkillsTalentsMap `json:"skills"`
	Talents []SkillsTalentsMap `json:"talents"`
	MechSkills []int `json:"mechSkills"`
	CounterData []ICounterData `json:"counters"`
	CustomCounters []ICounterData `json:"custom_counters"`
	CoreBonuses []ICoreBonus `json:"core_bonuses"`
	Licenses []string `json:"licenses"`
	Reserves []Reserves `json:"reserves"`
	Orgs []string `json:"orgs"`
	BondID string `json:"bondId"`
	XP int `json:"xp"`
	Stress int `json:"stress"`
	IsBroken bool `json:"isBroken"`
	Burdens []string `json:"burdens"`
	BondPowers []SkillsTalentsMap `json:"bondPowers"`
	PowerSelections int `json:"powerSelections"`
	MaxStress int `json:"maxStress"`
	BondAnswers []string `json:"bondAnswers"`
	MinorIdeal string `json:"minorIdeal"`
	Clocks []ICounterData `json:"clocks"`
	Portrait string `json:"portrait"`
	CloudPortrait string `json:"cloud_portrait"`
	Loadout Loadout `json:"loadout"`
	Brews []string `json:"brews"`
}

type SkillsTalentsMap struct{
	ID string `json:"id"`
	Rank int `json:"rank"`
}

type CombatStats struct{
	Moves int `json:"moves"`
	Kills int `json:"kills"`
	Damage int `json:"damage"`
	HPDamage int `json:"hp_damage"`
	StructureDamage int `json:"structure_damage"`
	Overshield int `json:"overshield"`
	HeatDamage int `json:"heat_damage"`
	ReactorDamage int `json:"reactor_damage"`
	OverchargeUses int `json:"overcharge_uses"`
	CoreUses int `json:"core_uses"`
}

type PilotState struct{
	ActiveMechID string `json:"active_mech_id"`
	RemoteMechID string `json:"remote_mech_id"`
	Stage string `json:"stage"`
	Turn int `json:"turn"`
	Actions int `json:"actions"`
	Braced bool `json:"braced"`
	Overcharged bool `json:"overcharged"`
	Prepare bool `json:"prepare"`
	BracedCooldown bool `json:"bracedCooldown"`
	Redundant bool `json:"redundant"`
	Stats CombatStats `json:"stats"`
	Deployed []string `json:"deployed"`
}

type SpecialEquipment struct{
	PGear []PilotGear
	Frames []Frame
	MechWeapons []Weapon
	WeaponMods []Mod
	MechSystems []System
	SystemMods []Mod
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