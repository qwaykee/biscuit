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
		Domain         string
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

func (c Cookie) ToHTTPCookie() http.Cookie {
	// TODO: Finish
	return http.Cookie{}
}