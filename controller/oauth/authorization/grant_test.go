package authorization_test

import (
	"bytes"
	"context"
	"database/sql"
	"encoding/json"
	"net/http"
	"testing"
	"time"

	"github.com/codefluence-x/altair/controller"
	"github.com/codefluence-x/altair/entity"
	"github.com/codefluence-x/altair/eobject"
	"github.com/codefluence-x/altair/formatter"
	"github.com/codefluence-x/altair/mock"
	"github.com/codefluence-x/altair/util"
	"github.com/codefluence-x/aurelia"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

type responseOneToken struct {
	Data entity.OauthAccessTokenJSON `json:"data"`
}

func TestGrant(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	t.Run("Method", func(t *testing.T) {
		authorizationService := mock.NewMockAuthorization(mockCtrl)
		assert.Equal(t, "POST", controller.Oauth().Authorization().Grant(authorizationService).Method())
	})

	t.Run("Path", func(t *testing.T) {
		authorizationService := mock.NewMockAuthorization(mockCtrl)
		assert.Equal(t, "/oauth/authorizations", controller.Oauth().Authorization().Grant(authorizationService).Path())
	})

	t.Run("Control", func(t *testing.T) {
		t.Run("Given request with json body", func(t *testing.T) {
			t.Run("Return oauth application data with status 202", func(t *testing.T) {
				apiEngine := gin.Default()

				authorizationRequest := entity.AuthorizationRequestJSON{
					ResponseType:    nil,
					ResourceOwnerID: util.IntToPointer(1),
					ClientUID:       util.StringToPointer(aurelia.Hash("x", "y")),
					ClientSecret:    util.StringToPointer(aurelia.Hash("z", "a")),
					RedirectURI:     util.StringToPointer("http://github.com"),
					Scopes:          util.StringToPointer("public users"),
				}
				encodedBytes, err := json.Marshal(authorizationRequest)
				assert.Nil(t, err)

				oauthAccessToken := entity.OauthAccessToken{
					ID:                 1,
					OauthApplicationID: 1,
					ResourceOwnerID:    *authorizationRequest.ResourceOwnerID,
					Token:              aurelia.Hash("x", "y"),
					Scopes: sql.NullString{
						String: *authorizationRequest.Scopes,
						Valid:  true,
					},
					ExpiresIn: time.Now().Add(time.Hour * 4),
					CreatedAt: time.Now().Truncate(time.Second),
				}
				oauthAccessTokenJSON := formatter.Oauth().AccessToken(authorizationRequest, oauthAccessToken)

				authorizationService := mock.NewMockAuthorization(mockCtrl)
				authorizationService.EXPECT().Grantor(gomock.Any(), authorizationRequest).Return(oauthAccessTokenJSON, nil)

				ctrl := controller.Oauth().Authorization().Grant(authorizationService)
				controller.Compile(apiEngine, ctrl)

				var response responseOneToken
				w := mock.PerformRequest(apiEngine, ctrl.Method(), ctrl.Path(), bytes.NewReader(encodedBytes))

				err = json.Unmarshal(w.Body.Bytes(), &response)
				assert.Nil(t, err)

				assert.Equal(t, http.StatusCreated, w.Code)
				assert.Equal(t, oauthAccessTokenJSON, response.Data)
			})

			t.Run("Unexpected error in authorization services", func(t *testing.T) {
				t.Run("Return entity error status", func(t *testing.T) {
					apiEngine := gin.Default()

					authorizationRequest := entity.AuthorizationRequestJSON{
						ResponseType:    nil,
						ResourceOwnerID: util.IntToPointer(1),
						ClientUID:       util.StringToPointer(aurelia.Hash("x", "y")),
						ClientSecret:    util.StringToPointer(aurelia.Hash("z", "a")),
						RedirectURI:     util.StringToPointer("http://github.com"),
						Scopes:          util.StringToPointer("public users"),
					}
					encodedBytes, err := json.Marshal(authorizationRequest)
					assert.Nil(t, err)

					oauthAccessToken := entity.OauthAccessToken{}
					oauthAccessTokenJSON := formatter.Oauth().AccessToken(authorizationRequest, oauthAccessToken)

					ctx := context.WithValue(context.Background(), "track_id", uuid.New().String())

					expectedError := &entity.Error{
						HttpStatus: http.StatusNotFound,
						Errors:     eobject.Wrap(eobject.NotFoundError(ctx, "client_uid & client_secret")),
					}

					authorizationService := mock.NewMockAuthorization(mockCtrl)
					authorizationService.EXPECT().Grantor(gomock.Any(), authorizationRequest).Return(oauthAccessTokenJSON, expectedError)

					ctrl := controller.Oauth().Authorization().Grant(authorizationService)
					controller.Compile(apiEngine, ctrl)

					var response mock.ErrorResponse
					w := mock.PerformRequest(apiEngine, ctrl.Method(), ctrl.Path(), bytes.NewReader(encodedBytes))

					err = json.Unmarshal(w.Body.Bytes(), &response)
					assert.Nil(t, err)

					assert.Equal(t, http.StatusNotFound, w.Code)
					assert.Equal(t, expectedError.Errors, response.Errors)
				})
			})
		})

		t.Run("Given invalid request body", func(t *testing.T) {
			t.Run("Return bad request", func(t *testing.T) {
				apiEngine := gin.Default()

				authorizationService := mock.NewMockAuthorization(mockCtrl)
				authorizationService.EXPECT().Grantor(gomock.Any(), gomock.Any()).Times(0)

				ctrl := controller.Oauth().Authorization().Grant(authorizationService)
				controller.Compile(apiEngine, ctrl)

				expectedError := &entity.Error{
					HttpStatus: http.StatusBadRequest,
					Errors:     eobject.Wrap(eobject.BadRequestError("request body")),
				}

				var response mock.ErrorResponse
				w := mock.PerformRequest(apiEngine, ctrl.Method(), ctrl.Path(), mock.MockErrorIoReader{})

				err := json.Unmarshal(w.Body.Bytes(), &response)
				assert.Nil(t, err)

				assert.Equal(t, expectedError.HttpStatus, w.Code)
				assert.Equal(t, expectedError.Errors, response.Errors)
			})
		})

		t.Run("Given request body but not json", func(t *testing.T) {
			t.Run("Return bad request", func(t *testing.T) {
				apiEngine := gin.Default()

				authorizationService := mock.NewMockAuthorization(mockCtrl)
				authorizationService.EXPECT().Grantor(gomock.Any(), gomock.Any()).Times(0)

				ctrl := controller.Oauth().Authorization().Grant(authorizationService)
				controller.Compile(apiEngine, ctrl)

				expectedError := &entity.Error{
					HttpStatus: http.StatusBadRequest,
					Errors:     eobject.Wrap(eobject.BadRequestError("request body")),
				}

				var response mock.ErrorResponse
				w := mock.PerformRequest(apiEngine, ctrl.Method(), ctrl.Path(), bytes.NewReader([]byte(`this is gonna be error`)))

				err := json.Unmarshal(w.Body.Bytes(), &response)
				assert.Nil(t, err)

				assert.Equal(t, expectedError.HttpStatus, w.Code)
				assert.Equal(t, expectedError.Errors, response.Errors)
			})
		})

	})
}
