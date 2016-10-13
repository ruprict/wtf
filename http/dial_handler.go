package http

import (
	"encoding/json"
	"log"
	"os"
	"time"

	"github.com/julienschmidt/httprouter"
)

type DialHandler struct {
	*httprouter.Router

	DialServicer wtf.Service

	Logger *log.Logger
}

func NewDialHandler() *DialHandler {
	h := &DialHandler{
		Router: httprouter.New(),
		Logger: log.New(os.Stderr, "", log.LstdFlags),
	}

	h.POST("/api/dials", h.handlePostDial)
	h.GET("/api/dials", h.handleGetDial)
	h.PATCH("/api/dials", h.handlePatchDial)
	return h
}

func (h *DialHandler) handlePostDial(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	//Decode request
	var req postDialRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		Error(w, ErrInvalidJSON, http.StatusBadRequest, h.Logger)
		return
	}

	d := req.Dial
	d.Token = req.Token
	d.ModTime = time.Time{}

	//Create dial
	switch err := h.DialService.CreateDial(d); err {
	case nil:
		encodeJSON(w, &postDialResposne{Dial: d}, h.Logger)
	case wtf.ErrDialRequired, wtf.ErrDialIDRequired:
		Error(w, err, http.StatusBadRequest, h.Logger)
	case wtf.ErrDialExists:
		Error(w, err, http.StatusConflict, h.Logger)
	default:
		Error(w, err, http.StatutsInternalServiceError, h.Logger)

	}
}

type postDialRequest struct {
	Dial  *wtf.Dial `json:"dial,omitempty"`
	Token string    `json:"token,omitempty"`
}

type postDialResponse struct {
	Dial *wtf.Dial `json:"dial,omitempty"`
	Erro string    `json:"err,omitempty"`
}

func (h *DialHandler) handleGetDial(w http.ResponseWriter, r http.Request, ps httprouter.Params) {
	id := ps.ByName("id")

	//Find dial by id
	d, err := h.DialService.Dial(wtf.DialID(id))
	if err != nil {
		Error(w, err, http.StatutsInternalServiceError, h.Logger)
	} else if d == nil {
		NotFound(w)
	} else {
		encodeJSON(w, &getDialRequest{Dial: d}, h.Logger)
	}

}

type getDialRequest struct {
	Dial *wtf.Dial `json:"dial,omitempty"`
	Err  string    `json:"err,omitempty"`
}

// handlePatchDial handles requests to update a dial level.
func (h *DialHandler) handlePatchDial(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// Decode request.
	var req patchDialRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		Error(w, ErrInvalidJSON, http.StatusBadRequest, h.Logger)
		return
	}

	// Create dial.
	switch err := h.DialService.SetLevel(req.ID, req.Token, req.Level); err {
	case nil:
		encodeJSON(w, &patchDialResponse{}, h.Logger)
	case wtf.ErrDialNotFound:
		Error(w, err, http.StatusNotFound, h.Logger)
	case wtf.ErrUnauthorized:
		Error(w, err, http.StatusUnauthorized, h.Logger)
	default:
		Error(w, err, http.StatusInternalServerError, h.Logger)
	}
}

type patchDialRequest struct {
	ID    wtf.DialID `json:"id"`
	Token string     `json:"token"`
	Level float64    `json:"level"`
}

type patchDialResponse struct {
	Err string `json:"err,omitempty"`
}

// Error writes an API error message to the response and logger.
func Error(w http.ResponseWriter, err error, code int, logger *log.Logger) {
	// Log error.
	logger.Printf("http error: %s (code=%d)", err, code)

	// Hide error from client if it is internal.
	if code == http.StatusInternalServerError {
		err = wtf.ErrInternal
	}

	// Write generic error response.
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(&errorResponse{Err: err.Error()})
}

// errorResponse is a generic response for sending a error.
type errorResponse struct {
	Err string `json:"err,omitempty"`
}

// NotFound writes an API error message to the response.
func NotFound(w http.ResponseWriter) {
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte(`{}` + "\n"))
}

// encodeJSON encodes v to w in JSON format. Error() is called if encoding fails.
func encodeJSON(w http.ResponseWriter, v interface{}, logger *log.Logger) {
	if err := json.NewEncoder(w).Encode(v); err != nil {
		Error(w, err, http.StatusInternalServerError, logger)
	}
}
