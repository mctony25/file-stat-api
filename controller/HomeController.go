package controller

import (
	"file-stat/app"
	"file-stat/response"
	"net/http"
	"time"
)

type HomeController struct{}

func (fc HomeController) GetInfo(w http.ResponseWriter, r *http.Request) {

	homeInfo := response.HomeInfo{
		AppVersion: app.Version,
		Timestamp:  time.Now(),
	}

	(response.JsonResponse{}).Send(w, homeInfo, http.StatusOK)

	return
}
