package handler

import (
	"crypto/rsa"
	"testing"

	"github.com/basriyasin/sp-user/repository"
	"github.com/golang-jwt/jwt"
	"github.com/golang/mock/gomock"
)

// get the dummy rsa private key for testing purposes
// this function should not be called in real flow
func getDummyRSAKey() *rsa.PrivateKey {
	key := []byte("-----BEGIN RSA PRIVATE KEY-----\n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\n-----END RSA PRIVATE KEY-----")
	rsaPrivateKey, err := jwt.ParseRSAPrivateKeyFromPEM(key)
	if err != nil {
		panic(err)
	}

	return rsaPrivateKey
}

func TestNewServer(t *testing.T) {
	var (
		ctrl     = gomock.NewController(t)
		mockRepo = repository.NewMockRepositoryInterface(ctrl)
	)

	test := []struct {
		name      string
		args      NewServerOptions
		expectErr bool
	}{
		{
			name:      "invalid rsa key",
			args:      NewServerOptions{mockRepo, nil},
			expectErr: false,
		},
		{
			name: "success",
			args: NewServerOptions{mockRepo, nil},
		},
	}

	for _, tt := range test {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				if tt.expectErr {
					if r := recover(); r == nil {
						t.Errorf("The code did not panic")
					}
				}
			}()

			NewServer(tt.args)
		})
	}

}
