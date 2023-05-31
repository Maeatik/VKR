package repository_test

import (
	"context"
	v1 "diploma/internal/entity/v1"
	"diploma/internal/repository"
	mock_repository "diploma/internal/repository/mocks"
	"diploma/internal/service"
	"fmt"

	"github.com/stretchr/testify/assert"

	"testing"

	"github.com/golang/mock/gomock"
)

func Test_getReport(t *testing.T) {
	type TestRequestParams struct {
		id          int
		name        string
		oldPassword string
		newPassword string
	}

	type mockBehavior func(mockAuth *mock_repository.MockAuthorization, mockServ *mock_repository.MockService, reqParams TestRequestParams)

	testTable := map[string]struct {
		name                 string
		inputStruct          TestRequestParams
		mockBehavior         mockBehavior
		expexctedRequestBody error
	}{
		"1. GoodData": {
			inputStruct: TestRequestParams{
				id:          1,
				name:        "user",
				oldPassword: service.GeneratePasswordHash("oldPassword"),
				newPassword: service.GeneratePasswordHash("newPassword"),
			},

			mockBehavior: func(mockAuth *mock_repository.MockAuthorization, mockServ *mock_repository.MockService, reqParams TestRequestParams) {
				ctx := context.Background()
				mockServ.EXPECT().GetUser(ctx, reqParams.id).Return(v1.User{
					Name:     reqParams.name,
					Password: reqParams.oldPassword,
				}, nil).AnyTimes()

				mockServ.EXPECT().UpdateUsers(ctx, reqParams.id, reqParams.name, reqParams.newPassword).Return(nil).AnyTimes()
			},
			expexctedRequestBody: nil,
		},
		"2. BadData": {
			inputStruct: TestRequestParams{
				id:          1,
				name:        "user",
				oldPassword: service.GeneratePasswordHash("oldPassword123"),
				newPassword: service.GeneratePasswordHash("newPassword"),
			},

			mockBehavior: func(mockAuth *mock_repository.MockAuthorization, mockServ *mock_repository.MockService, reqParams TestRequestParams) {
				ctx := context.Background()
				mockServ.EXPECT().GetUser(ctx, reqParams.id).Return(v1.User{
					Name:     reqParams.name,
					Password: reqParams.oldPassword,
				}, nil).AnyTimes()
			},
			expexctedRequestBody: fmt.Errorf("dismatches passwords"),
		},
	}

	for name, testCase := range testTable {
		t.Run(name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			ctx := context.Background()
			repoAuth := mock_repository.NewMockAuthorization(c)
			repoService := mock_repository.NewMockService(c)

			testCase.mockBehavior(repoAuth, repoService, testCase.inputStruct)

			repository := repository.Repository{
				Authorization: repoAuth,
				Service:       repoService,
			}

			serviceMock := service.NewServices(&repository)

			err := serviceMock.ChangePassword(ctx, 1, "oldPassword", "newPassword")

			assert.Equal(t, err, testCase.expexctedRequestBody)
		})
	}
}
