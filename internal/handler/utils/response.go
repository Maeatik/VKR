package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	v1 "diploma/internal/entity/v1"
	"diploma/pkg/logger"
)

const (
	begin = `{"data":{"name":"Детализация ежедневного отчета №%s.zip","file":"`
	end   = `", "contentType":"zip"},"error":false,"errorText":"","additionalErrors":null}`
)

type Responder struct {
	logger *logger.Logger
}

func NewResponder(logger *logger.Logger) *Responder {
	return &Responder{logger: logger}
}

func (r *Responder) RespondWithError(w http.ResponseWriter, code int, message string, handler string) {
	r.logger.Info(message)
	errorResp := v1.Response{
		Data:             nil,
		Error:            true,
		ErrorText:        message,
		AdditionalErrors: nil,
	}
	r.RespondWithJSON(w, code, errorResp, handler)
	// respondWithJSON(w, code, map[string]string{"error": message})
}

func (r *Responder) RespondWithJSON(w http.ResponseWriter, code int, payload interface{}, handler string) {
	response, err := json.Marshal(payload)
	// fmt.Println(string(response))
	if err != nil {
		fmt.Println(err)
		r.RespondWithError(w, http.StatusInternalServerError, err.Error(), handler)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func (r *Responder) RespondWithFile(w http.ResponseWriter, code int, reportID string, file io.ReadCloser) {
	handler := "excelFile"
	defer file.Close()
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	beginning := fmt.Sprintf(begin, reportID)
	w.Write([]byte(beginning))
	_, err := io.Copy(w, file)
	if err != nil {
		fmt.Println(err)
		r.RespondWithError(w, http.StatusInternalServerError, err.Error(), handler)
	}
	w.Write([]byte(end))
}
