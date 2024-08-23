package biscuit

import (
	"time"
	"net/http"
)

type (
	Cookie struct {
		CreationDate   time.Time
		ExpirationDate time.Time
		LastAccessDate time.Time
		LastUpdateDate time.Time
		Host           string
		HttpOnly       bool
		Name           string
		Value          string
		EncryptedValue []byte
		Path           string
		SameSite       string
		Secure         bool
	}
)

func GetCookies(browserName BrowserName, filters ...Filter) ([]Cookie, error) {
	browser, err := NewBrowser(browserName)
	if err != nil {
		return nil, err
	}

	cookies, err := browser.GetCookies()
	if err != nil {
		return nil, err
	}

	if len(filters) > 0 {
		filteredCookies := make([]Cookie, 0, len(cookies))

		for _, cookie := range cookies {
			keep := true

			for _, filter := range filters {
				if !filter(cookie) {
					keep = false
					break
				}
			}

			if keep {
				filteredCookies = append(filteredCookies, cookie)
			}
		}

		return filteredCookies, nil
	}
	
	return cookies, nil
}

func (c Cookie) ToHTTPCookie() *http.Cookie {
	var sameSite http.SameSite
	switch c.SameSite {
	case "Strict":
		sameSite = http.SameSiteStrictMode
	case "Lax":
		sameSite = http.SameSiteLaxMode
	case "None":
		sameSite = http.SameSiteNoneMode
	default:
		sameSite = http.SameSiteDefaultMode
	}

	return &http.Cookie{
		Name: c.Name,
		Value: c.Value,
		Path: c.Path,
		Expires: c.ExpirationDate,
		// RawExpires string    // for reading cookies only
		// MaxAge      int
		Secure: c.Secure,
		HttpOnly: c.HttpOnly,
		SameSite: sameSite,
		// Partitioned bool
		// Raw         string
		// Unparsed:    []string // Raw text of unparsed attribute-value pairs
	}
}