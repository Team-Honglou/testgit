package models

import (
	"testing"

	"github.com/logdisplayplatform/logdisplayplatform/pkg/components/simplejson"
	"github.com/logdisplayplatform/logdisplayplatform/pkg/setting"
	. "github.com/smartystreets/goconvey/convey"
)

func TestDashboardModel(t *testing.T) {

	Convey("Generate full dashboard url", t, func() {
		setting.AppUrl = "http://logdisplayplatform.local/"
		fullUrl := GetFullDashboardUrl("uid", "my-dashboard")
		So(fullUrl, ShouldEqual, "http://logdisplayplatform.local/d/uid/my-dashboard")
	})

	Convey("Generate relative dashboard url", t, func() {
		setting.AppUrl = ""
		fullUrl := GetDashboardUrl("uid", "my-dashboard")
		So(fullUrl, ShouldEqual, "/d/uid/my-dashboard")
	})

	Convey("When generating slug", t, func() {
		dashboard := NewDashboard("LogDisplayPlatform Play Home")
		dashboard.UpdateSlug()

		So(dashboard.Slug, ShouldEqual, "logdisplayplatform-play-home")
	})

	Convey("Can slugify title", t, func() {
		slug := SlugifyTitle("LogDisplayPlatform Play Home")

		So(slug, ShouldEqual, "logdisplayplatform-play-home")
	})

	Convey("Given a dashboard json", t, func() {
		json := simplejson.New()
		json.Set("title", "test dash")

		Convey("With tags as string value", func() {
			json.Set("tags", "")
			dash := NewDashboardFromJson(json)

			So(len(dash.GetTags()), ShouldEqual, 0)
		})
	})

	Convey("Given a new dashboard folder", t, func() {
		json := simplejson.New()
		json.Set("title", "test dash")

		cmd := &SaveDashboardCommand{Dashboard: json, IsFolder: true}
		dash := cmd.GetDashboardModel()

		Convey("Should set IsFolder to true", func() {
			So(dash.IsFolder, ShouldBeTrue)
		})
	})

	Convey("Given a child dashboard", t, func() {
		json := simplejson.New()
		json.Set("title", "test dash")

		cmd := &SaveDashboardCommand{Dashboard: json, FolderId: 1}
		dash := cmd.GetDashboardModel()

		Convey("Should set FolderId", func() {
			So(dash.FolderId, ShouldEqual, 1)
		})
	})
}