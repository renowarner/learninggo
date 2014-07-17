package main

import (
	"errors"
	"testing"
	"time"

	"github.com/tebeka/selenium"
)

var (
	wd  selenium.WebDriver
	err error

	mail     = "http://mail.google.com/mail/h/"
	userName = "lithotesting2@gmail.com"
	userPass = "Lithosphere1"
)

func TestCreatingWebDriver(t *testing.T) {
	caps := selenium.Capabilities{"browserName": "firefox"}
	wd, err = selenium.NewRemote(caps, "")
	if err != nil {
		t.Fatal("Unable to intialize wd.  ", err)
	}
}

func TestLoadingLoginPage(t *testing.T) {
	err = wd.Get(mail)
}

func TestEmailLogin(t *testing.T) {
	err = findAndFillByCss("#Email", "Email field; ", userName)
	if err != nil {
		t.Error(err)

	}
	err = findAndFillByCss("#Passwd", "Password field", userPass)
	if err != nil {
		t.Error(err)
	}
	err = findAndClickButton("#signIn", "Sign in Button; ")
	if err != nil {
		t.Error(err)
	}
	time.Sleep(time.Millisecond * 1000)
}

func TestOpenMail(t *testing.T) {
	email, err := wd.FindElement(selenium.ByPartialLinkText, "Test Subject")
	if err != nil {
		t.Error(err)
	}
	err = email.Click()
	if err != nil {
		t.Error(err)
	}
}

/*
func TestOpenLink(t *testing.T) {
	email, err := wd.FindElement(selenium.ByLinkText, "http://www.speedtest.net/")
	if err != nil {
		t.Error(err)
	}
	err = email.Click()
	if err != nil {
		t.Error(err)
	}
}
*/
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func findAndFillByCss(fieldID, errID, fillValue string) error {
	elem, err := wd.FindElement(selenium.ByCSSSelector, fieldID)
	if err != nil {
		return errors.New("Unable to find field: " + errID + " " + err.Error())
	}

	err = elem.Clear()
	if err != nil {
		return errors.New("Unable to clear field: " + errID + " " + err.Error())
	}

	err = elem.SendKeys(fillValue)
	if err != nil {
		return errors.New("Unable to send keys to field: " + errID + " " + err.Error())
	}
	return nil
}

func findAndClickButton(fieldID, errID string) error {
	elem, err := wd.FindElement(selenium.ByCSSSelector, fieldID)
	if err != nil {
		return errors.New("Unable to find button: " + errID + " " + err.Error())
	}
	err = elem.Click()
	if err != nil {
		return errors.New("Unable to click button: " + errID + " " + err.Error())
	}
	return nil
}
