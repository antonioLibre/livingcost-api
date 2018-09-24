package dao

import (
	"log"

	"gopkg.in/mgo.v2"
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

type LivingcostsDAO struct {
	Server   string
	Database string
}

var db *mgo.Database

const (
	COLLECTION = "livingcosts"
)

// Establish a connection to database
func (m *LivingcostsDAO) Connect() {
	session, err := mgo.Dial(m.Server)
	if err != nil {
		log.Fatal(err)
	}
	db = session.DB(m.Database)
}

// Find list of livingcosts
func (m *LivingcostsDAO) FindAll() ([]Livingcost, error) {
	var livingcosts []Livingcost
	err := db.C(COLLECTION).Find(bson.M{}).All(&livingcosts)
	return livingcosts, err
}

// Find a Livingcost by its id
func (m *LivingcostsDAO) FindById(id string) (Livingcost, error) {
	var Livingcost Livingcost
	err := db.C(COLLECTION).FindId(bson.ObjectIdHex(id)).One(&Livingcost)
	return Livingcost, err
}

// Insert a Livingcost into database
func (m *LivingcostsDAO) Insert(Livingcost Livingcost) error {
	err := db.C(COLLECTION).Insert(&Livingcost)
	return err
}

// Delete an existing Livingcost
func (m *LivingcostsDAO) Delete(Livingcost Livingcost) error {
	err := db.C(COLLECTION).Remove(&Livingcost)
	return err
}

// Update an existing Livingcost
func (m *LivingcostsDAO) Update(Livingcost Livingcost) error {
	err := db.C(COLLECTION).UpdateId(Livingcost.ID, &Livingcost)
	return err
}
