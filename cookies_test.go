package biscuit

import (
	"testing"
	"log"
)

func BenchmarkGetAllCookies(b *testing.B) {
    for i := 0; i < b.N; i++ {
		cookies, err := GetCookies(All, NameContains("token", "auth"))
		if err != nil {
			log.Println(err)
		}

		_ = cookies
    }
}