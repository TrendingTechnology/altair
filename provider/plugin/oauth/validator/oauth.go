package validator

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/codefluence-x/altair/provider/plugin/oauth/entity"
	"github.com/codefluence-x/altair/provider/plugin/oauth/eobject"
	"github.com/codefluence-x/altair/util"
)

type application struct {
}

func Oauth() *application {
	return &application{}
}

func (a *application) ValidateApplication(ctx context.Context, data entity.OauthApplicationJSON) *entity.Error {
	var entityError = &entity.Error{}

	if data.OwnerType == nil {
		entityError.Errors = append(entityError.Errors, eobject.ValidationError("object `owner_type` is nil or not exists"))
	} else {
		if *data.OwnerType != "confidential" && *data.OwnerType != "public" {
			entityError.Errors = append(entityError.Errors, eobject.ValidationError("object `owner_type` must be either of `confidential` or `public`"))
		}
	}

	if len(entityError.Errors) > 0 {
		entityError.HttpStatus = http.StatusUnprocessableEntity
		return entityError
	}

	return nil
}

func (a *application) ValidateAuthorizationGrant(ctx context.Context, r entity.AuthorizationRequestJSON, application entity.OauthApplication) *entity.Error {
	var entityErr = &entity.Error{}

	if r.ResponseType == nil {
		entityErr.Errors = append(entityErr.Errors, eobject.ValidationError(`response_type can't be empty`))
	}

	if r.ResourceOwnerID == nil {
		entityErr.Errors = append(entityErr.Errors, eobject.ValidationError(`resource_owner_id can't be empty`))
	}

	if r.RedirectURI == nil {
		entityErr.Errors = append(entityErr.Errors, eobject.ValidationError(`redirect_uri can't be empty`))
	}

	if len(entityErr.Errors) > 0 {
		entityErr.HttpStatus = http.StatusUnprocessableEntity
		return entityErr
	}

	if r.Scopes == nil {
		r.Scopes = util.StringToPointer("")
	}

	requestScopes := strings.Fields(*r.Scopes)
	applicationScopes := strings.Fields(application.Scopes.String)

	var invalidScope []string

	for _, rs := range requestScopes {

		scopeNotExists := true

		for _, as := range applicationScopes {
			if rs == as {
				scopeNotExists = false
				break
			}
		}

		if scopeNotExists {
			invalidScope = append(invalidScope, rs)
		}
	}

	if len(invalidScope) > 0 {
		return &entity.Error{
			HttpStatus: http.StatusForbidden,
			Errors:     eobject.Wrap(eobject.ForbiddenError(ctx, "application", fmt.Sprintf("your requested scopes `(%v)` is not exists in application", invalidScope))),
		}
	}

	if *r.ResponseType == "token" && application.OwnerType != "confidential" {
		return &entity.Error{
			HttpStatus: http.StatusForbidden,
			Errors:     eobject.Wrap(eobject.ForbiddenError(ctx, "access_token", "your response type is not allowed in this application")),
		}
	}

	return nil
}

func (a *application) ValidateTokenGrant(ctx context.Context, r entity.AccessTokenRequestJSON) *entity.Error {
	var entityErr = &entity.Error{}

	if r.GrantType == nil {
		entityErr.Errors = append(entityErr.Errors, eobject.ValidationError(`grant_type can't be empty`))
	} else if *r.GrantType != "authorization_code" {
		entityErr.Errors = append(entityErr.Errors, eobject.ValidationError(`grant_type must be set to 'authorization_code'`))
	}

	if r.Code == nil {
		entityErr.Errors = append(entityErr.Errors, eobject.ValidationError(`code can't be empty`))
	}

	if r.RedirectURI == nil {
		entityErr.Errors = append(entityErr.Errors, eobject.ValidationError(`redirect_uri can't be empty`))
	}

	if len(entityErr.Errors) > 0 {
		entityErr.HttpStatus = http.StatusUnprocessableEntity
		return entityErr
	}

	return nil
}
