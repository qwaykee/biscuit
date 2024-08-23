package biscuit

import (
	"fmt"
	"log"
	"path/filepath"
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"syscall"
	"unsafe"
	"encoding/json"
	"encoding/base64"
	"encoding/hex"
	"crypto/aes"
	"crypto/cipher"
	"time"
	"os"
)

type ChromiumBase struct {
	EncryptedAESKey  []byte
	DecryptedAESKey  string
	Version          int
	LocalStatePathFn func() string
	CookiePathFn     func() (string, string)
	CookiesPath		 []string
}

func (cb *ChromiumBase) initializeChromium() error {
	data, err := os.ReadFile(cb.LocalStatePathFn())
	if err != nil {
		return err
	}

	temp := struct {
		OsCrypt struct {
			EncryptedKey string `json:"encrypted_key"`
		} `json:"os_crypt"`
		Browser struct {
			Version int `json:"last_whats_new_version"`
		} `json:"browser"`
	}{}

	if err := json.Unmarshal(data, &temp); err != nil {
		return err
	}

	cb.EncryptedAESKey = []byte(temp.OsCrypt.EncryptedKey)
	cb.Version = temp.Browser.Version

	aesKey, err := cb.getAesKey()
	if err != nil {
		return err
	}

	cb.DecryptedAESKey = aesKey

	userDataPath, cookieScanPath := cb.CookiePathFn()

	userData, err := os.ReadDir(userDataPath)
	if err != nil {
		return err
	}

	for _, profile := range userData {
		path := filepath.Join(userDataPath, profile.Name(), cookieScanPath)

		if _, err := os.Stat(path); !os.IsNotExist(err) {
			cb.CookiesPath = append(cb.CookiesPath, path)
		}
	}

	return nil
}

func (cb *ChromiumBase) GetCookies() ([]Cookie, error) {
	var cookies []Cookie

	for _, cookiePath := range cb.CookiesPath {
		newCookies, err := cb.getCookiesFrom(cookiePath)
		if err != nil {
			return cookies, err
		}

		cookies = append(cookies, newCookies...)
	}

	return cookies, nil
}

func (cb *ChromiumBase) getCookiesFrom(cookiePath string) ([]Cookie, error) {
	db, err := sql.Open("sqlite3", fmt.Sprintf("file:%s?cache=shared&mode=ro", cookiePath))
	if err != nil {
		return nil, err
	}
	defer db.Close()

	rows, err := db.Query(`
		SELECT 
		creation_utc,
		expires_utc,
		last_access_utc,
		last_update_utc,
		host_key,
		is_httponly,
		name,
		value,
		encrypted_value,
		path,
		samesite,
		is_secure
		FROM cookies
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
		var lastUpdateDate int64
		
		if err := rows.Scan(
			&creationDate,
			&expirationDate,
			&lastAccessDate,
			&lastUpdateDate,
			&cookie.Domain,
			&cookie.HttpOnly,
			&cookie.Name,
			&cookie.Value,
			&cookie.EncryptedValue,
			&cookie.Path,
			&cookie.SameSite,
			&cookie.Secure,
		); err != nil {
			log.Println(err)
		}

		cookie.CreationDate = time.Unix(creationDate, 0)
		cookie.ExpirationDate = time.Unix(expirationDate, 0)
		cookie.LastAccessDate = time.Unix(lastAccessDate, 0)
		cookie.LastUpdateDate = time.Unix(lastUpdateDate, 0)

		if cookie.Value != "" || len(cookie.EncryptedValue) == 0 {
			cookies = append(cookies, cookie)
			// https://codereview.chromium.org/24734007
			// cookie value is already decrypted (Chromium version < 33)
			continue
		}

		// TODO: Decrypt for every OS
		switch detectOS() {
		case windows:
			if cb.Version != 0 && cb.Version >= 80 {
				if decryptedValue, err := decryptWithAES(cb.DecryptedAESKey, cookie.EncryptedValue); err != nil {
					log.Println(err)
					// return []Cookie{}, err
				} else {
					cookie.Value = decryptedValue
				}
			} else {
				if decryptedValue, err := decryptWithDPAPI(cookie.EncryptedValue); err != nil {
					log.Println(err)
					// return []Cookie{}, err
				} else {
					cookie.Value = string(decryptedValue)
				}
			}
		}

		cookies = append(cookies, cookie)
	}

	return cookies, nil
}

func (cb *ChromiumBase) getAesKey() (string, error) {
    encryptedAesKey, err := base64.StdEncoding.DecodeString(string(cb.EncryptedAESKey))
	if err != nil {
		return "", err
	}

    aesKey, err := decryptWithDPAPI(encryptedAesKey)
	if err != nil {
		return "", err
	}

    return hex.EncodeToString(aesKey), nil
}

type dataBlob struct {
	cbData uint32
	pbData *byte
}

var (
	crypt32            = syscall.NewLazyDLL("crypt32.dll")
	cryptUnprotectData = crypt32.NewProc("CryptUnprotectData")
)

func decryptWithDPAPI(encrypted []byte) ([]byte, error) {
	if len(encrypted) < 1 {
		return nil, fmt.Errorf("encrypted data is too short")
	}

	encryptedWithoutPrefix := cutAnyPrefix(string(encrypted), "v10", "v11", "DPAPI")
	encrypted = []byte(encryptedWithoutPrefix)

	var outBlob dataBlob

	inBlob := dataBlob {
		cbData: uint32(len(encrypted)),
		pbData: &encrypted[0],
	}

	ret, _, err := cryptUnprotectData.Call(uintptr(unsafe.Pointer(&inBlob)), 0, 0, 0, 0, 0, uintptr(unsafe.Pointer(&outBlob)))
	if ret == 0 { // ret == 0 means failure in WinAPI
		return nil, fmt.Errorf("decryption failed: %v", err)
	}

	decrypted := make([]byte, outBlob.cbData)
	copy(decrypted, (*[1 << 30]byte)(unsafe.Pointer(outBlob.pbData))[:outBlob.cbData:outBlob.cbData])
	return decrypted, nil
}

func decryptWithAES(keyString string, encrypted []byte) (string, error) {
	// https://gist.github.com/donvito/efb2c643b724cf6ff453da84985281f8
	key, err := hex.DecodeString(keyString)
	if err != nil {
		return "", err
	}

	enc := cutAnyPrefix(string(encrypted), "v10", "v11")
	encBytes := []byte(enc)

	//Create a new Cipher Block from the key
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	//Create a new GCM
	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	//Get the nonce size
	nonceSize := aesGCM.NonceSize()

	//Extract the nonce from the encrypted data
	nonce, ciphertext := encBytes[:nonceSize], encBytes[nonceSize:]

	//Decrypt the data
	plaintext, err := aesGCM.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return "", err
	}

	return string(plaintext), nil
}