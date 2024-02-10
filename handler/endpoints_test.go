package handler

import (
	"database/sql"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/basriyasin/sp-user/repository"
	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestRegister(t *testing.T) {
	var (
		// dependencies mock
		ctrl     = gomock.NewController(t)
		mockRepo = repository.NewMockRepositoryInterface(ctrl)

		// request and response mock
		any     = gomock.Any()
		mockErr = errors.New("an error")
		mockReq = `{"name": "narto", "phone": "+6281122334455", "password": "Aa123!@#"}`

		// echo server mock
		server = NewServer(NewServerOptions{mockRepo, getDummyRSAKey()})
		e      = echo.New()
	)

	test := []struct {
		name      string
		req       string
		mock      func()
		expectErr bool
	}{
		{
			name:      "err bind request",
			req:       "asd",
			expectErr: true,
		},
		{
			name:      "err invalid request",
			req:       `{"name": "a", "phone":"+62", "passsword":"x123"}`,
			expectErr: true,
		},
		{
			name:      "err save profile",
			req:       mockReq,
			expectErr: true,
			mock: func() {
				mockRepo.EXPECT().SaveProfile(any, any).Return(int64(0), mockErr)
			},
		},
		{
			name:      "success",
			req:       mockReq,
			expectErr: false,
			mock: func() {
				mockRepo.EXPECT().SaveProfile(any, any).Return(int64(1), nil)
			},
		},
	}

	for _, tt := range test {
		t.Run(tt.name, func(t *testing.T) {
			if tt.mock != nil {
				tt.mock()
			}

			req := httptest.NewRequest(http.MethodPost, "/register", strings.NewReader(tt.req))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			c := e.NewContext(req, httptest.NewRecorder())
			err := server.Register(c)

			if tt.expectErr {
				assert.Error(t, err)
			}
		})
	}
}

func TestAuthenticate(t *testing.T) {
	var (
		// dependencies mock
		ctrl     = gomock.NewController(t)
		mockRepo = repository.NewMockRepositoryInterface(ctrl)

		// request and response mock
		any      = gomock.Any()
		mockErr  = errors.New("an error")
		mockReq  = `{"phone": "+6281122334455", "password": "Aa123!@#"}`
		mockUser = repository.User{Password: "$2a$10$OAN2DNAF79q/njdQAlnYF./iKq.5XYq/txWdlJnA1czE4IiAnZ86K", UpdatedAt: sql.NullTime{Valid: true}}

		// echo server mock
		e       = echo.New()
		reqPath = "/authenticate"
		server  = NewServer(NewServerOptions{mockRepo, getDummyRSAKey()})
	)

	test := []struct {
		name      string
		req       string
		expectErr bool
		mock      func()
	}{
		{
			name:      "err bind request",
			req:       "asd",
			expectErr: true,
		},
		{
			name:      "err no row get profile by phone",
			req:       mockReq,
			expectErr: true,
			mock: func() {
				mockRepo.EXPECT().GetProfileByPhone(any, any).Return(mockUser, sql.ErrNoRows)
			},
		},
		{
			name: "err get profile by phone",
			req:  mockReq,
			mock: func() {
				mockRepo.EXPECT().GetProfileByPhone(any, any).Return(mockUser, mockErr)
			},
			expectErr: true,
		},
		{
			name: "err mismatch password",
			req:  mockReq,
			mock: func() {
				mockRepo.EXPECT().GetProfileByPhone(any, any).Return(repository.User{Password: "123"}, nil)
			},
			expectErr: true,
		},
		{
			name: "err update login count",
			req:  mockReq,
			mock: func() {
				mockRepo.EXPECT().GetProfileByPhone(any, any).Return(mockUser, nil)
				mockRepo.EXPECT().UpdateLoginCount(any, any, any).Return(mockErr)
			},
			expectErr: true,
		},
		{
			name: "success",
			req:  mockReq,
			mock: func() {
				mockRepo.EXPECT().GetProfileByPhone(any, any).Return(mockUser, nil)
				mockRepo.EXPECT().UpdateLoginCount(any, any, any).Return(nil)
			},
		},
	}

	for _, tt := range test {
		t.Run(tt.name, func(t *testing.T) {
			if tt.mock != nil {
				tt.mock()
			}

			req := httptest.NewRequest(http.MethodPost, reqPath, strings.NewReader(tt.req))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			err := server.Authenticate(c)

			if tt.expectErr {
				assert.Error(t, err)
			}
		})
	}
}

func TestProfile(t *testing.T) {
	var (
		// dependencies mock
		ctrl     = gomock.NewController(t)
		mockRepo = repository.NewMockRepositoryInterface(ctrl)

		// request and response mock
		any      = gomock.Any()
		mockErr  = errors.New("an error")
		mockUser = repository.User{Password: "$2a$10$OAN2DNAF79q/njdQAlnYF./iKq.5XYq/txWdlJnA1czE4IiAnZ86K", UpdatedAt: sql.NullTime{Valid: true}}

		// echo server mock
		e       = echo.New()
		reqPath = "/profile"
	)

	test := []struct {
		name      string
		token     string
		expectErr bool
		mock      func()
	}{
		{
			name:      "err empty token",
			token:     "",
			expectErr: true,
		},
		{
			name:      "err invalid token",
			token:     "Invalid",
			expectErr: true,
		},
		{
			name:      "err expired token",
			token:     dummyExpiredToken,
			expectErr: true,
		},
		{
			name:      "err get user profile by id",
			token:     dummyValidToken,
			expectErr: true,
			mock: func() {
				mockRepo.EXPECT().GetProfileByID(any, any).Return(mockUser, mockErr)
			},
		},
		{
			name:  "success",
			token: dummyValidToken,
			mock: func() {
				mockRepo.EXPECT().GetProfileByID(any, any).Return(mockUser, nil)
			},
		},
	}

	for _, tt := range test {
		t.Run(tt.name, func(t *testing.T) {
			if tt.mock != nil {
				tt.mock()
			}

			server := NewServer(NewServerOptions{mockRepo, getDummyRSAKey()})
			req := httptest.NewRequest(http.MethodPost, reqPath, nil)
			req.Header.Set(echo.HeaderAuthorization, tt.token)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			err := server.Profile(c)

			if tt.expectErr {
				assert.Error(t, err)
			}
		})
	}
}

func TestUpdateProfile(t *testing.T) {
	var (
		// dependencies mock
		ctrl     = gomock.NewController(t)
		mockRepo = repository.NewMockRepositoryInterface(ctrl)

		// request and response mock
		any          = gomock.Any()
		mockErr      = errors.New("an error")
		mockUser     = repository.User{Password: "$2a$10$OAN2DNAF79q/njdQAlnYF./iKq.5XYq/txWdlJnA1czE4IiAnZ86K", UpdatedAt: sql.NullTime{Valid: true}}
		mockValidReq = `{"name": "narto", "phone": "+6281122334455"}`
		// echo server mock
		e       = echo.New()
		reqPath = "/profile"
		server  = NewServer(NewServerOptions{mockRepo, getDummyRSAKey()})
	)

	test := []struct {
		name      string
		token     string
		req       string
		expectErr bool
		mock      func()
	}{
		{
			name:      "err invalid token",
			token:     "Invalid",
			expectErr: true,
		},
		{
			name:      "err invalid payload",
			token:     dummyValidToken,
			req:       "asd",
			expectErr: true,
		},
		{
			name:      "err get user profile by id",
			token:     dummyValidToken,
			req:       mockValidReq,
			expectErr: true,
			mock: func() {
				mockRepo.EXPECT().GetProfileByID(any, any).Return(mockUser, mockErr)
			},
		},
		{
			name:      "err update user",
			token:     dummyValidToken,
			req:       mockValidReq,
			expectErr: true,
			mock: func() {
				mockRepo.EXPECT().GetProfileByID(any, any).Return(mockUser, nil)
				mockRepo.EXPECT().UpdateUserByID(any, any).Return(mockErr)
			},
		},
		{
			name:  "has invalid name && invalid phone",
			token: dummyValidToken,
			req:   `{"name": "123", "phone": "+12233"}`,
			mock: func() {
				mockRepo.EXPECT().GetProfileByID(any, any).Return(mockUser, nil)
			},
		},
		{
			name:  "success",
			token: dummyValidToken,
			req:   mockValidReq,
			mock: func() {
				mockRepo.EXPECT().GetProfileByID(any, any).Return(mockUser, nil)
				mockRepo.EXPECT().UpdateUserByID(any, any).Return(nil)
			},
		},
	}

	for _, tt := range test {
		t.Run(tt.name, func(t *testing.T) {
			if tt.mock != nil {
				tt.mock()
			}

			req := httptest.NewRequest(http.MethodPut, reqPath, strings.NewReader(tt.req))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			req.Header.Set(echo.HeaderAuthorization, tt.token)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			err := server.UpdateProfile(c)

			if tt.expectErr {
				assert.Error(t, err)
			}
		})
	}
}
