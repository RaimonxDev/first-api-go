package auth

import (
	"crypto/rsa"
	"github.com/golang-jwt/jwt/v4"
	"os"
	"sync"
)

// Las dejamos privadas para que ningun paquete las pueda usar
var (
	signKey   *rsa.PrivateKey
	verifyKey *rsa.PublicKey
	once      sync.Once
)

// LoadFiles
// sus argumentos son de tipo string porque le enviamos solo la ruta del archivo
func LoadFiles(privateFile, publicFile string) error {
	var err error
	once.Do(func() {
		err = readAndLoadFiles(privateFile, publicFile)
	})
	return err
}

func readAndLoadFiles(privateFile, publicFile string) error {

	privateBytes, err := os.ReadFile(privateFile)
	if err != nil {
		return err
	}
	publicBytes, err := os.ReadFile(publicFile)
	if err != nil {
		return err
	}
	return parseRSA(privateBytes, publicBytes)
}

// Function parse rsa keys
func parseRSA(privateBytes, publicBytes []byte) error {
	var err error
	signKey, err = jwt.ParseRSAPrivateKeyFromPEM(privateBytes)
	if err != nil {
		return err
	}
	verifyKey, err = jwt.ParseRSAPublicKeyFromPEM(publicBytes)
	if err != nil {
		return err
	}
	return nil
}
