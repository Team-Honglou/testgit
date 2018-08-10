package notifiers

import (
	"testing"

	"github.com/logdisplayplatform/logdisplayplatform/pkg/components/simplejson"
	m "github.com/logdisplayplatform/logdisplayplatform/pkg/models"
	. "github.com/smartystreets/goconvey/convey"
)

func TestEmailNotifier(t *testing.T) {
	Convey("Email notifier tests", t, func() {

		Convey("Parsing alert notification from settings", func() {
			Convey("empty settings should return error", func() {
				json := `{ }`

				settingsJSON, _ := simplejson.NewJson([]byte(json))
				model := &m.AlertNotification{
					Name:     "ops",
					Type:     "email",
					Settings: settingsJSON,
				}

				_, err := NewEmailNotifier(model)
				So(err, ShouldNotBeNil)
			})

			Convey("from settings", func() {
				json := `
				{
					"addresses": "ops@logdisplayplatform.org"
				}`

				settingsJSON, _ := simplejson.NewJson([]byte(json))
				model := &m.AlertNotification{
					Name:     "ops",
					Type:     "email",
					Settings: settingsJSON,
				}

				not, err := NewEmailNotifier(model)
				emailNotifier := not.(*EmailNotifier)

				So(err, ShouldBeNil)
				So(emailNotifier.Name, ShouldEqual, "ops")
				So(emailNotifier.Type, ShouldEqual, "email")
				So(emailNotifier.Addresses[0], ShouldEqual, "ops@logdisplayplatform.org")
			})

			Convey("from settings with two emails", func() {
				json := `
				{
					"addresses": "ops@logdisplayplatform.org;dev@logdisplayplatform.org"
				}`

				settingsJSON, err := simplejson.NewJson([]byte(json))
				So(err, ShouldBeNil)

				model := &m.AlertNotification{
					Name:     "ops",
					Type:     "email",
					Settings: settingsJSON,
				}

				not, err := NewEmailNotifier(model)
				emailNotifier := not.(*EmailNotifier)

				So(err, ShouldBeNil)
				So(emailNotifier.Name, ShouldEqual, "ops")
				So(emailNotifier.Type, ShouldEqual, "email")
				So(len(emailNotifier.Addresses), ShouldEqual, 2)

				So(emailNotifier.Addresses[0], ShouldEqual, "ops@logdisplayplatform.org")
				So(emailNotifier.Addresses[1], ShouldEqual, "dev@logdisplayplatform.org")
			})
		})
	})
}