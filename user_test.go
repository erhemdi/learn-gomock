package user

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func Test_UserService_GetData(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	type args struct {
		userID int64
	}

	tests := []struct {
		name         string
		args         args
		mock         func(redisMock *MockUserRedisIface, dbMock *MockUserDBIface)
		expectedUser User
		expectedErr  error
	}{
		{
			name: "failed get data from redis",
			args: args{
				userID: 123,
			},
			mock: func(redisMock *MockUserRedisIface, dbMock *MockUserDBIface) {
				redisMock.EXPECT().GetData(
					gomock.Eq(int64(123)),
				).Return(
					User{}, errors.New("error redis connection"),
				)
			},
			expectedUser: User{},
			expectedErr:  errors.New("error redis connection"),
		},
		{
			name: "user data found in redis",
			args: args{
				userID: 123,
			},
			mock: func(redisMock *MockUserRedisIface, dbMock *MockUserDBIface) {
				redisMock.EXPECT().GetData(
					gomock.Eq(int64(123)),
				).Return(
					User{
						ID:   123,
						Name: "Bob",
					}, nil,
				)
			},
			expectedUser: User{
				ID:   123,
				Name: "Bob",
			},
			expectedErr: nil,
		},
		{
			name: "failed get data from DB",
			args: args{
				userID: 123,
			},
			mock: func(redisMock *MockUserRedisIface, dbMock *MockUserDBIface) {
				redisMock.EXPECT().GetData(
					gomock.Eq(int64(123)),
				).Return(
					User{}, nil,
				)

				dbMock.EXPECT().GetData(
					gomock.Eq(int64(123)),
				).Return(
					User{}, errors.New("error db connection"),
				)
			},
			expectedUser: User{},
			expectedErr:  errors.New("error db connection"),
		},
		{
			name: "success get data from DB",
			args: args{
				userID: 123,
			},
			mock: func(redisMock *MockUserRedisIface, dbMock *MockUserDBIface) {
				redisMock.EXPECT().GetData(
					gomock.Eq(int64(123)),
				).Return(
					User{}, nil,
				)

				dbMock.EXPECT().GetData(
					gomock.Eq(int64(123)),
				).Return(
					User{
						ID:   123,
						Name: "Bob",
					}, nil,
				)
			},
			expectedUser: User{
				ID:   123,
				Name: "Bob",
			},
			expectedErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			redisMock := NewMockUserRedisIface(ctrl)
			dbMock := NewMockUserDBIface(ctrl)
			tt.mock(redisMock, dbMock)

			userService := UserService{
				Redis: redisMock,
				DB:    dbMock,
			}

			user, err := userService.GetData(tt.args.userID)
			assert.Equal(t, tt.expectedUser, user)
			assert.Equal(t, tt.expectedErr, err)
		})
	}
}
