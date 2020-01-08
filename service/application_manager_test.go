package service_test

import (
	"context"
	"database/sql"
	"errors"
	"net/http"
	"testing"

	"github.com/codefluence-x/altair/entity"
	"github.com/codefluence-x/altair/eobject"
	"github.com/codefluence-x/altair/formatter"
	"github.com/codefluence-x/altair/mock"
	"github.com/codefluence-x/altair/service"
	"github.com/codefluence-x/altair/util"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestApplicationManager(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	t.Run("List", func(t *testing.T) {
		t.Run("Given limit and offset", func(t *testing.T) {
			t.Run("Return a formatted result", func(t *testing.T) {

				oauthApplications := []entity.OauthApplication{
					entity.OauthApplication{ID: 1},
					entity.OauthApplication{ID: 2},
				}

				applicationFormatter := formatter.OauthApplication()

				ctx := context.Background()

				oauthModel := mock.NewMockOauthApplicationModel(mockCtrl)

				gomock.InOrder(
					oauthModel.EXPECT().Paginate(ctx, 0, 10).Return(oauthApplications, nil),
					oauthModel.EXPECT().Count(ctx).Return(len(oauthApplications), nil),
				)

				applicationManager := service.ApplicationManager(applicationFormatter, oauthModel)

				oauthApplicationJSON, total, err := applicationManager.List(ctx, 0, 10)

				assert.Nil(t, err)
				assert.Equal(t, len(oauthApplications), total)
				assert.Equal(t, applicationFormatter.ApplicationList(ctx, oauthApplications), oauthApplicationJSON)

			})

			t.Run("Error oauth application paginate return internal server error", func(t *testing.T) {

				ctx := context.Background()

				oauthModel := mock.NewMockOauthApplicationModel(mockCtrl)

				expectedError := &entity.Error{
					HttpStatus: http.StatusInternalServerError,
					Errors:     eobject.Wrap(eobject.InternalServerError(ctx)),
				}

				gomock.InOrder(
					oauthModel.EXPECT().Paginate(ctx, 0, 10).Return([]entity.OauthApplication(nil), expectedError),
					oauthModel.EXPECT().Count(gomock.Any()).Times(0),
				)

				applicationManager := service.ApplicationManager(formatter.OauthApplication(), oauthModel)

				oauthApplicationJSON, total, err := applicationManager.List(ctx, 0, 10)

				assert.NotNil(t, err)
				assert.Equal(t, expectedError, err)
				assert.Equal(t, 0, total)
				assert.Equal(t, []entity.OauthApplicationJSON(nil), oauthApplicationJSON)
			})

			t.Run("Error oauth application count return internal server error", func(t *testing.T) {
				oauthApplications := []entity.OauthApplication{
					entity.OauthApplication{ID: 1},
					entity.OauthApplication{ID: 2},
				}

				ctx := context.Background()

				oauthModel := mock.NewMockOauthApplicationModel(mockCtrl)

				expectedError := &entity.Error{
					HttpStatus: http.StatusInternalServerError,
					Errors:     eobject.Wrap(eobject.InternalServerError(ctx)),
				}

				gomock.InOrder(
					oauthModel.EXPECT().Paginate(ctx, 0, 10).Return(oauthApplications, nil),
					oauthModel.EXPECT().Count(ctx).Return(0, expectedError),
				)

				applicationManager := service.ApplicationManager(formatter.OauthApplication(), oauthModel)

				oauthApplicationJSON, total, err := applicationManager.List(ctx, 0, 10)

				assert.NotNil(t, err)
				assert.Equal(t, expectedError, err)
				assert.Equal(t, 0, total)
				assert.Equal(t, []entity.OauthApplicationJSON(nil), oauthApplicationJSON)
			})
		})
	})

	t.Run("One", func(t *testing.T) {
		t.Run("Given context and oauth application id", func(t *testing.T) {
			t.Run("Return oauth application data", func(t *testing.T) {
				ctx := context.Background()

				expectedData := entity.OauthApplication{ID: 1}

				oauthModel := mock.NewMockOauthApplicationModel(mockCtrl)
				oauthModel.EXPECT().One(ctx, 1).Return(expectedData, nil)

				applicationManager := service.ApplicationManager(formatter.OauthApplication(), oauthModel)
				oauthApplicationJSON, err := applicationManager.One(ctx, 1)

				assert.Nil(t, err)
				assert.Equal(t, formatter.OauthApplication().Application(ctx, expectedData), oauthApplicationJSON)
			})

			t.Run("Oauth application is not found", func(t *testing.T) {
				t.Run("Return 404", func(t *testing.T) {
					ctx := context.Background()

					oauthModel := mock.NewMockOauthApplicationModel(mockCtrl)
					oauthModel.EXPECT().One(ctx, 1).Return(entity.OauthApplication{}, sql.ErrNoRows)

					applicationManager := service.ApplicationManager(formatter.OauthApplication(), oauthModel)
					oauthApplicationJSON, err := applicationManager.One(ctx, 1)

					expectedError := &entity.Error{
						HttpStatus: http.StatusNotFound,
						Errors:     eobject.Wrap(eobject.NotFoundError(ctx, "oauth_application")),
					}

					assert.NotNil(t, err)
					assert.Equal(t, expectedError, err)
					assert.Equal(t, entity.OauthApplicationJSON{}, oauthApplicationJSON)
				})
			})

			t.Run("Unexpected error", func(t *testing.T) {
				t.Run("Return internal server error", func(t *testing.T) {
					ctx := context.Background()

					oauthModel := mock.NewMockOauthApplicationModel(mockCtrl)
					oauthModel.EXPECT().One(ctx, 1).Return(entity.OauthApplication{}, errors.New("Unexpected error"))

					applicationManager := service.ApplicationManager(formatter.OauthApplication(), oauthModel)
					oauthApplicationJSON, err := applicationManager.One(ctx, 1)

					expectedError := &entity.Error{
						HttpStatus: http.StatusInternalServerError,
						Errors:     eobject.Wrap(eobject.InternalServerError(ctx)),
					}

					assert.NotNil(t, err)
					assert.Equal(t, expectedError, err)
					assert.Equal(t, entity.OauthApplicationJSON{}, oauthApplicationJSON)
				})
			})
		})
	})

	t.Run("Create", func(t *testing.T) {
		t.Run("Given context and oauth application json data", func(t *testing.T) {
			t.Run("Return last inserted id", func(t *testing.T) {
				ctx := context.Background()
				oauthApplication := entity.OauthApplication{
					ID: 1,
				}
				expectedOauthApplicationJSON := formatter.OauthApplication().Application(ctx, oauthApplication)
				oauthModel := mock.NewMockOauthApplicationModel(mockCtrl)

				gomock.InOrder(
					oauthModel.EXPECT().Create(ctx, expectedOauthApplicationJSON).Return(1, nil),
					oauthModel.EXPECT().One(ctx, 1).Return(oauthApplication, nil),
				)

				applicationManager := service.ApplicationManager(formatter.OauthApplication(), oauthModel)
				formattedResult, err := applicationManager.Create(ctx, expectedOauthApplicationJSON)

				assert.Nil(t, err)
				assert.Equal(t, expectedOauthApplicationJSON, formattedResult)
			})

			t.Run("Unexpected error", func(t *testing.T) {
				t.Run("Return internal server error", func(t *testing.T) {
					ctx := context.Background()

					data := entity.OauthApplicationJSON{
						ID: util.IntToPointer(1),
					}

					oauthModel := mock.NewMockOauthApplicationModel(mockCtrl)
					oauthModel.EXPECT().One(gomock.Any(), gomock.Any()).Times(0)

					gomock.InOrder(
						oauthModel.EXPECT().Create(ctx, data).Return(0, errors.New("Unexpected error")),
					)

					applicationManager := service.ApplicationManager(formatter.OauthApplication(), oauthModel)
					formattedResult, err := applicationManager.Create(ctx, data)

					expectedError := &entity.Error{
						HttpStatus: http.StatusInternalServerError,
						Errors:     eobject.Wrap(eobject.InternalServerError(ctx)),
					}

					assert.NotNil(t, err)
					assert.Equal(t, expectedError, err)
					assert.Equal(t, entity.OauthApplicationJSON{}, formattedResult)
				})
			})
		})
	})
}
