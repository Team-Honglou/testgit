package login

import (
	"errors"
	"testing"

	m "github.com/logdisplayplatform/logdisplayplatform/pkg/models"
	. "github.com/smartystreets/goconvey/convey"
)

func TestAuthenticateUser(t *testing.T) {
	Convey("Authenticate user", t, func() {
		authScenario("When a user authenticates having too many login attempts", func(sc *authScenarioContext) {
			mockLoginAttemptValidation(ErrTooManyLoginAttempts, sc)
			mockLoginUsingLogDisplayPlatformDB(nil, sc)
			mockLoginUsingLdap(true, nil, sc)
			mockSaveInvalidLoginAttempt(sc)

			err := AuthenticateUser(sc.loginUserQuery)

			Convey("it should result in", func() {
				So(err, ShouldEqual, ErrTooManyLoginAttempts)
				So(sc.loginAttemptValidationWasCalled, ShouldBeTrue)
				So(sc.logdisplayplatformLoginWasCalled, ShouldBeFalse)
				So(sc.ldapLoginWasCalled, ShouldBeFalse)
				So(sc.saveInvalidLoginAttemptWasCalled, ShouldBeFalse)
			})
		})

		authScenario("When logdisplayplatform user authenticate with valid credentials", func(sc *authScenarioContext) {
			mockLoginAttemptValidation(nil, sc)
			mockLoginUsingLogDisplayPlatformDB(nil, sc)
			mockLoginUsingLdap(true, ErrInvalidCredentials, sc)
			mockSaveInvalidLoginAttempt(sc)

			err := AuthenticateUser(sc.loginUserQuery)

			Convey("it should result in", func() {
				So(err, ShouldEqual, nil)
				So(sc.loginAttemptValidationWasCalled, ShouldBeTrue)
				So(sc.logdisplayplatformLoginWasCalled, ShouldBeTrue)
				So(sc.ldapLoginWasCalled, ShouldBeFalse)
				So(sc.saveInvalidLoginAttemptWasCalled, ShouldBeFalse)
			})
		})

		authScenario("When logdisplayplatform user authenticate and unexpected error occurs", func(sc *authScenarioContext) {
			customErr := errors.New("custom")
			mockLoginAttemptValidation(nil, sc)
			mockLoginUsingLogDisplayPlatformDB(customErr, sc)
			mockLoginUsingLdap(true, ErrInvalidCredentials, sc)
			mockSaveInvalidLoginAttempt(sc)

			err := AuthenticateUser(sc.loginUserQuery)

			Convey("it should result in", func() {
				So(err, ShouldEqual, customErr)
				So(sc.loginAttemptValidationWasCalled, ShouldBeTrue)
				So(sc.logdisplayplatformLoginWasCalled, ShouldBeTrue)
				So(sc.ldapLoginWasCalled, ShouldBeFalse)
				So(sc.saveInvalidLoginAttemptWasCalled, ShouldBeFalse)
			})
		})

		authScenario("When a non-existing logdisplayplatform user authenticate and ldap disabled", func(sc *authScenarioContext) {
			mockLoginAttemptValidation(nil, sc)
			mockLoginUsingLogDisplayPlatformDB(m.ErrUserNotFound, sc)
			mockLoginUsingLdap(false, nil, sc)
			mockSaveInvalidLoginAttempt(sc)

			err := AuthenticateUser(sc.loginUserQuery)

			Convey("it should result in", func() {
				So(err, ShouldEqual, ErrInvalidCredentials)
				So(sc.loginAttemptValidationWasCalled, ShouldBeTrue)
				So(sc.logdisplayplatformLoginWasCalled, ShouldBeTrue)
				So(sc.ldapLoginWasCalled, ShouldBeTrue)
				So(sc.saveInvalidLoginAttemptWasCalled, ShouldBeFalse)
			})
		})

		authScenario("When a non-existing logdisplayplatform user authenticate and invalid ldap credentials", func(sc *authScenarioContext) {
			mockLoginAttemptValidation(nil, sc)
			mockLoginUsingLogDisplayPlatformDB(m.ErrUserNotFound, sc)
			mockLoginUsingLdap(true, ErrInvalidCredentials, sc)
			mockSaveInvalidLoginAttempt(sc)

			err := AuthenticateUser(sc.loginUserQuery)

			Convey("it should result in", func() {
				So(err, ShouldEqual, ErrInvalidCredentials)
				So(sc.loginAttemptValidationWasCalled, ShouldBeTrue)
				So(sc.logdisplayplatformLoginWasCalled, ShouldBeTrue)
				So(sc.ldapLoginWasCalled, ShouldBeTrue)
				So(sc.saveInvalidLoginAttemptWasCalled, ShouldBeTrue)
			})
		})

		authScenario("When a non-existing logdisplayplatform user authenticate and valid ldap credentials", func(sc *authScenarioContext) {
			mockLoginAttemptValidation(nil, sc)
			mockLoginUsingLogDisplayPlatformDB(m.ErrUserNotFound, sc)
			mockLoginUsingLdap(true, nil, sc)
			mockSaveInvalidLoginAttempt(sc)

			err := AuthenticateUser(sc.loginUserQuery)

			Convey("it should result in", func() {
				So(err, ShouldBeNil)
				So(sc.loginAttemptValidationWasCalled, ShouldBeTrue)
				So(sc.logdisplayplatformLoginWasCalled, ShouldBeTrue)
				So(sc.ldapLoginWasCalled, ShouldBeTrue)
				So(sc.saveInvalidLoginAttemptWasCalled, ShouldBeFalse)
			})
		})

		authScenario("When a non-existing logdisplayplatform user authenticate and ldap returns unexpected error", func(sc *authScenarioContext) {
			customErr := errors.New("custom")
			mockLoginAttemptValidation(nil, sc)
			mockLoginUsingLogDisplayPlatformDB(m.ErrUserNotFound, sc)
			mockLoginUsingLdap(true, customErr, sc)
			mockSaveInvalidLoginAttempt(sc)

			err := AuthenticateUser(sc.loginUserQuery)

			Convey("it should result in", func() {
				So(err, ShouldEqual, customErr)
				So(sc.loginAttemptValidationWasCalled, ShouldBeTrue)
				So(sc.logdisplayplatformLoginWasCalled, ShouldBeTrue)
				So(sc.ldapLoginWasCalled, ShouldBeTrue)
				So(sc.saveInvalidLoginAttemptWasCalled, ShouldBeFalse)
			})
		})

		authScenario("When logdisplayplatform user authenticate with invalid credentials and invalid ldap credentials", func(sc *authScenarioContext) {
			mockLoginAttemptValidation(nil, sc)
			mockLoginUsingLogDisplayPlatformDB(ErrInvalidCredentials, sc)
			mockLoginUsingLdap(true, ErrInvalidCredentials, sc)
			mockSaveInvalidLoginAttempt(sc)

			err := AuthenticateUser(sc.loginUserQuery)

			Convey("it should result in", func() {
				So(err, ShouldEqual, ErrInvalidCredentials)
				So(sc.loginAttemptValidationWasCalled, ShouldBeTrue)
				So(sc.logdisplayplatformLoginWasCalled, ShouldBeTrue)
				So(sc.ldapLoginWasCalled, ShouldBeTrue)
				So(sc.saveInvalidLoginAttemptWasCalled, ShouldBeTrue)
			})
		})
	})
}

