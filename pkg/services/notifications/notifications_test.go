package notifications

import (
	"testing"

	"github.com/logdisplayplatform/logdisplayplatform/pkg/bus"
	m "github.com/logdisplayplatform/logdisplayplatform/pkg/models"
	"github.com/logdisplayplatform/logdisplayplatform/pkg/setting"
	. "github.com/smartystreets/goconvey/convey"
)

type testTriggeredAlert struct {
	ActualValue float64
	Name        string
	State       string
}

func TestNotifications(t *testing.T) {

	Convey("Given the notifications service", t, func() {
		setting.StaticRootPath = "../../../public/"

		ns := &NotificationService{}
		ns.Bus = bus.New()
		ns.Cfg = setting.NewCfg()
		ns.Cfg.Smtp.Enabled = true
		ns.Cfg.Smtp.TemplatesPattern = "emails/*.html"
		ns.Cfg.Smtp.FromAddress = "from@address.com"
		ns.Cfg.Smtp.FromName = "LogDisplayPlatform Admin"

		err := ns.Init()
		So(err, ShouldBeNil)

		Convey("When sending reset email password", func() {
			err := ns.sendResetPasswordEmail(&m.SendResetPasswordEmailCommand{User: &m.User{Email: "asd@asd.com"}})
			So(err, ShouldBeNil)

			sentMsg := <-ns.mailQueue
			So(sentMsg.Body, ShouldContainSubstring, "body")
			So(sentMsg.Subject, ShouldEqual, "Reset your LogDisplayPlatform password - asd@asd.com")
			So(sentMsg.Body, ShouldNotContainSubstring, "Subject")
		})
	})
}
