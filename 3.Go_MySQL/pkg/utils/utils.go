package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func ParseBody(r *http.Request, x interface{}) {
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(x)
	if err != nil {
		fmt.Println("Error parsing body: ", err)
		return
	}
	defer r.Body.Close()
}
