package handler

import (
	"file-stat/response"
	"net/http"
)

func NotFound(w http.ResponseWriter, r *http.Request) {
	notFound := response.Error{
		ErrorCode: http.StatusNotFound,
		Message: "Requested resource not found",
	}

	(response.JsonResponse{}).Send(w, notFound, http.StatusNotFound)

	return
}

func NotFoundHandler() http.Handler { return http.HandlerFunc(NotFound) }


