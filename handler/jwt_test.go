package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

const (
	dummyValidToken   = "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJjcmVhdGVkX2F0IjoiMjAyNC0wMi0xMFQwOToyMToyNy41ODQ1OTlaIiwiZXhwIjo1MzA3NTYxMDAyLCJpZCI6MSwibG9naW5fY291bnQiOjE3LCJuYW1lIjoibmFydG94bCIsInBhc3N3b3JkIjoiJDJhJDEwJC5EeHR6cGF6anBxMy5YVHNSYm1DSi5rRkMvVzFZTnhuOHNTV05WY3NBUEN0aGFraFh3OExhIiwicGhvbmUiOiIrNjI4MTEyMjMzNDQ1NSIsInVwZGF0ZWRfYXQiOnsiVGltZSI6IjAwMDEtMDEtMDFUMDA6MDA6MDBaIiwiVmFsaWQiOmZhbHNlfX0.gygYz0XREHjd6_Xv3wVCavIDo6atLNqZMcNoZEKXaCI"
	dummyExpiredToken = "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJjcmVhdGVkX2F0IjoiMjAyNC0wMi0xMFQwOToyMToyNy41ODQ1OTlaIiwiZXhwIjoxNzA3NTY0MzMwLCJpZCI6MSwibG9naW5fY291bnQiOjE1LCJuYW1lIjoibmFydG94bCIsInBhc3N3b3JkIjoiJDJhJDEwJC5EeHR6cGF6anBxMy5YVHNSYm1DSi5rRkMvVzFZTnhuOHNTV05WY3NBUEN0aGFraFh3OExhIiwicGhvbmUiOiIrNjI4MTEyMjMzNDQ1NSIsInVwZGF0ZWRfYXQiOnsiVGltZSI6IjAwMDEtMDEtMDFUMDA6MDA6MDBaIiwiVmFsaWQiOmZhbHNlfX0.gPcJkDC4oxQv0ASssgI0mAL0-PSs-eZ-zXKCXmn0B_Q"
)

func TestVerifyToken(t *testing.T) {
	var (
		e            = echo.New()
		mockValidKey = getDummyRSAKey()
	)

	test := []struct {
		name      string
		token     string
		expectErr bool
	}{
		{
			name:      "empty token",
			token:     "",
			expectErr: true,
		},
		{
			name:      "err invalid key",
			token:     "invalid",
			expectErr: true,
		},
		{
			name:      "err expired token",
			token:     dummyExpiredToken,
			expectErr: true,
		},
		{
			name:  "success",
			token: dummyValidToken,
		},
	}

	for _, tt := range test {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodPost, "/register", nil)
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			req.Header.Set(echo.HeaderAuthorization, tt.token)
			c := e.NewContext(req, httptest.NewRecorder())
			_, err := verifyToken(c, mockValidKey)
			if tt.expectErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
