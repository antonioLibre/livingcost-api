package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/BurntSushi/toml"
	. "github.com/antonioLibre/livingcost-api/dao"
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2/bson"
)

// Represents database server and credentials
type Config struct {
	Server   string
	Database string
}

// Read and parse the configuration file
func (c *Config) Read() {
	if _, err := toml.DecodeFile("config.toml", &c); err != nil {
		log.Fatal(err)
	}
}

var config = Config{}
var dao = LivingcostsDAO{}

// GET list of livingcosts
func AllLivingcostsEndPoint(w http.ResponseWriter, r *http.Request) {
	livingcosts, err := dao.FindAll()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, livingcosts)
}

// GET a livingcost by its ID
func FindLivingcostEndpoint(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	livingcost, err := dao.FindById(params["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Livingcost ID")
		return
	}
	respondWithJson(w, http.StatusOK, livingcost)
}

// POST a new livingcost
func CreateLivingcostEndPoint(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var livingcost Livingcost
	if err := json.NewDecoder(r.Body).Decode(&livingcost); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	livingcost.ID = bson.NewObjectId()
	if err := dao.Insert(livingcost); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusCreated, livingcost)
}

// PUT update an existing livingcost
func UpdateLivingcostEndPoint(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	_, err := dao.FindById(params["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Livingcost ID")
		return
	}

	// barrio := params["barrio"]
	// localidad := params["localidad"]
	// sectorCatastral := params["sectorCatastral"]
	// valorm2 := params["valorm2"]
	defer r.Body.Close()
	var livingcost2 Livingcost
	livingcost2.ID = bson.ObjectIdHex(params["id"])

	if err := json.NewDecoder(r.Body).Decode(&livingcost2); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	if err := dao.Update(livingcost2); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, map[string]string{"result": "success"})
}

// DELETE an existing livingcost
func DeleteLivingcostEndPoint(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	livingcost, err := dao.FindById(params["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Livingcost ID")
		return
	}

	if err := dao.Delete(livingcost); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, map[string]string{"result": "success"})
}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJson(w, code, map[string]string{"error": msg})
}

func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

// Parse the configuration file 'config.toml', and establish a connection to DB
func init() {
	config.Read()
	dao.Server = config.Server
	dao.Database = config.Database
	dao.Connect()
}

// Define HTTP request routes
func main() {
	r := mux.NewRouter()
	r.HandleFunc("/livingcosts", AllLivingcostsEndPoint).Methods("GET")
	r.HandleFunc("/livingcosts", CreateLivingcostEndPoint).Methods("POST")
	r.HandleFunc("/livingcost/{id}", UpdateLivingcostEndPoint).Methods("PUT")
	r.HandleFunc("/livingcost/{id}", DeleteLivingcostEndPoint).Methods("DELETE")
	r.HandleFunc("/livingcosts/{id}", FindLivingcostEndpoint).Methods("GET")
	if err := http.ListenAndServe(":3000", r); err != nil {
		log.Fatal(err)
	}
}
