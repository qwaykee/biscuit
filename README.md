# Biscuit

NOTE: This library isn't stable yet, expect many changes in the future

A simple library to read cookies from multiple browsers accross all operating systems with a Kooky-like syntax

Currently supports:
- Chrome. edge, opera on Windows only
- Firefox, librewolf, waterfox, zen on Windows only
- Electron apps

Although the list is small now, the code is made to easily add new chromium and firefox based browsers which covers most browsers as well as others browser

### Why this project in the first place?

Trying to retrieve cookies for a personal project, I found that most libraries didn't work for me or didn't meet my requirements.
This library aims to be internally as simple as possible and to have a great developer experience.

## Example usage

```golang
package main

import (
	"github.com/qwaykee/biscuit"
	"github.com/rodaine/table"
)

func main() {
	// get all cookies which name contains "token" or "auth"
	// from any chrome profile at default installation path
	cookies, err := biscuit.GetCookies(biscuit.Chrome, biscuit.NameContains("token", "auth"))
	if err != nil {
		panic(err)
	}

	tbl := table.New("Host", "Name", "Value")

	for _, cookie := range cookies {
		tbl.AddRow(cookie.Host, cookie.Name, maxSize(cookie.Value, 80))
	}

	tbl.Print()
}

func maxSize(text string, size int) string {
	if len(text) < size {
		return text
	} else {
		return text[:size]
	}
}
```

## More informations

More informations are written in [notes.md](notes.md), which also provides instructions on how to decrypt chromium cookies
And how to add chromium and firefox based browsers
