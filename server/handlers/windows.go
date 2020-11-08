package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	data "github.com/pellison512/viewfinder/server/data/v2"
	"github.com/pellison512/viewfinder/server/helpers/v2"
)

type WindowsHandler struct {
	dataSvc data.DataSvc
}

type WindowInfoReq struct {
	WindowText string `json:"windowText"`
	Left       int    `json:"left"`
	Top        int    `json:"top"`
	Right      int    `json:"right"`
	Bottom     int    `json:"bottom"`
}

func NewWindowsHandler(dataSvc data.DataSvc) (WindowsHandler, error) {
	return WindowsHandler{
		dataSvc: dataSvc,
	}, nil
}

func (handler *WindowsHandler) POSTWindowsHandler(w http.ResponseWriter, req *http.Request) {

	var winInfoReq WindowInfoReq
	if err := json.NewDecoder(req.Body).Decode(&winInfoReq); err != nil {
		helpers.WriteErrResponse(w, err, "error unmarshalling json", http.StatusBadRequest)
	}
	handler.dataSvc.StoreWindow(data.WindowData{
		Title:  winInfoReq.WindowText,
		Left:   winInfoReq.Left,
		Top:    winInfoReq.Top,
		Bottom: winInfoReq.Bottom,
		Right:  winInfoReq.Right,
	})

	w.WriteHeader(http.StatusCreated)
}

type WindowInfoResponse struct {
	Left   int `json:"left"`
	Top    int `json:"top"`
	Right  int `json:"right"`
	Bottom int `json:"bottom"`
}

func (handler *WindowsHandler) GETWindowsHandler(w http.ResponseWriter, req *http.Request) {
	title := mux.Vars(req)["title"]

	window, err := handler.dataSvc.GetWindow(title)
	if err != nil {
		//TODO hide db error message
		helpers.WriteErrResponse(w, err, "error getting window from database", http.StatusNotFound)
		return
	}
	resp := WindowInfoResponse{
		Left:   window.Left,
		Top:    window.Top,
		Right:  window.Right,
		Bottom: window.Bottom,
	}

	helpers.WriteJSONResponse(w, resp, http.StatusOK)
	return
}
