package library

import (
	"encoding/json"
	"net/http"
)

type Responsee struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

//set error message in Error struct
func SetResponse(rsp Responsee, message string, emptyArr interface{}) Responsee {
	rsp.Message = message
	rsp.Data = emptyArr
	return rsp
}

func Res_401(w http.ResponseWriter, msg string) {
	var err Responsee
	err = SetResponse(err, msg, []string{})
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(401)
	json.NewEncoder(w).Encode(err)
	return
}

func Res_400(w http.ResponseWriter, msg string) {
	var err Responsee
	err = SetResponse(err, msg, []string{})
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(400)
	json.NewEncoder(w).Encode(err)
	return
}

func Res_500(w http.ResponseWriter, msg string) {
	var err Responsee
	err = SetResponse(err, msg, []string{})
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(500)
	json.NewEncoder(w).Encode(err)
	return
}

func Res_Unknown(w http.ResponseWriter, msg string) {
	var err Responsee
	err = SetResponse(err, msg, []string{})
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(err)
	return
}

func Res_200(w http.ResponseWriter, msg string, data interface{}) {
	var rsp Responsee
	rsp = SetResponse(rsp, msg, data)
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(rsp)
	return
}

func Res_204(w http.ResponseWriter, msg string) {
	var err Responsee
	err = SetResponse(err, msg, []string{})
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(204)
	json.NewEncoder(w).Encode(err)
	return
}

func PingHandler(w http.ResponseWriter, r *http.Request) {
	var res Responsee
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(200)
	// res.Status = 200
	res.Message = "Server Active"
	res.Data = []int{}
	json.NewEncoder(w).Encode(res)
}
