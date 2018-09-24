package models

import (
	"gopkg.in/mgo.v2/bson"
)

type Livingcost struct {
	ID              bson.ObjectId `bson:"_id" json:"id"`
	Barrio          string        `bson:"barrio" json:"barrio"`
	Estrato         int           `bson:"estrato" json:"estrato"`
	Localidad       int           `bson:"localidad" json:"localidad"`
	SectroCatastral string        `bson:"sectroCatastral" json:"sectroCatastral"`
	Valorm2         int           `bson:"valorm2" json:"valorm2"`
}
