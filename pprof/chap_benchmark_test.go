package main

import (
	"testing"

	"golang.org/x/crypto/bcrypt"
)

func BenchmarkEqual(b *testing.B) {
	b.StopTimer()
	passwd := []byte("somepasswordyoulike")
	hash, _ := bcrypt.GenerateFromPassword(passwd, 10)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		bcrypt.CompareHashAndPassword(hash, passwd)
	}
}