type authScenarioContext struct {
	loginUserQuery                   *m.LoginUserQuery
	logdisplayplatformLoginWasCalled            bool
	ldapLoginWasCalled               bool
	loginAttemptValidationWasCalled  bool
	saveInvalidLoginAttemptWasCalled bool
}

type authScenarioFunc func(sc *authScenarioContext)

func mockLoginUsingLogDisplayPlatformDB(err error, sc *authScenarioContext) {
	loginUsingLogDisplayPlatformDB = func(query *m.LoginUserQuery) error {
		sc.logdisplayplatformLoginWasCalled = true
		return err
	}
}

func mockLoginUsingLdap(enabled bool, err error, sc *authScenarioContext) {
	loginUsingLdap = func(query *m.LoginUserQuery) (bool, error) {
		sc.ldapLoginWasCalled = true
		return enabled, err
	}
}

func mockLoginAttemptValidation(err error, sc *authScenarioContext) {
	validateLoginAttempts = func(username string) error {
		sc.loginAttemptValidationWasCalled = true
		return err
	}
}

func mockSaveInvalidLoginAttempt(sc *authScenarioContext) {
	saveInvalidLoginAttempt = func(query *m.LoginUserQuery) {
		sc.saveInvalidLoginAttemptWasCalled = true
	}
}

func authScenario(desc string, fn authScenarioFunc) {
	Convey(desc, func() {
		origLoginUsingLogDisplayPlatformDB := loginUsingLogDisplayPlatformDB
		origLoginUsingLdap := loginUsingLdap
		origValidateLoginAttempts := validateLoginAttempts
		origSaveInvalidLoginAttempt := saveInvalidLoginAttempt

		sc := &authScenarioContext{
			loginUserQuery: &m.LoginUserQuery{
				Username:  "user",
				Password:  "pwd",
				IpAddress: "192.168.1.1:56433",
			},
		}

		defer func() {
			loginUsingLogDisplayPlatformDB = origLoginUsingLogDisplayPlatformDB
			loginUsingLdap = origLoginUsingLdap
			validateLoginAttempts = origValidateLoginAttempts
			saveInvalidLoginAttempt = origSaveInvalidLoginAttempt
		}()

		fn(sc)
	})
}
