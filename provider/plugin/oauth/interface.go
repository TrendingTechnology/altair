package oauth

import (
	"context"
	"database/sql"

	"github.com/codefluence-x/altair/entity"
)

type DBExecutable interface {
	Exec(query string, args ...interface{}) (sql.Result, error)
}

type OauthApplicationModel interface {
	Name() string
	Paginate(ctx context.Context, offset, limit int) ([]entity.OauthApplication, error)
	One(ctx context.Context, ID int) (entity.OauthApplication, error)
	OneByUIDandSecret(ctx context.Context, clientUID, clientSecret string) (entity.OauthApplication, error)
	Count(ctx context.Context) (int, error)
	Create(ctx context.Context, data entity.OauthApplicationInsertable, txs ...*sql.Tx) (int, error)
}

type OauthAccessTokenModel interface {
	Name() string
	One(ctx context.Context, ID int) (entity.OauthAccessToken, error)
	OneByToken(ctx context.Context, token string) (entity.OauthAccessToken, error)
	Create(ctx context.Context, data entity.OauthAccessTokenInsertable, txs ...*sql.Tx) (int, error)
	Revoke(ctx context.Context, token string) error
}

type OauthAccessGrantModel interface {
	Name() string
	One(ctx context.Context, ID int) (entity.OauthAccessGrant, error)
	Create(ctx context.Context, data entity.OauthAccessGrantInsertable, txs ...*sql.Tx) (int, error)
}

type ApplicationManager interface {
	List(ctx context.Context, offset, limit int) ([]entity.OauthApplicationJSON, int, *entity.Error)
	One(ctx context.Context, ID int) (entity.OauthApplicationJSON, *entity.Error)
	Create(ctx context.Context, e entity.OauthApplicationJSON) (entity.OauthApplicationJSON, *entity.Error)
}

type Authorization interface {
	Grantor(ctx context.Context, authorizationReq entity.AuthorizationRequestJSON) (interface{}, *entity.Error)
	Grant(ctx context.Context, authorizationReq entity.AuthorizationRequestJSON) (entity.OauthAccessGrantJSON, *entity.Error)
	GrantToken(ctx context.Context, authorizationReq entity.AuthorizationRequestJSON) (entity.OauthAccessTokenJSON, *entity.Error)
	Token(ctx context.Context, accessTokenReq entity.AccessTokenRequestJSON) (entity.OauthAccessTokenJSON, *entity.Error)
	RevokeToken(ctx context.Context, revokeAccessTokenReq entity.RevokeAccessTokenRequestJSON) *entity.Error
}