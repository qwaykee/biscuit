package biscuit

import "errors"

type Browser interface {
	Name() string
	GetCookies() ([]Cookie, error)
}

type BrowserName string

var (
	// candy code to use like biscuit.GetCookies(biscuit.Firefox)
	// instead of having to initialize the Browser first
	Chrome BrowserName = "chrome"
	Firefox BrowserName = "firefox"
	Librewolf BrowserName = "librewolf"
	Waterfox BrowserName = "waterfox"
	Zen BrowserName = "zen"
	Electron BrowserName = "electron"
)

func NewBrowser(browserName BrowserName) (Browser, error) {
	switch browserName {
	case Chrome:
		return NewChrome()
	case Firefox:
		return NewFirefox(), nil
	case Librewolf:
		return NewFirefoxGeneric("librewolf"), nil
	case Waterfox:
		return NewFirefoxGeneric("Waterfox"), nil
	case Zen:
		return NewFirefoxGeneric("Zen"), nil
	case Electron:
		// TODO: Fix placeholder discord to allow any electron apps
		return NewElectron("discord")
	default:
		return nil, errors.New("selected browser isn't implemented yet")
	}
}