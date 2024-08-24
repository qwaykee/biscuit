# TODO

- [ ] Chrome (chromium), firefox, safari support
- [ ] Edge, opera, epiphany, arc (chromium), firefox and derivates, internet explorer support
- [x] !!Edge support!! (can't find the cookie location)
- [ ] Windows, mac os, linux support
- [x] Cookie filters when reading
- [ ] Test for linux and mac os
- [ ] Generic chromium and firefox based browsers support
- [x] Switch to BrowsersFactory
- [ ] Add behavior settings: 

```golang
type Behavior struct {
	// if the database can't be opened because
	// it's already used by another process, try killing it
	TryKillingProcess bool


	// if the database can't be opened, retry after x seconds
	// let nil if you don't want to retry
	AutoReconnect time.Duration
}
```

https://www.cyberark.com/resources/threat-research-blog/the-current-state-of-browser-cookies

# File structure

- browsers.go: Implements Browser interface and NewBrowsers() variables
- cookies.go: Implements Cookie type and main function (GetCookies, which handles cookies filtering)
- filters.go: Implements all the default filters
- platform_detection.go: Simple function to detect the current running OS and normalize it to constants

- browser_base.go: Base for all "browser"-based browsers (examples: chromium_base.go, firefox_base.go)
- browser_browsername: Completion of needed informations for "browser"-based browsers  (examples: chromium_chrome.go, firefox_firefox.go)
- browser_generic: Generic completion of needed informations (see [browsers.go](browsers.go))
- electron_appname.go: Wrapper around chromium_electron.go or custom GetCookies implementation if needed

# Path handling

Using HOME environment variable avoids having to handle an error
but there is a small chance that HOME is not populated

# Chromium

## Add a chromium-based browser

Since chromium-based browsers use the same cookie saving logic, the functions are the same.

Chromium saves the first profile created under the "Default" directory,
Any other profile will be saved under "Profile *i*" with i as 1, 2, 3...
Unlike firefox, the Windows cookies path doesn't follow the same pattern as Linux and Mac, because of this,
the cookies files retrieval logic must return two values for each chromium-based browser (see [chromeCookiePath()](chromium_chrome.go)).

The AES encryption key remains the same for any profile as long as it's on the same session.

To add a browser based on chromium, simply copy chromium_chrome.go
and replace the localStatePath and cookiePath funcs by the corresponding path,
same for the Chrome struct, the Name and NewChrome function,
save as chromium_browsername.go
Test and if it works on all platform, commit and push

## Cookies retrieval logic

### On Windows

0. Close the browser: the database can't be accessed if it's already used by chromium

1. Decrypt cookies encrypted_value: try decrypting with CryptUnprotectData (> Chromium version 80), if doesn't work, try decrypting using AES256-GCM with the AES key

2. Decrypt AES key: found in %LOCALAPPDATA%\Google\Chrome\User Data\Local State -> json file (os_crypt > encrypted_key), decode base64, remove "DPAPI" prefix and handle to CryptUnprotectData process (Crypt32.dll)

3. Fetch cookies: found in %LOCALAPPDATA%\Google\Chrome\User Data\Default\Network\Cookies, table "cookies", have to decrypt "encrypted_value" column

### On MacOS

- Need password to keychain (TODO)

### On Linux

- Need password to keychain (TODO)


# Firefox

## Add a firefox-based browser

Since firefox-based browsers use the same cookie saving logic, the functions are the same.

Firefox uses random IDs for profiles by default (examples: fixz6yfd.default, 157azy95.default-release, 584wjw45.default-default)
when installed on the computer, one of two directory must contain the cookies.sqlite file ; if downloaded as the portable version,
the default and only directory won't contain a random ID and will be named "Default" instead and will contain the cookies.sqlite file.

firefox_base.go scans all the directory in "Profiles" in order to find cookies files (the scan is non-recursive)
This way, only the Profiles path is needed instead of the cookies path and most firefox-based browsers should work

To add a browser based on firefox, simply copy firefox_firefox.go
replace the needed functions and informations, save as firefox_browsername.go
Test and if it works on all platform, commit and push

## Cookies retrieval logic

- No need to close the browser
- Fetch cookies: found in %APPDATA%/Mozilla/Firefox/Profiles/(randomid).default-release/cookies.sqlite, table "moz_cookies", column "value" is already decrypted