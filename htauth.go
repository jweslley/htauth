package main

import (
	"crypto/sha1"
	"encoding/base64"
	"flag"
	"fmt"
	"os"
	"syscall"

	"golang.org/x/crypto/bcrypt"
	"golang.org/x/crypto/ssh/terminal"
)

var programVersion = "dev"

var (
	hash    = flag.String("h", "bcrypt", "Hashing encryption for passwords. Available hashing algorithms: bcrypt, sha1, plain.")
	realm   = flag.String("r", "", "The realm name to which the user name belongs. Used only to generate passwords for digest authentication. See http://tools.ietf.org/html/rfc2617#section-3.2.1 for more details.")
	version = flag.Bool("v", false, "Print version information and exit")
)

type Hashing struct {
	Prefix string
	Hash   func(password string) (string, error)
}

var hashings = map[string]*Hashing{
	"bcrypt": &Hashing{Hash: hashBcrypt},
	"sha1":   &Hashing{Prefix: "{SHA}", Hash: hashSha1},
	"plain":  &Hashing{Hash: hashPlain},
}

func hashPlain(password string) (string, error) {
	return password, nil
}

func hashSha1(password string) (string, error) {
	s := sha1.New()
	s.Write([]byte(password))
	sum := []byte(s.Sum(nil))
	return base64.StdEncoding.EncodeToString(sum), nil
}

func hashBcrypt(password string) (string, error) {
	sum, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(sum), nil
}

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: htauth [OPTION]... username\n")
		fmt.Fprintf(os.Stderr, "Generate encrypted passwords for basic and digest authentication.\n\n")
		flag.PrintDefaults()
	}
	flag.Parse()

	if *version {
		fmt.Printf("%s %s\n", os.Args[0], programVersion)
		return
	}

	if len(flag.Args()) == 0 {
		flag.Usage()
		os.Exit(1)
	}

	var username = flag.Args()[0]

	hashing, ok := hashings[*hash]
	if !ok {
		fmt.Fprintf(os.Stderr, "Unknown hashing algorithm: %s\n", *hash)
		os.Exit(1)
	}

	fmt.Fprintf(os.Stderr, "Enter password for %s: ", username)
	password, err := terminal.ReadPassword(int(syscall.Stdin))
	if err != nil {
		panic(err)
	}
	fmt.Fprintf(os.Stderr, "\n")

	hashedPassword, err := hashing.Hash(string(password))
	if err != nil {
		panic(err)
	}

	if *realm == "" {
		fmt.Printf("%s:%s%s\n", username, hashing.Prefix, hashedPassword)
	} else {
		fmt.Printf("%s:%s:%s%s\n", username, *realm, hashing.Prefix, hashedPassword)
	}
}
