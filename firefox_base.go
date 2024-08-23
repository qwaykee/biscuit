package biscuit

import (
	"fmt"
	"log"
	"time"
	"os"
	"path/filepath"
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

type FirefoxBase struct {
	ProfilesPathFn     func() string
}

func NewFirefoxBase() *FirefoxBase {
	return &FirefoxBase{}
}

func (fb *FirefoxBase) Name() string {
	return "Name"
}

func (fb *FirefoxBase) GetCookies() ([]Cookie, error) {
	var cookiePaths []string

	profilesPath := fb.ProfilesPathFn()

	profiles, err := os.ReadDir(profilesPath)
	if err != nil {
		return nil, err
	}

	for _, profile := range profiles {
		path := filepath.Join(profilesPath, profile.Name(), "cookies.sqlite")

		if _, err := os.Stat(path); !os.IsNotExist(err) {
			cookiePaths = append(cookiePaths, path)
		}
	}

	var cookies []Cookie

	for _, cookiePath := range cookiePaths {
		newCookies, err := firefoxGetCookies(cookiePath)
		if err != nil {
			return cookies, err
		}

		cookies = append(cookies, newCookies...)
	}

	return cookies, nil
}

func firefoxGetCookies(cookiePath string) ([]Cookie, error) {
	db, err := sql.Open("sqlite3", fmt.Sprintf("file:%s?cache=shared&mode=ro", cookiePath))
	if err != nil {
		return nil, err
	}
	defer db.Close()

	rows, err := db.Query(`
		SELECT 
		creationTime,
		expiry,
		lastAccessed,
		host,
		isHttpOnly,
		name,
		value,
		path,
		samesite,
		isSecure
		FROM moz_cookies
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var cookies []Cookie

	for rows.Next() {
		var cookie Cookie
		var creationDate int64
		var expirationDate int64
		var lastAccessDate int64
		
		if err := rows.Scan(
			&creationDate,
			&expirationDate,
			&lastAccessDate,
			&cookie.Host,
			&cookie.HttpOnly,
			&cookie.Name,
			&cookie.Value,
			&cookie.Path,
			&cookie.SameSite,
			&cookie.Secure,
		); err != nil {
			log.Println(err)
		}

		cookie.CreationDate = time.Unix(creationDate, 0)
		cookie.ExpirationDate = time.Unix(expirationDate, 0)
		cookie.LastAccessDate = time.Unix(lastAccessDate, 0)

		cookies = append(cookies, cookie)
	}

	return cookies, nil
}