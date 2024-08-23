package biscuit

import (
	"path/filepath"
	"os"
)

type FirefoxBrowser struct {
	FirefoxBase
}

func NewFirefox() *FirefoxBrowser {
	return &FirefoxBrowser{
		FirefoxBase: FirefoxBase{
			ProfilesPathFn: firefoxProfilesPath,
		},
	}
}

func (browser *FirefoxBrowser) Name() string {
	return "Firefox"
}

func firefoxProfilesPath() string {
	switch detectOS() {
	case windows:
		return filepath.Join(os.Getenv("APPDATA"), "Mozilla", "Firefox", "Profiles")
	case mac:
		// TODO: Replace windows placeholder
		return filepath.Join(os.Getenv("APPDATA"), "Mozilla", "Firefox", "Profiles")
	case linux:
		// TODO: Replace windows placeholder
		return filepath.Join(os.Getenv("APPDATA"), "Mozilla", "Firefox", "Profiles")
	default:
		return "unknown profiles path"
	}
}