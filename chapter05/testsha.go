package main

import (
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
)

func main() {

	// sha1
	fmt.Printf("%x\n", sha1.Sum([]byte("我是Jack")))
	sha1Hasher := sha1.New()
	sha1Hasher.Write([]byte("我是Jack"))
	fmt.Printf("%x\n", sha1Hasher.Sum(nil))
	// sha256
	fmt.Printf("%x\n", sha256.Sum256([]byte("我是Jack")))
	sha256Hasher := sha256.New()
	sha256Hasher.Write([]byte("我是Jack"))
	fmt.Printf("%x\n", sha256Hasher.Sum(nil))
	// sha512
	fmt.Printf("%x\n", sha512.Sum512([]byte("我是Jack")))
	sha512Hasher := sha512.New()
	sha512Hasher.Write([]byte("我是Jack"))
	fmt.Printf("%x\n", sha512Hasher.Sum(nil))

}
