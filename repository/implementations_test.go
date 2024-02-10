package repository

import (
	"context"
	"errors"
	"testing"
	"time"

	"gopkg.in/DATA-DOG/go-sqlmock.v1"
)

var mockProfileColumn = []string{"id", "name", "phone", "password", "login_count", "created_at", "updated_at"}

func TestSaveProfile(t *testing.T) {
	var (
		// mock dependencies
		db, mock, _ = sqlmock.New()
		mockQuery   = "insert into profile"

		// mock request and responser
		mockErr = errors.New("an error")

		r = NewRepository(NewRepositoryOptions{Db: db})
	)

	test := []struct {
		name      string
		mock      func()
		expectErr bool
	}{
		{
			name:      "error scan",
			expectErr: true,
			mock: func() {
				mock.ExpectQuery(mockQuery).WillReturnError(mockErr)
			},
		},
		{
			name: "success",
			mock: func() {
				mock.ExpectQuery(mockQuery).WillReturnRows(
					sqlmock.NewRows([]string{"id"}).AddRow(1),
				)
			},
		},
	}

	for _, tt := range test {
		t.Run(tt.name, func(t *testing.T) {
			if tt.mock != nil {
				tt.mock()
			}
		})

		_, err := r.SaveProfile(context.Background(), User{})
		if (err != nil) != tt.expectErr {
			t.Error(err)
		}
	}
}

func TestGetProfileByPhone(t *testing.T) {
	var (
		// mock dependencies
		db, mock, _ = sqlmock.New()
		mockQuery   = "select (.+) from profile where phone"

		// mock request and responser
		mockErr = errors.New("an error")

		r = NewRepository(NewRepositoryOptions{Db: db})
	)

	test := []struct {
		name      string
		mock      func()
		expectErr bool
	}{
		{
			name:      "error scan",
			expectErr: true,
			mock: func() {
				mock.ExpectQuery(mockQuery).WillReturnError(mockErr)
			},
		},
		{
			name: "success",
			mock: func() {
				mock.ExpectQuery(mockQuery).WillReturnRows(
					sqlmock.NewRows(mockProfileColumn).
						AddRow(1, "narto", "+62", "pass", 1, time.Now(), nil),
				)
			},
		},
	}

	for _, tt := range test {
		t.Run(tt.name, func(t *testing.T) {
			if tt.mock != nil {
				tt.mock()
			}
		})

		_, err := r.GetProfileByPhone(context.Background(), "")
		if (err != nil) != tt.expectErr {
			t.Error(err)
		}
	}
}

func TestGetProfileByID(t *testing.T) {
	var (
		// mock dependencies
		db, mock, _ = sqlmock.New()
		mockQuery   = "select (.+) from profile (.+) id ="

		// mock request and responser
		mockErr = errors.New("an error")

		r = NewRepository(NewRepositoryOptions{Db: db})
	)

	test := []struct {
		name      string
		mock      func()
		expectErr bool
	}{
		{
			name:      "error scan",
			expectErr: true,
			mock: func() {
				mock.ExpectQuery(mockQuery).WillReturnError(mockErr)
			},
		},
		{
			name: "success",
			mock: func() {
				mock.ExpectQuery(mockQuery).WillReturnRows(
					sqlmock.NewRows(mockProfileColumn).
						AddRow(1, "narto", "+62", "pass", 1, time.Now(), nil),
				)
			},
		},
	}

	for _, tt := range test {
		t.Run(tt.name, func(t *testing.T) {
			if tt.mock != nil {
				tt.mock()
			}
		})

		_, err := r.GetProfileByID(context.Background(), 1)
		if (err != nil) != tt.expectErr {
			t.Error(err)
		}
	}
}

func TestGetUpdateLoginCount(t *testing.T) {
	var (
		// mock dependencies
		db, mock, _ = sqlmock.New()
		mockQuery   = "update profile set login_count (.+) where id ="

		// mock request and responser
		mockErr = errors.New("an error")

		r = NewRepository(NewRepositoryOptions{Db: db})
	)

	test := []struct {
		name      string
		mock      func()
		expectErr bool
	}{
		{
			name:      "error scan",
			expectErr: true,
			mock: func() {
				mock.ExpectExec(mockQuery).WillReturnError(mockErr)
			},
		},
		{
			name: "success",
			mock: func() {
				mock.ExpectExec(mockQuery).WillReturnResult(sqlmock.NewResult(0, 1))
			},
		},
	}

	for _, tt := range test {
		t.Run(tt.name, func(t *testing.T) {
			if tt.mock != nil {
				tt.mock()
			}
		})

		err := r.UpdateLoginCount(context.Background(), 1, 1)
		if (err != nil) != tt.expectErr {
			t.Error(err)
		}
	}
}

func TestUpdateUserByID(t *testing.T) {
	var (
		// mock dependencies
		db, mock, _ = sqlmock.New()
		mockQuery   = "update profile (.+) where id ="

		// mock request and responser
		mockErr = errors.New("an error")

		r = NewRepository(NewRepositoryOptions{Db: db})
	)

	test := []struct {
		name      string
		mock      func()
		expectErr bool
	}{
		{
			name:      "error scan",
			expectErr: true,
			mock: func() {
				mock.ExpectExec(mockQuery).WillReturnError(mockErr)
			},
		},
		{
			name: "success",
			mock: func() {
				mock.ExpectExec(mockQuery).WillReturnResult(sqlmock.NewResult(0, 1))
			},
		},
	}

	for _, tt := range test {
		t.Run(tt.name, func(t *testing.T) {
			if tt.mock != nil {
				tt.mock()
			}
		})

		err := r.UpdateUserByID(context.Background(), User{})
		if (err != nil) != tt.expectErr {
			t.Error(err)
		}
	}
}
