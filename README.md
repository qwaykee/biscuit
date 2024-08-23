# Biscuit

NOTE: This library isn't stable yet, expect many changes in the future

A simple library to read cookies from multiple browsers accross all operating systems

Currently supports:
- Chrome on Windows only
- Firefox, librewolf, waterfox on Windows only
- Electron apps

Although the list is small now, the code is made to easily add new chromium and firefox based browsers which covers most browsers

## Example usage

```golang
package main

import (
	"github.com/qwaykee/biscuit"
	"github.com/rodaine/table"
)

func main() {
	cookies, err := biscuit.GetCookies(biscuit.Chrome)
	if err != nil {
		panic(err)
	}

	tbl := table.New("Domain", "Name", "Value")

	for _, cookie := range cookies {
		tbl.AddRow(cookie.Domain, cookie.Name, maxSize(cookie.Value, 80))
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