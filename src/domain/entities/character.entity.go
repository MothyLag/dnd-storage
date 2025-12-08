package entities

type Character struct{
	OwnerId string `json:"ownerId,omitempty"`
	ID string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	Level int `json:"level,omitempty"`
	CodeName string `json:"codeName,omitempty"`
	Class string `json:"class,omitempty"`
	Race string `json:"race,omitempty"`
	Background string `json:"background,omitempty"`
	Age int `json:"age,omitempty"`
	Speed int `json:"speed,omitempty"`
	Stats CharacterStats `json:"stats,omitempty"`
	Hp CharacterHP `json:"hp,omitempty"`
	AC int `json:"ac,omitempty"`
	PowerDie string `json:"powerDie,omitempty"`
	Rp CharacterRp `json:"rp,omitempty"`
	SaveDC int `json:"saveDc,omitempty"`
	AttackMode int `json:"attackMode,omitempty"`
	SavingThrows CharacterSavingThrows `json:"saving_throws,omitempty"`
	Competences []string `json:"competences,omitempty"`
	Image FileUpload `json:"image,omitempty"`
}

type CharacterStats struct{
	INT int `json:"INT"`
	COMB int `json:"COMB"`
	VEL int `json:"VEL"`
	RES int `json:"RES"`
	CAR int `json:"CAR"`
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
