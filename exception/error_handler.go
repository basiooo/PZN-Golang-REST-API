package exception

import (
	"golang-restfull-api/helpers"
	"golang-restfull-api/model/web"
	"net/http"

	"github.com/go-playground/validator"
)

func ErrorHandler(w http.ResponseWriter, r *http.Request, error interface{}) {
	if notFoundError(w, r, error) {
		return
	}
	if validationErrors(w, r, error) {
		return
	}

	internalServerError(w, r, error)
}

func internalServerError(w http.ResponseWriter, r *http.Request, error interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)
	webResponse := web.WebResponse{
		Code:   http.StatusInternalServerError,
		Status: "Internal Server Error",
		Data:   error,
	}
	helpers.WriteToResponseBody(w, webResponse)
}
func validationErrors(writer http.ResponseWriter, request *http.Request, err interface{}) bool {
	exception, ok := err.(validator.ValidationErrors)
	if ok {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusBadRequest)

		webResponse := web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD REQUEST",
			Data:   exception.Error(),
		}

		helpers.WriteToResponseBody(writer, webResponse)
		return true
	} else {
		return false
	}
}
func notFoundError(w http.ResponseWriter, r *http.Request, error interface{}) bool {
	exception, ok := error.(NotFoundError)
	if ok {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		webResponse := web.WebResponse{
			Code:   http.StatusNotFound,
			Status: "NOT FOUND",
			Data:   exception.Error,
		}
		helpers.WriteToResponseBody(w, webResponse)
		return true
	} else {
		return false
	}
}
