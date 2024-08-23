package biscuit

import (
	"strings"
	"errors"
)

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
	Electron BrowserName = "electron"
)

func NewBrowser(browserName BrowserName) (Browser, error) {
	switch browserName {
	case Chrome:
		return NewChrome()
	case Firefox:
		return NewFirefox(), nil
	case Librewolf:
		return NewLibrewolf(), nil
	case Waterfox:
		return NewWaterfox(), nil
	case Electron:
		// TODO: Fix placeholder discord to allow any electron apps
		return NewElectron("discord")
	default:
		return nil, errors.New("selected browser isn't implemented yet")
	}
}

func cutAnyPrefix(s string, prefixes ...string) string {
	for _, p := range prefixes {
		s, _ = strings.CutPrefix(s, p)
	}

	return s
}