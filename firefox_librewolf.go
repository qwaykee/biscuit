package biscuit

import (
	"path/filepath"
	"os"
)

type LibrewolfBrowser struct {
	FirefoxBase
}

func NewLibrewolf() *LibrewolfBrowser {
	return &LibrewolfBrowser{
		FirefoxBase: FirefoxBase{
			ProfilesPathFn: librewolfProfilesPath,
		},
	}
}

func (browser *LibrewolfBrowser) Name() string {
	return "Firefox"
}

func librewolfProfilesPath() string {
	switch detectOS() {
	case windows:
		return filepath.Join(os.Getenv("APPDATA"), "librewolf", "Profiles")
	case mac:
		// TODO: Replace windows placeholder
		return filepath.Join(os.Getenv("APPDATA"), "librewolf", "Profiles")
	case linux:
		// TODO: Replace windows placeholder
		return filepath.Join(os.Getenv("APPDATA"), "librewolf", "Profiles")
	default:
		return "unknown profiles path"
	}
}