package main

import (
	"errors"

	"github.com/tebeka/selenium"
)

var (
	wd  selenium.WebDriver
	err error

	nexus    = "http://www.lolnexus.com/"
	username = "orphalian"
)

func main() {
	webDriver()

	loadPage(nexus)

	findAndFillByCss("#name-search", "Player Field", username)

	findAndClickButton("#new-search-submit", "Search Button")
}

///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// Utility Functions
///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func webDriver() {
	caps := selenium.Capabilities{"browserName": "firefox"}
	wd, err = selenium.NewRemote(caps, "")
	if err != nil {
		errors.New("Unable to intialize wd.")
	}
}

func loadPage(nexUrl string) {
	err = wd.Get(nexUrl)
	if err != nil {
		errors.New("Unable to load LolNexus.")
	}
}

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
