package handler

import (
	"crypto/rsa"

	"github.com/basriyasin/sp-user/repository"
)

type Server struct {
	Repository    repository.RepositoryInterface
	rsaPrivateKey *rsa.PrivateKey
}

type NewServerOptions struct {
	Repository    repository.RepositoryInterface
	RSAPrivateKey *rsa.PrivateKey
}

// create new serer repository and parse the RSA key, it will panic when invalid RSA key provided
func NewServer(opts NewServerOptions) *Server {
	return &Server{
		Repository:    opts.Repository,
		rsaPrivateKey: opts.RSAPrivateKey,
	}
}
