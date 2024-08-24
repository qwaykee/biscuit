package biscuit

import (
	"path/filepath"
	"strings"
	"os"
)

type firefoxGenericBrowser struct {
	name string
	firefoxBase
}

func newFirefoxGeneric(name string) func() (Browser, error) {
	return func() (Browser, error) {
		// name is case sensitive for the path to work
		return &firefoxGenericBrowser{
			name: name,
			firefoxBase: firefoxBase{
				ProfilesPathFn: firefoxGenericProfilePath(name),
			},
		}, nil
	}
}

func (browser *firefoxGenericBrowser) Name() string {
	return strings.ToTitle(browser.name)
}

func firefoxGenericProfilePath(name string) func() string {
	return func() string {
		switch detectOS() {
		case windows:
			return filepath.Join(os.Getenv("APPDATA"), name, "Profiles")
		case mac:
			return filepath.Join(os.Getenv("HOME"), "Library", "Application Support", name, "Profiles")
		case linux:
			// TODO: Add support for flatpak (~/.var)
			return filepath.Join(os.Getenv("HOME"), "." + strings.ToLower(name))
		default:
			return "unknown profiles path"
		}
	}
}