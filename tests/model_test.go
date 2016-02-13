package test

import (
	"app/models"
	"app/services"
	"github.com/andboson/carbon"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestGetModel(t *testing.T) {
	AddMockDataModel("tst")

	var model = &models.Model{}
	model = model.GetByName("tst")
	Convey("Subject: Test Find Model in DB \n", t, func() {

		Convey("The Result Name must be equal `tst`", func() {
			So(model.Name, ShouldEqual, "tst")
		})

		Convey("The Result Updated At must be equal "+carbon.Now().SubDays(2).ToDateTimeString(), func() {
			So(model.Date.Day(), ShouldEqual, carbon.Now().SubDays(2).Day())
		})

	})
}

///mock
func AddMockDataModel(name string) {
	services.DB.Where("name = ?", name).Delete(models.Model{})
	model := models.Model{
		Name: name,
		Date: carbon.Now().SubDays(2).Time,
	}
	services.DB.Create(&model)
}
