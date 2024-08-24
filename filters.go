package biscuit

// NOTE: the actual filter logic is handled by the GetCookies func (cookies.go)

import (
	"strings"
	"time"
)

func executeOnAny(fn func(string, string) bool, compareTo string, list []string) bool {
	for _, str := range list {
		if fn(compareTo, str) {
			return true
		}
	}
	return false
}

type Filter func(c Cookie) bool

func Host(domain string) Filter {
	return func(c Cookie) bool {
		return c.Host == domain
	}
}
func HostContains(substr ...string) Filter {
	return func(c Cookie) bool {
		return executeOnAny(strings.Contains, c.Host, substr)
	}
}
func HostHasPrefix(prefixes ...string) Filter {
	return func(c Cookie) bool {
		return executeOnAny(strings.HasPrefix, c.Host, prefixes)
	}
}
func HostHasSuffix(suffixes ...string) Filter {
	return func(c Cookie) bool {
		return executeOnAny(strings.HasSuffix, c.Host, suffixes)
	}
}

func Name(name string) Filter {
	return func(c Cookie) bool {
		return c.Name == name
	}
}

func NameContains(substr ...string) Filter {
	return func(c Cookie) bool {
		return executeOnAny(strings.Contains, c.Name, substr)
	}
}

func NameHasPrefix(prefixes ...string) Filter {
	return func(c Cookie) bool {
		return executeOnAny(strings.HasPrefix, c.Name, prefixes)
	}
}

func NameHasSuffix(suffixes ...string) Filter {
	return func(c Cookie) bool {
		return executeOnAny(strings.HasSuffix, c.Name, suffixes)
	}
}

func Path(path string) Filter {
	return func(c Cookie) bool {
		return c.Path == path
	}
}
func PathContains(substr ...string) Filter {
	return func(c Cookie) bool {
		return executeOnAny(strings.Contains, c.Path, substr)
	}
}
func PathHasPrefix(prefixes ...string) Filter {
	return func(c Cookie) bool {
		return executeOnAny(strings.HasPrefix, c.Path, prefixes)
	}
}
func PathHasSuffix(suffixes ...string) Filter {
	return func(c Cookie) bool {
		return executeOnAny(strings.HasSuffix, c.Path, suffixes)
	}
}
func PathDepth(depth int) Filter {
	return func(c Cookie) bool {
		return strings.Count(strings.TrimRight(c.Path, `/`), `/`) == depth
	}
}

func ExpiresAfter(time time.Time) Filter {
	return func(c Cookie) bool {
		return c.ExpirationDate.After(time)
	}
}
func ExpiresBefore(time time.Time) Filter {
	return func(c Cookie) bool {
		return c.ExpirationDate.Before(time)
	}
}

func Value(value string) Filter {
	return func(c Cookie) bool {
		return c.Value == value
	}
}

func ValueContains(substr ...string) Filter {
	return func(c Cookie) bool {
		return executeOnAny(strings.Contains, c.Value, substr)
	}
}

func ValueHasPrefix(prefixes ...string) Filter {
	return func(c Cookie) bool {
		return executeOnAny(strings.HasPrefix, c.Value, prefixes)
	}
}

func ValueHasSuffix(suffixes ...string) Filter {
	return func(c Cookie) bool {
		return executeOnAny(strings.HasSuffix, c.Value, suffixes)
	}
}

func ValueLen(length int) Filter {
	return func(c Cookie) bool {
		return len(c.Value) == length
	}
}

func CreatedAfter(time time.Time) Filter {
	return func(c Cookie) bool {
		return c.CreationDate.After(time)
	}
}
func CreatedBefore(time time.Time) Filter {
	return func(c Cookie) bool {
		return c.CreationDate.Before(time)
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

var Expired Filter = func(c Cookie) bool {
	return time.Now().After(c.ExpirationDate)
}