package biscuit

// NOTE: the actual filter logic is handled by the GetCookies func (cookies.go)

import (
	"strings"
	"time"
)

type Filter func(c Cookie) bool

func Host(domain string) Filter {
	return func(c Cookie) bool {
		return c.Host == domain
	}
}
func HostContains(substr string) Filter {
	return func(c Cookie) bool {
		return strings.Contains(c.Host, substr)
	}
}
func HostHasPrefix(prefix string) Filter {
	return func(c Cookie) bool {
		return strings.HasPrefix(c.Host, prefix)
	}
}
func HostHasSuffix(suffix string) Filter {
	return func(c Cookie) bool {
		return strings.HasSuffix(c.Host, suffix)
	}
}

func Name(name string) Filter {
	return func(c Cookie) bool {
		return c.Name == name
	}
}

func NameContains(substr string) Filter {
	return func(c Cookie) bool {
		return strings.Contains(c.Name, substr)
	}
}

func NameHasPrefix(prefix string) Filter {
	return func(c Cookie) bool {
		return strings.HasPrefix(c.Name, prefix)
	}
}

func NameHasSuffix(suffix string) Filter {
	return func(c Cookie) bool {
		return strings.HasSuffix(c.Name, suffix)
	}
}

func Path(path string) Filter {
	return func(c Cookie) bool {
		return c.Path == path
	}
}
func PathContains(substr string) Filter {
	return func(c Cookie) bool {
		return strings.Contains(c.Path, substr)
	}
}
func PathHasPrefix(prefix string) Filter {
	return func(c Cookie) bool {
		return strings.HasPrefix(c.Path, prefix)
	}
}
func PathHasSuffix(suffix string) Filter {
	return func(c Cookie) bool {
		return strings.HasSuffix(c.Path, suffix)
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

func ValueContains(substr string) Filter {
	return func(c Cookie) bool {
		return strings.Contains(c.Value, substr)
	}
}

func ValueHasPrefix(prefix string) Filter {
	return func(c Cookie) bool {
		return strings.HasPrefix(c.Value, prefix)
	}
}

func ValueHasSuffix(suffix string) Filter {
	return func(c Cookie) bool {
		return strings.HasSuffix(c.Value, suffix)
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