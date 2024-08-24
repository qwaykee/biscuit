package biscuit

import "errors"

type Browser interface {
	Name() string
	GetCookies() ([]Cookie, error)
}

type BrowserName string

var (
	// candy code to use like biscuit.GetCookies(biscuit.Firefox)
	// instead of having to initialize the Browser first
	All BrowserName = "all" // try getting cookies from every browsers (see GetCookies -> cookies.go)

	Chrome BrowserName = "chrome"
	Electron BrowserName = "electron"
	Opera BrowserName = "opera"
	Edge BrowserName = "edge"
	Firefox BrowserName = "firefox"
	Librewolf BrowserName = "librewolf"
	Waterfox BrowserName = "waterfox"
	Zen BrowserName = "zen"
)

type BrowserFactory func() (Browser, error)

var browserRegistry = make(map[BrowserName]BrowserFactory)

func RegisterBrowser(name BrowserName, factory BrowserFactory) {
    browserRegistry[name] = factory
}

func NewBrowser(name BrowserName) (Browser, error) {
    if browser, exists := browserRegistry[name]; exists {
        return browser()
    }
    return nil, errors.New("selected browser isn't implemented yet")
}

func init() {
	RegisterBrowser(Chrome, NewChrome)
	RegisterBrowser(Opera, NewOpera)
	RegisterBrowser(Edge, NewEdge)
	RegisterBrowser(Electron, NewElectron("discord"))
	RegisterBrowser(Firefox, NewFirefox)
	RegisterBrowser(Librewolf, NewFirefoxGeneric("librewolf"))
	RegisterBrowser(Waterfox, NewFirefoxGeneric("Waterfox"))
	RegisterBrowser(Zen, NewFirefoxGeneric("Zen"))
}