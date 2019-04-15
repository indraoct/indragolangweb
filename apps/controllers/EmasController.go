package controllers

import (
	"encoding/json"
	"net/http"
)

type PayResponseSuccess struct {
	ResponseCode    string      `json:"response_code"`
	ResponseDesc    struct{
		Id      string      `json:"id"`
		En      string      `json:"en"`
	}      `json:"response_desc"`
}
func ResponseDummy(w http.ResponseWriter, req *http.Request){
	w.Header().Set("Content-Type", "application/json")
	var response PayResponseSuccess
	
	response.ResponseCode = "00"
	response.ResponseDesc.Id = "Sukses"
	response.ResponseDesc.En = "Success"
	
	 data,_:= json.Marshal(response)
	 w.Write(data)
	return

}