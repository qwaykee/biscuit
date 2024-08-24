package biscuit

import (
	"path/filepath"
	"os"
)

type firefoxBrowser struct {
	firefoxBase
}

func newFirefox() (Browser, error) {
	return &firefoxBrowser{
		firefoxBase: firefoxBase{
			ProfilesPathFn: firefoxProfilesPath,
		},
	}, nil
}

func (browser *firefoxBrowser) Name() string {
	return "Firefox"
}

func firefoxProfilesPath() string {
	switch detectOS() {
	case windows:
		return filepath.Join(os.Getenv("APPDATA"), "Mozilla", "Firefox", "Profiles")
	case mac:
		return filepath.Join(os.Getenv("HOME"), "Library", "Application Support", "Firefox", "Profiles")
	case linux:
		return filepath.Join(os.Getenv("HOME"), ".mozilla", "firefox")
	default:
		return "unknown profiles path"
	}
}