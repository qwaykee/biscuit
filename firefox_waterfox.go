package biscuit

import (
	"path/filepath"
	"os"
)

type WaterfoxBrowser struct {
	FirefoxBase
}

func NewWaterfox() *WaterfoxBrowser {
	return &WaterfoxBrowser{
		FirefoxBase: FirefoxBase{
			ProfilesPathFn: waterfoxProfilesPath,
		},
	}
}

func (browser *WaterfoxBrowser) Name() string {
	return "Firefox"
}

func waterfoxProfilesPath() string {
	switch detectOS() {
	case windows:
		return filepath.Join(os.Getenv("APPDATA"), "Waterfox", "Profiles")
	case mac:
		// TODO: Replace windows placeholder
		return filepath.Join(os.Getenv("APPDATA"), "Waterfox", "Profiles")
	case linux:
		// TODO: Replace windows placeholder
		return filepath.Join(os.Getenv("APPDATA"), "Waterfox", "Profiles")
	default:
		return "unknown profiles path"
	}
}