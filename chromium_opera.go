package biscuit

import (
	"path/filepath"
	"os"
)

type OperaBrowser struct {
	ChromiumBase
}

func NewOpera() (Browser, error) {
	chrome := &OperaBrowser{
		ChromiumBase: ChromiumBase{
			LocalStatePathFn: operaLocalStatePath,
			CookiePathFn: operaCookiePath,
		},
	}

	if err := chrome.initializeChromium(); err != nil {
		return nil, err
	}

	return chrome, nil
}

func (browser *OperaBrowser) Name() string {
	return "Opera"
}

func operaLocalStatePath() string {
	switch detectOS() {
	case windows:
		return filepath.Join(os.Getenv("APPDATA"), "Opera Software", "Opera Stable", "Local State")
	case mac:
		return filepath.Join(os.Getenv("HOME"), "Library", "Application Support", "Opera Software", "Opera Stable", "Local State")
	case linux:
		// TODO: Replace google-chrome
		return filepath.Join(os.Getenv("HOME"), ".config", "google-chrome", "Local State")
	default:
		return "unknown local state path"
	}
}
	
func operaCookiePath() (string, string) {
	var userDataPath string
	var cookieScanPath string

	switch detectOS() {
	case windows:
		userDataPath = filepath.Join(os.Getenv("APPDATA"), "Opera Software", "Opera Stable")
		cookieScanPath = filepath.Join("Network", "Cookies")
	case mac:
		userDataPath = filepath.Join(os.Getenv("HOME"), "Library", "Application Support", "Opera Software", "Opera Stable")
		cookieScanPath = "Cookies"
	case linux:
		// TODO: Replace google-chrome
		userDataPath = filepath.Join(os.Getenv("HOME"), ".config", "google-chrome")
		cookieScanPath = "Cookies"
	default:
		return "", ""
	}

	return userDataPath, cookieScanPath
}