package helpers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func WriteErrResponse(w http.ResponseWriter, err error, errStr string, status int) {
	http.Error(w, fmt.Sprintf("%s: :%v", errStr, err), status)
}

func WriteJSONResponse(w http.ResponseWriter, resp interface{}, status int) {
	bytes, err := json.Marshal(resp)
	if err != nil {
		WriteErrResponse(w, err, "error encoding json", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(status)
	w.Write(bytes)
}
