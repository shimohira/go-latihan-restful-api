package exception

import (
	"latihan-restful-api/helper"
	"latihan-restful-api/model/web"
	"net/http"
)

func ErrorHandler(w http.ResponseWriter, r *http.Request, err interface{}) {
	internalServerError(w, r, err)
}

func internalServerError(w http.ResponseWriter, r *http.Request, err interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)

	webResponse := web.WebResponse{
		Code:   http.StatusInternalServerError,
		Status: "INTERNAL SERVER ERROR",
		Data:   err,
	}

	helper.WriteToResponseBody(w, webResponse)
}
