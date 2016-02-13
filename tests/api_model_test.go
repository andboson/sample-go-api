package test

import (
	"app/controllers"
	"app/models"
	"bytes"
	"compress/gzip"
	"encoding/json"
	conf "github.com/andboson/configlog"
	"github.com/julienschmidt/httprouter"
	. "github.com/smartystreets/goconvey/convey"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetModelApi(t *testing.T) {
	AddMockDataModel("test-name")

	conf.AppConfig.Set("enable_gzip", true)
	b := bytes.NewBufferString(`{"name":"test-name"}`)
	r, _ := http.NewRequest("POST", "/api/v1/model.find_by_name", b)
	r.Header.Add("Accept-Encoding", "deflate, gzip")
	w := httptest.NewRecorder()
	var reader io.ReadCloser

	router := httprouter.New()
	router.POST("/api/v1/model.find_by_name", controllers.GetModelByName)
	router.ServeHTTP(w, r)

	var response=&models.Model{}
	var responseText string
	switch w.Header().Get("Content-Encoding") {
	case "gzip":
		reader, _ = gzip.NewReader(w.Body)
		var uncompressed []byte
		uncompressed, _ = ioutil.ReadAll(reader)
		responseText = string(uncompressed)
		defer reader.Close()
	default:
		responseText = w.Body.String()
	}

	log.Printf("\n response == ", responseText, w.Header().Get("Content-Encoding"))
	json.Unmarshal([]byte(responseText), &response)
	log.Printf("\n response mapped: %+v", response)

	Convey("Subject: Test GetByNamey \n", t, func() {
		Convey("Status Code Should Be 200", func() {
			So(w.Code, ShouldEqual, 200)
		})
		Convey("The Result Name must be equal `test-name`", func() {
			So(response.Name, ShouldEqual, "test-name")
		})

	})
}
