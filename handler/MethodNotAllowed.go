package handler

import (
	"file-stat/response"
	"net/http"
)

func MethodNotAllowed(w http.ResponseWriter, r *http.Request) {
	methodNotAllowed := response.Error{
		ErrorCode: http.StatusMethodNotAllowed,
		Message: "Requested method is not allowed",
	}

	(response.JsonResponse{}).Send(w, methodNotAllowed, http.StatusMethodNotAllowed)

	return
}

func MethodNotAllowedHandler() http.Handler { return http.HandlerFunc(MethodNotAllowed) }


