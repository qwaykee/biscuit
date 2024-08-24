package biscuit

import (
	"path/filepath"
	"os"
)

type EdgeBrowser struct {
	ChromiumBase
}

func NewEdge() (Browser, error) {
	chrome := &EdgeBrowser{
		ChromiumBase: ChromiumBase{
			LocalStatePathFn: edgeLocalStatePath,
			CookiePathFn: edgeCookiePath,
		},
	}

	if err := chrome.initializeChromium(); err != nil {
		return nil, err
	}

	return chrome, nil
}

func (browser *EdgeBrowser) Name() string {
	return "Edge"
}

func edgeLocalStatePath() string {
	switch detectOS() {
	case windows:
		return filepath.Join(os.Getenv("LOCALAPPDATA"), "Microsoft", "Edge", "User Data", "Local State")
	case mac:
		return filepath.Join(os.Getenv("HOME"), "Library", "Application Support", "Microsoft", "Edge", "User Data", "Local State")
	case linux:
		// TODO: Replace google-chrome
		return filepath.Join(os.Getenv("HOME"), ".config", "google-chrome", "User Data", "Local State")
	default:
		return "unknown local state path"
	}
}
	
func edgeCookiePath() (string, string) {
	var userDataPath string
	var cookieScanPath string

	switch detectOS() {
	case windows:
		userDataPath = filepath.Join(os.Getenv("LOCALAPPDATA"), "Microsoft", "Edge", "User Data")
		cookieScanPath = filepath.Join("Network", "Cookies")
	case mac:
		userDataPath = filepath.Join(os.Getenv("HOME"), "Library", "Application Support", "Microsoft", "Edge", "User Data")
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