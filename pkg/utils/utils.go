package utils

import (
	"io"
	"net/http"
	"encoding/json"
)

func ParseBody(r *http.Request, x interface{}){
	body,err:=io.ReadAll(r.Body)
	if err==nil{
		err:=json.Unmarshal([]byte(body),x)
		if err!=nil{
			return 
		}
	}
}