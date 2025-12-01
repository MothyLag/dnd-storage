package repository

import (
	"dnd-storage/src/domain/entities"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)
type Character struct {
	OwnerId		  primitive.ObjectID   `bson:"ownerId,omitempty"`
	ID            primitive.ObjectID   `bson:"_id,omitempty"`
	Name          string               `bson:"name,omitempty"`
	Level         int                  `bson:"level,omitempty"`
	CodeName      string               `bson:"codeName,omitempty"`
	Class         string               `bson:"class,omitempty"`
	Race          string               `bson:"race,omitempty"`
	Background    string               `bson:"background,omitempty"`
	Age           int                  `bson:"age,omitempty"`
	Speed         int                  `bson:"speed,omitempty"`
	Stats         CharacterStats       `bson:"stats,omitempty"`
	Hp            CharacterHP          `bson:"hp,omitempty"`
	Ac            int                  `bson:"ac,omitempty"`
	PowerDie      string               `bson:"powerDie,omitempty"`
	Rp            CharacterRp          `bson:"rp,omitempty"`
	SaveDC        int                  `bson:"saveDC,omitempty"`
	AttackMode    int                  `bson:"attackMode,omitempty"`
	SavingThrows  CharacterSavingThrows `bson:"savingThrows,omitempty"`
	Competences   []string             `bson:"competences,omitempty"`
	ImageUrl      string               `bson:"imageUrl,omitempty"`
}

type CharacterStats struct {
	INT  int `bson:"int"`
	COMB int `bson:"comb"`
	VEL  int `bson:"vel"`
	RES  int `bson:"res"`
	CAR  int `bson:"car"`
}

type CharacterHP struct {
	Current int `bson:"current"`
	Max     int `bson:"max"`
	Temp    int `bson:"temp"`
}

type CharacterRp struct {
	Current int `bson:"current"`
	Max     int `bson:"max"`
}

type CharacterSavingThrows struct {
	Primary   string `bson:"primary"`
	Secondary string `bson:"secondary"`
}

type CharacterMongoRepository struct{
	collection *mongo.Collection
}

func NewCharacterMongoRepository(db *mongo.Database) *CharacterMongoRepository{
	return &CharacterMongoRepository{collection: db.Collection("characters")}
}

func (m *Character) toCharacter() entities.Character{
	return entities.Character{
		OwnerId:      m.ID.Hex(),
		ID:			  m.ID.Hex(),
		Name:         m.Name,
		Level:        m.Level,
		CodeName:     m.CodeName,
		Class:        m.Class,
		Race:         m.Race,
		Background:   m.Background,
		Age:          m.Age,
		Speed:        m.Speed,
		Stats:        m.Stats.ToEntity(),
		Hp:           m.Hp.ToEntity(),
		AC:           m.Ac,
		PowerDie:     m.PowerDie,
		Rp:           m.Rp.ToEntity(),
		SaveDC:       m.SaveDC,
		AttackMode:   m.AttackMode,
		SavingThrows: m.SavingThrows.ToEntity(),
		Competences:  m.Competences,
	}
}

func (s *CharacterStats) ToEntity() entities.CharacterStats {
	return entities.CharacterStats{
		INT:  s.INT,
		COMB: s.COMB,
		VEL:  s.VEL,
		RES:  s.RES,
		CAR:  s.CAR,
	}
}

func (h *CharacterHP) ToEntity() entities.CharacterHP {
	return entities.CharacterHP{
		Current: h.Current,
		Max:     h.Max,
		Temp:    h.Temp,
	}
}

func (r *CharacterRp)ToEntity() entities.CharacterRp {
	return entities.CharacterRp{
		Current: r.Current,
		Max:     r.Max,
	}
}

func (s *CharacterSavingThrows) ToEntity() entities.CharacterSavingThrows {
	return entities.CharacterSavingThrows{
		Primary:   s.Primary,
		Secondary: s.Secondary,
	}
}

func (CharacterStats) FromEntity(e entities.CharacterStats) CharacterStats {
	return CharacterStats{
		INT:  e.INT,
		COMB: e.COMB,
		VEL:  e.VEL,
		RES:  e.RES,
		CAR:  e.CAR,
	}
}

func (CharacterHP) FromEntity(e entities.CharacterHP) CharacterHP {
	return CharacterHP{
		Current: e.Current,
		Max:     e.Max,
		Temp:    e.Temp,
	}
}

func (CharacterRp) FromEntity(e entities.CharacterRp) CharacterRp {
	return CharacterRp{
		Current: e.Current,
		Max:     e.Max,
	}
}

func (CharacterSavingThrows) FromEntity(e entities.CharacterSavingThrows) CharacterSavingThrows {
	return CharacterSavingThrows{
		Primary:   e.Primary,
		Secondary: e.Secondary,
	}
}

func (Character) FromEntity(e entities.Character) (Character,error){
	var oid primitive.ObjectID
	var err error

	if e.ID != ""{
		oid,err = primitive.ObjectIDFromHex(e.ID)	
		if err != nil{
			return Character{},err
		}
	}

	ownerId,err := primitive.ObjectIDFromHex(e.OwnerId)
	if err != nil{
		return Character{},err
	}

	return Character{
		OwnerId:      ownerId,
		ID:           oid,
		Name:         e.Name,
		Level:        e.Level,
		CodeName:     e.CodeName,
		Class:        e.Class,
		Race:         e.Race,
		Background:   e.Background,
		Age:          e.Age,
		Speed:        e.Speed,
		Stats:        CharacterStats{}.FromEntity(e.Stats),
		Hp:           CharacterHP{}.FromEntity(e.Hp),
		Ac:           e.AC,
		PowerDie:     e.PowerDie,
		Rp:           CharacterRp{}.FromEntity(e.Rp),
		SaveDC:       e.SaveDC,
		AttackMode:   e.AttackMode,
		SavingThrows: CharacterSavingThrows{}.FromEntity(e.SavingThrows),
		Competences:  e.Competences,

	},nil
}
