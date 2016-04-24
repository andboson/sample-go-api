package controllers

import (
	. "app/common"
	"app/models"
	"encoding/json"
	conf "github.com/andboson/configlog"
	"github.com/julienschmidt/httprouter"
	"io/ioutil"
	"net/http"
	"time"
)

type InputRequest struct {
	Name string `json:"name"`
}

type ModelController struct {
	Request  *http.Request
	Response http.ResponseWriter
	HttpLib
}

// http handler GetImages
func GetModelByName(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	controller := ModelController{Request: r, Response: w}
	controller.getModelByName()
}

// get model by name
func (d *ModelController) getModelByName() {
	var input InputRequest
	var start time.Time
	var model = &models.Model{}

	debug, _ := conf.AppConfig.Bool("debug")
	InitLogger(d.GetRequestHeaders(d.Request))
	defer LogErr()
	body, error := ioutil.ReadAll(d.Request.Body)
	if debug {
		start = time.Now()
	}
	error = json.Unmarshal(body, &input)
	if debug {
		Log.WithField("request", string(body)).Printf("model.find_by_name")
	}
	if error != nil {
		Log.WithField("error", error).Printf("read body or json error")
	}
	if input.Name == "" {
		d.makeIncorrectRequestParamsReponse("Name param is required")
		return
	}

	model = model.GetByName(input.Name)

	if debug {
		times := time.Since(start)
		Log.WithField("model.find_by_name_timing", times).Printf("human %s", times)
	}

	if model.Name != "" {
		d.ResponseWriter(d.Request, d.Response, model)
	} else {
		d.makeNotFound("model.not_found", "Model not found")
		return
	}

	return
}

/////////

// not found
func (d *ModelController) makeNotFound(method, message string) {
	d.Response.WriteHeader(http.StatusNotFound)
	d.Response.Write([]byte(`{"model.` + method + `": "` + message + `"}`))
}

// incorrect params
func (d *ModelController) makeIncorrectRequestParamsReponse(error string) {
	d.Response.WriteHeader(http.StatusNotAcceptable)
	d.Response.Write([]byte(error))
}
