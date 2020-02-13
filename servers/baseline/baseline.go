package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"

	"github.com/sirupsen/logrus"
)

type EchoRequest struct {
	Message string `json:"message"`
}

type EchoReply struct {
	Message string `json:"message"`
}

type ErrorReply struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type Response struct {
	Data  interface{} `json:"data"`
	Error *ErrorReply `json:"error"`
}

func Respond(res http.ResponseWriter, code int, data interface{}, err error) {
	response := Response{}
	if err != nil {
		response.Error = &ErrorReply{Code: code, Message: err.Error()}
	} else {
		response.Data = data
	}
	res.WriteHeader(code)
	if err := json.NewEncoder(res).Encode(response); err != nil {
		logrus.Errorf("Error encoding response: %s", err.Error())
	}
}

func EchoHandler(res http.ResponseWriter, req *http.Request) {
	// logrus.WithField("request-uri", req.RequestURI).Info("Echo()")
	echoReq := EchoRequest{}
	if req.Method != http.MethodPost {
		Respond(res, http.StatusMethodNotAllowed, nil, fmt.Errorf("Only POST Supported"))
		return
	}
	defer req.Body.Close()
	if err := json.NewDecoder(req.Body).Decode(&echoReq); err != nil {
		Respond(res, http.StatusInternalServerError, nil, fmt.Errorf("Could not decode request body: %w", err))
		return
	}
	if echoReq.Message == "user-error" {
		Respond(res, http.StatusBadRequest, nil, errors.New("The User Made an Error"))
		return
	}
	Respond(res, http.StatusOK, EchoReply{Message: echoReq.Message}, nil)
}

func main() {
	listen := os.Getenv("LISTEN")
	if listen == "" {
		listen = ":8081"
	}
	http.HandleFunc("/v1/echo", EchoHandler)
	logrus.WithField("listen", listen).Info("REST service listening")
	http.ListenAndServe(listen, nil)
}
