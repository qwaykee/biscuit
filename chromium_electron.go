package biscuit

import (
	"path/filepath"
	"os"
)

type electronApp struct {
	chromiumBase
	AppName string
}

func newElectron(appName string) func() (Browser, error) {
	return func() (Browser, error) {
		electron := &electronApp{
			AppName: appName,
			chromiumBase: chromiumBase{
				LocalStatePathFn: electronLocalStatePath(appName),
				CookiePathFn:     electronCookiePath(appName),
			},
		}

		if err := electron.initializeChromium(); err != nil {
			return nil, err
		}

		return electron, nil
	}
}

func (e *electronApp) Name() string {
	return e.AppName
}

func electronLocalStatePath(appName string) func() string {
	return func() string {
		switch detectOS() {
		case windows:
			return filepath.Join(os.Getenv("APPDATA"), appName, "Local State")
		case mac:
			return filepath.Join(os.Getenv("HOME"), "Library", "Application Support", appName, "Local State")
		case linux:
			return filepath.Join(os.Getenv("HOME"), ".config", appName, "Local State")
		default:
			return "unsupported OS for Electron"
		}
	}
}

func electronCookiePath(appName string) func() (string, string) {
	return func() (string, string) {
		var userDataPath string
		var cookieScanPath string

		switch detectOS() {
		case windows:
			userDataPath = filepath.Join(os.Getenv("APPDATA"), appName)
			cookieScanPath = filepath.Join("Network", "Cookies")
		case mac:
			userDataPath = filepath.Join(os.Getenv("HOME"), "Library", "Application Support", appName)
			cookieScanPath = "Cookies"
		case linux:
			userDataPath = filepath.Join(os.Getenv("HOME"), ".config", appName)
			cookieScanPath = "Cookies"
		default:
			return "", ""
		}

		return userDataPath, cookieScanPath
	}
}