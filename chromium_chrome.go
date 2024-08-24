package biscuit

import (
	"path/filepath"
	"os"
)

type ChromeBrowser struct {
	ChromiumBase
}

func NewChrome() (Browser, error) {
	chrome := &ChromeBrowser{
		ChromiumBase: ChromiumBase{
			LocalStatePathFn: chromeLocalStatePath,
			CookiePathFn: chromeCookiePath,
		},
	}

	if err := chrome.initializeChromium(); err != nil {
		return nil, err
	}

	return chrome, nil
}

func (browser *ChromeBrowser) Name() string {
	return "Chrome"
}

func chromeLocalStatePath() string {
	switch detectOS() {
	case windows:
		return filepath.Join(os.Getenv("LOCALAPPDATA"), "Google", "Chrome", "User Data", "Local State")
	case mac:
		return filepath.Join(os.Getenv("HOME"), "Library", "Application Support", "Google", "Chrome", "Local State")
	case linux:
		return filepath.Join(os.Getenv("HOME"), ".config", "google-chrome", "Local State")
	default:
		return "unknown local state path"
	}
}
	
func chromeCookiePath() (string, string) {
	var userDataPath string
	var cookieScanPath string

	switch detectOS() {
	case windows:
		userDataPath = filepath.Join(os.Getenv("LOCALAPPDATA"), "Google", "Chrome", "User Data")
		cookieScanPath = filepath.Join("Network", "Cookies")
	case mac:
		userDataPath = filepath.Join(os.Getenv("HOME"), "Library", "Application Support", "Google", "Chrome")
		cookieScanPath = "Cookies"
	case linux:
		userDataPath = filepath.Join(os.Getenv("HOME"), ".config", "google-chrome")
		cookieScanPath = "Cookies"
	default:
		return "", ""
	}

	return userDataPath, cookieScanPath
}