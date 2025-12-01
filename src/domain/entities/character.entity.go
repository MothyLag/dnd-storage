package entities

type Character struct{
	OwnerId string `json:"ownerId,omitempty"`
	ID string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	Level int `json:"level,omitempty"`
	CodeName string `json:"code_name,omitempty"`
	Class string `json:"class,omitempty"`
	Race string `json:"race,omitempty"`
	Background string `json:"background,omitempty"`
	Age int `json:"age,omitempty"`
	Speed int `json:"speed,omitempty"`
	Stats CharacterStats `json:"stats,omitempty"`
	Hp CharacterHP `json:"hp,omitempty"`
	AC int `json:"ac,omitempty"`
	PowerDie string `json:"power_die,omitempty"`
	Rp CharacterRp `json:"rp,omitempty"`
	SaveDC int `json:"save_dc,omitempty"`
	AttackMode int `json:"attack_mode,omitempty"`
	SavingThrows CharacterSavingThrows `json:"saving_throws,omitempty"`
	Competences []string `json:"competences,omitempty"`
	Image FileUpload `json:"image,omitempty"`
}

type CharacterStats struct{
	INT int `json:"int"`
	COMB int `json:"comb"`
	VEL int `json:"vel"`
	RES int `json:"res"`
	CAR int `json:"car"`
}

type CharacterHP struct{
	Current int `json:"current"`
	Max int `json:"max"`
	Temp int `json:"temp"`
}

type CharacterRp struct{
	Current int `json:"current"`
	Max int `json:"max"`
}

type CharacterSavingThrows struct{
	Primary string `json:"primary"`
	Secondary string `json:"secondary"`
}
