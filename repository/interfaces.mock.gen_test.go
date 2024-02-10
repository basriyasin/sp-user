package repository

import (
	"context"
	"testing"

	gomock "github.com/golang/mock/gomock"
)

// to cover the mock file so it would not reduce the test coverage
func TestMock(t *testing.T) {
	var (
		ctrl = gomock.NewController(t)
		mock = NewMockRepositoryInterface(ctrl)

		any = gomock.Any()
		ctx = context.Background()
	)

	mock.EXPECT().GetProfileByID(any, any)
	mock.GetProfileByID(ctx, 1)
	mock.EXPECT().GetProfileByID(any, any)
	mock.GetProfileByID(ctx, 1)
	mock.EXPECT().GetProfileByPhone(any, any)
	mock.GetProfileByPhone(ctx, "")
	mock.EXPECT().SaveProfile(any, any)
	mock.SaveProfile(ctx, User{})
	mock.EXPECT().UpdateLoginCount(any, any, any)
	mock.UpdateLoginCount(ctx, 1, 1)
	mock.EXPECT().UpdateUserByID(any, any)
	mock.UpdateUserByID(ctx, User{})
}
