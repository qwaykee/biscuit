package biscuit

// NOTE: the actual filter logic is handled by the GetCookies func (cookies.go)

import (
	"strings"
	"time"
)

type (
	Filter func(c Cookie) bool
)

// func (c Cookies) Filter(filters ...Filter) []Cookie {
// 	var newCookies []Cookie
// 
// 	for _, filter := range filters {
// 		for _, cookie := range c {
// 			if filter(cookie) {
// 				newCookies = append(newCookies, cookie)
// 			}
// 		}
// 	}
// 
// 	return newCookies
// }

func Name(name string) Filter {
	return func(c Cookie) bool {
		return c.Name == name
	}
}

func NameContains(name string) Filter {
	return func(c Cookie) bool {
		return strings.Contains(c.Name, name)
	}
}

var Secure Filter = func(c Cookie) bool {
	return c.Secure
}

var HTTPOnly Filter = func(c Cookie) bool {
	return c.HttpOnly
}

var Valid Filter = func(c Cookie) bool {
	return time.Now().Before(c.ExpirationDate)
}