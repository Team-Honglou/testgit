package login

import (
	"testing"

	"github.com/logdisplayplatform/logdisplayplatform/pkg/bus"
	m "github.com/logdisplayplatform/logdisplayplatform/pkg/models"
	. "github.com/smartystreets/goconvey/convey"
)

func TestLogDisplayPlatformLogin(t *testing.T) {
	Convey("Login using LogDisplayPlatform DB", t, func() {
		logdisplayplatformLoginScenario("When login with non-existing user", func(sc *logdisplayplatformLoginScenarioContext) {
			sc.withNonExistingUser()
			err := loginUsingLogDisplayPlatformDB(sc.loginUserQuery)

			Convey("it should result in user not found error", func() {
				So(err, ShouldEqual, m.ErrUserNotFound)
			})

			Convey("it should not call password validation", func() {
				So(sc.validatePasswordCalled, ShouldBeFalse)
			})

			Convey("it should not pupulate user object", func() {
				So(sc.loginUserQuery.User, ShouldBeNil)
			})
		})

		logdisplayplatformLoginScenario("When login with invalid credentials", func(sc *logdisplayplatformLoginScenarioContext) {
			sc.withInvalidPassword()
			err := loginUsingLogDisplayPlatformDB(sc.loginUserQuery)

			Convey("it should result in invalid credentials error", func() {
				So(err, ShouldEqual, ErrInvalidCredentials)
			})

			Convey("it should call password validation", func() {
				So(sc.validatePasswordCalled, ShouldBeTrue)
			})

			Convey("it should not pupulate user object", func() {
				So(sc.loginUserQuery.User, ShouldBeNil)
			})
		})

		logdisplayplatformLoginScenario("When login with valid credentials", func(sc *logdisplayplatformLoginScenarioContext) {
			sc.withValidCredentials()
			err := loginUsingLogDisplayPlatformDB(sc.loginUserQuery)

			Convey("it should not result in error", func() {
				So(err, ShouldBeNil)
			})

			Convey("it should call password validation", func() {
				So(sc.validatePasswordCalled, ShouldBeTrue)
			})

			Convey("it should pupulate user object", func() {
				So(sc.loginUserQuery.User, ShouldNotBeNil)
				So(sc.loginUserQuery.User.Login, ShouldEqual, sc.loginUserQuery.Username)
				So(sc.loginUserQuery.User.Password, ShouldEqual, sc.loginUserQuery.Password)
			})
		})
	})
}

type logdisplayplatformLoginScenarioContext struct {
	loginUserQuery         *m.LoginUserQuery
	validatePasswordCalled bool
}

type logdisplayplatformLoginScenarioFunc func(c *logdisplayplatformLoginScenarioContext)

func logdisplayplatformLoginScenario(desc string, fn logdisplayplatformLoginScenarioFunc) {
	Convey(desc, func() {
		origValidatePassword := validatePassword

		sc := &logdisplayplatformLoginScenarioContext{
			loginUserQuery: &m.LoginUserQuery{
				Username:  "user",
				Password:  "pwd",
				IpAddress: "192.168.1.1:56433",
			},
			validatePasswordCalled: false,
		}

		defer func() {
			validatePassword = origValidatePassword
		}()

		fn(sc)
	})
}

func mockPasswordValidation(valid bool, sc *logdisplayplatformLoginScenarioContext) {
	validatePassword = func(providedPassword string, userPassword string, userSalt string) error {
		sc.validatePasswordCalled = true

		if !valid {
			return ErrInvalidCredentials
		}

		return nil
	}
}

func (sc *logdisplayplatformLoginScenarioContext) getUserByLoginQueryReturns(user *m.User) {
	bus.AddHandler("test", func(query *m.GetUserByLoginQuery) error {
		if user == nil {
			return m.ErrUserNotFound
		}

		query.Result = user
		return nil
	})
}

func (sc *logdisplayplatformLoginScenarioContext) withValidCredentials() {
	sc.getUserByLoginQueryReturns(&m.User{
		Id:       1,
		Login:    sc.loginUserQuery.Username,
		Password: sc.loginUserQuery.Password,
		Salt:     "salt",
	})
	mockPasswordValidation(true, sc)
}

func (sc *logdisplayplatformLoginScenarioContext) withNonExistingUser() {
	sc.getUserByLoginQueryReturns(nil)
}

func (sc *logdisplayplatformLoginScenarioContext) withInvalidPassword() {
	sc.getUserByLoginQueryReturns(&m.User{
		Password: sc.loginUserQuery.Password,
		Salt:     "salt",
	})
	mockPasswordValidation(false, sc)
}
