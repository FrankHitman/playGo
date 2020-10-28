package main

import (
	"crypto/sha256"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	h := sha256.New()
	h.Write([]byte("password"))
	shPassword := fmt.Sprintf("%x\n", h.Sum(nil))

	fmt.Println(shPassword)
	// sum := sha256.Sum256([]byte("password"))
	// fmt.Printf("%x\n", sum)

	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(shPassword), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println("generate from password err ", err)
		return
	}
	fmt.Println(string(hashedBytes))

	if err := bcrypt.CompareHashAndPassword(hashedBytes, []byte(shPassword)); err != nil {
		fmt.Printf("password not correct %s", err)
		return
	}
	fmt.Println("success")

}

// $2a$10$GjdmB5kSbGmce/CYkxNvu.kT2gU1.kGf/5Ws4JEp6MHFB88Npgag2
// 5e884898da28047151d0e56f8dc6292773603d0d6aabbdd62a11ef721d1542d8
//
// $2a$10$6Fg1Xx4QSdAxlMdEWdZZh.5evMcHbLj6pKpIKj2eCtus7Brc9nLvq
// success
