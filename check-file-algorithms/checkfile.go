package main

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"fmt"
	"hash"
	"io"
	"log"
	"os"

	"github.com/spf13/pflag"
)

func main() {
	var fileName = pflag.StringP("filename", "f", "", "name of the file to be encrypted.")
	var algorithms = pflag.StringP("algorithms", "a", "sha256", "algorithms to be used, such as sha256, md5, sha1.")
	var input = pflag.StringP("input", "i", "", "strings to be encrypted.")

	pflag.Parse()

	var h hash.Hash
	if *algorithms == "sha256" {
		h = sha256.New()

	} else if *algorithms == "md5" {
		h = md5.New()

	} else if *algorithms == "sha1" {
		h = sha1.New()

	} else {
		log.Println("error algorithms input, please retry after modify")
		os.Exit(0)
	}

	if *fileName != "" {
		f, err := os.Open(*fileName)
		if err != nil {
			log.Fatal("open file failed: ", err)
		}
		defer f.Close()

		if _, err := io.Copy(h, f); err != nil {
			log.Fatal(err)
		}
		log.Printf("%s's %s is: %x\n", *fileName, *algorithms, h.Sum(nil))
	}

	if *input != "" {
		h.Write([]byte(*input))
		fmt.Printf("%s's %s is: %x\n", *input, *algorithms, h.Sum(nil))

	}

}

// go build -o checkfile algorithms/checkfile.go
// CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -v -o checkfile-mac64 algorithms/checkfile.go
// CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -v -o checkfile-linux64 algorithms/checkfile.go
// CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -v -o checkfile.exe algorithms/checkfile.go
// cp checkfile ~/go/bin/
// checkfile -f yourfilename -a md5
