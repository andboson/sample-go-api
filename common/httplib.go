package common

import (
	"compress/gzip"
	"encoding/json"
	conf "github.com/andboson/configlog"
	"net"
	"net/http"
	"reflect"
	"strconv"
	"strings"
)

//request headers
type HttpLib struct {
	XCallingService string `json:"x_calling_service"`
	XCallingMethod  string `json:"x_calling_method"`
}

func (h *HttpLib) GetRequestHeaders(request *http.Request) *HttpLib {
	response := &HttpLib{
		XCallingService: request.Header.Get("X-Calling-Service"),
		XCallingMethod:  request.Header.Get("X-Calling-Method"),
	}

	return response
}

//gzip/plain writer
func (h *HttpLib) ResponseWriter(request *http.Request, w http.ResponseWriter, response interface{}) {
	var err error
	var encoded []byte
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	enableGzip, _ := conf.AppConfig.Bool("enable_gzip")

	if reflect.TypeOf(response).String() != "[]uint8" {
		encoded, err = json.Marshal(response)
		if err == nil {
			encodedString := strings.Replace(string(encoded), `""`, `null`, -1)
			encoded = []byte(encodedString)
		}
	} else {
		encoded = response.([]byte)
	}
	if enableGzip && strings.Contains(request.Header.Get("Accept-Encoding"), "gzip") {
		w.Header().Set("Content-Encoding", "gzip")
		gz := gzip.NewWriter(w)
		defer gz.Close()
		gz.Write(encoded)
	} else {
		w.Write(encoded)
	}
	if err != nil {
		Log.WithError(err).Printf("Error encoding response, response: %+v", response)
	}
}

var Port = conf.AppConfig.UString("port")

//detect free port
func DetectPortIsFree() bool {
	var result bool
	intPort, _ := strconv.Atoi(Port)
	tcpAddr := &net.TCPAddr{}
	tcpAddr.Port = intPort
	l, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		result = false
	} else {
		result = true
		l.Close()
	}

	return result
}
