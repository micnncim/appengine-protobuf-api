package response

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gogo/protobuf/proto"
	"google.golang.org/appengine/log"
)

type Result struct {
	Message string `json:"message,omitempty"`
}

// Created sends body with created status (201)
func Created(w http.ResponseWriter, r *http.Request, m interface{}) {
	sendBody(w, r, m, http.StatusCreated)
}

// Accepted sends body with accepted status (202)
func Accepted(w http.ResponseWriter, r *http.Request, m interface{}) {
	sendBody(w, r, m, http.StatusAccepted)
}

// NoContent responses no content status (204)
func NoContent(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNoContent)
}

// BadRequest sends error (400)
func BadRequest(w http.ResponseWriter, r *http.Request, err error) {
	sendError(w, r, err, http.StatusBadRequest)
}

// Unauthorized sends error (401)
func Unauthorized(w http.ResponseWriter, r *http.Request, err error) {
	sendError(w, r, err, http.StatusUnauthorized)
}

// Forbidden sends error (403)
func Forbidden(w http.ResponseWriter, r *http.Request, err error) {
	sendError(w, r, err, http.StatusForbidden)
}

// NotFound sends error (404)
func NotFound(w http.ResponseWriter, r *http.Request, err error) {
	sendError(w, r, err, http.StatusNotFound)
}

// Unknown sends error (500)
func InternalServerError(w http.ResponseWriter, r *http.Request, err error) {
	sendError(w, r, err, http.StatusInternalServerError)
}

// NotImplemented sends error (501)
func NotImplemented(w http.ResponseWriter, r *http.Request, err error) {
	sendError(w, r, err, http.StatusNotImplemented)
}

// Unavailable sends error (503)
func Unavailable(w http.ResponseWriter, r *http.Request, err error) {
	sendError(w, r, err, http.StatusServiceUnavailable)
}

func sendBody(w http.ResponseWriter, r *http.Request, msg interface{}, status int) {
	data, err := json.Marshal(msg)
	if err != nil {
		BadRequest(w, r, err)
		return
	}
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(status)
	w.Write(data)
}

func Send(w http.ResponseWriter, r *http.Request, code int, body proto.Marshaler) {
	var contentType string
	var b []byte

	switch r.Header.Get("Accept") {
	case "application/protobuf", "application/x-protobuf":
		contentType = "application/protobuf"
		var err error
		b, err = body.Marshal()
		if err != nil {
			code = http.StatusInternalServerError
		}
	default:
		contentType = "application/json"
		var err error
		b, err = json.Marshal(body)
		if err != nil {
			code = http.StatusInternalServerError
		}
	}

	w.Header().Set("Content-Type", contentType)
	w.Header().Set("Content-Length", strconv.Itoa(len(b)))
	w.WriteHeader(code)
	w.Write(b)
}

func Succeed(w http.ResponseWriter, r *http.Request, body proto.Marshaler) {
	Send(w, r, http.StatusOK, body)
}

func sendError(w http.ResponseWriter, r *http.Request, err error, status int) {
	log.Errorf(r.Context(), "%s", err.Error())
	sendBody(w, r, Result{Message: err.Error()}, status)
}
