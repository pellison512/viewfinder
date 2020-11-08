package handlers

import (
	"encoding/json"
	"net/http"

	data "github.com/pellison512/viewfinder/server/data/v2"
	"github.com/pellison512/viewfinder/server/helpers/v2"
)

type WindowsHandler struct {
	dataSvc data.Data
}

type WindowInfoReq struct {
	WindowText string `json:"windowText"`
	Left       int    `json:"left"`
	Top        int    `json:"top"`
	Right      int    `json:"right"`
	Bottom     int    `json:"bottom"`
}

func(handler *WindowsHandler) PostWindowsHandler(w http.ResponseWriter, req *http.Request) {

	var winInfoReq WindowInfoReq
	if err := json.NewDecoder(req.Body).Decode(&winInfoReq); err != nil {
		helpers.WriteErrResponse(w, err, "error unmarshalling json", http.StatusBadRequest)
	}
	//handler.dataSvc.
}
