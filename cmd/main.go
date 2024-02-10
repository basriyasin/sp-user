package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"database/sql"
	"encoding/pem"
	"log"
	"os"

	"github.com/basriyasin/sp-user/generated"
	"github.com/basriyasin/sp-user/handler"
	"github.com/basriyasin/sp-user/repository"
	"github.com/golang-jwt/jwt"

	"github.com/labstack/echo/v4"
)

const (
	// RSA Key for JWT
	RSAKeyFileName = "id_rsa"
	RSAKeySize     = 4096

	// environtment
	EnvDatabaseURL = "DATABASE_URL"
	HTTPPort       = ":1323"
)

func main() {
	e := echo.New()
	e.Debug = true

	var server generated.ServerInterface = newServer()

	generated.RegisterHandlers(e, server)
	e.Logger.Fatal(e.Start(HTTPPort))
}

// construct new server object for echo handler
func newServer() *handler.Server {
	repo := initRepository()

	rsaPrivateKey := getRSAKey()
	opts := handler.NewServerOptions{
		Repository:    repo,
		RSAPrivateKey: rsaPrivateKey,
	}
	return handler.NewServer(opts)
}

// init the repository dependencies
func initRepository() repository.RepositoryInterface {
	dbDsn := os.Getenv(EnvDatabaseURL)
	db, err := sql.Open("postgres", dbDsn)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	return repository.NewRepository(repository.NewRepositoryOptions{
		Db: db,
	})
}

// get the RSA key for JWT signature or create new on if the key corrupted or not exists
func getRSAKey() *rsa.PrivateKey {
	key, err := os.ReadFile(RSAKeyFileName)
	if err == nil {
		return parseRSAKey(key)
	}

	return generateRSAKey()
}

// parse the existing RSA key or create new one if the existing key corrupted
func parseRSAKey(key []byte) *rsa.PrivateKey {
	rsaPrivateKey, err := jwt.ParseRSAPrivateKeyFromPEM(key)
	if err != nil {
		return generateRSAKey()
	}

	return rsaPrivateKey
}

// generate a new RSA key, save it into a file, and return the generated key
func generateRSAKey() *rsa.PrivateKey {
	pKey, err := rsa.GenerateKey(rand.Reader, RSAKeySize)
	if err != nil {
		panic(err)
	}

	pcksKey := x509.MarshalPKCS1PrivateKey(pKey)
	writer, err := os.OpenFile(RSAKeyFileName, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0400)
	if err != nil {
		log.Fatalf("Failed to open %s for writing: %v", RSAKeyFileName, err)
	}
	defer writer.Close()

	err = pem.Encode(writer, &pem.Block{
		Type:  "PRIVATE KEY",
		Bytes: pcksKey,
	})
	if err != nil {
		panic(err)
	}

	return pKey
}
