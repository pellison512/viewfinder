package handlers

import (
	"encoding/json"
	"log"
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

	var winInfoReq []WindowInfoReq
	if err := json.NewDecoder(req.Body).Decode(&winInfoReq); err != nil {
		helpers.WriteErrResponse(w, err, "error unmarshalling json", http.StatusBadRequest)
	}
	log.Printf("window Req: %+v", winInfoReq)
	for _, winInfo := range winInfoReq {
		handler.dataSvc.StoreWindow(data.WindowData{
			Title:  winInfo.WindowText,
			Left:   winInfo.Left,
			Top:    winInfo.Top,
			Bottom: winInfo.Bottom,
			Right:  winInfo.Right,
		})
	}

	w.WriteHeader(http.StatusCreated)
}

type WindowInfoResponse struct {
	Title  string `json:"windowText,omitmepty"`
	Left   int    `json:"left"`
	Top    int    `json:"top"`
	Right  int    `json:"right"`
	Bottom int    `json:"bottom"`
}

func (handler *WindowsHandler) GETWindowsHandler(w http.ResponseWriter, req *http.Request) {
	title := mux.Vars(req)["title"]

	//TODO remove all this
	log.Printf("getting title: %s", title)
	if title == "React App" {

	}

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

func (handler *WindowsHandler) GETAllWindowsHandler(w http.ResponseWriter, req *http.Request) {

	windows, err := handler.dataSvc.GetAllWindows()
	if err != nil {
		//TODO hide db error message
		helpers.WriteErrResponse(w, err, "error getting window from database", http.StatusNotFound)
		return
	}
	resp := make([]WindowInfoResponse, 0)
	for _, window := range windows {
		resp = append(resp, WindowInfoResponse{
			Title:  window.Title,
			Left:   window.Left,
			Top:    window.Top,
			Right:  window.Right,
			Bottom: window.Bottom,
		})
	}

	helpers.WriteJSONResponse(w, resp, http.StatusOK)
	return
}
