package oauth

import (
	"time"

	"github.com/codefluence-x/altair/core"
	"github.com/codefluence-x/altair/entity"
	"github.com/codefluence-x/altair/provider/plugin/oauth/controller"
	"github.com/codefluence-x/altair/provider/plugin/oauth/formatter"
	"github.com/codefluence-x/altair/provider/plugin/oauth/model"
	"github.com/codefluence-x/altair/provider/plugin/oauth/service"
	"github.com/codefluence-x/altair/provider/plugin/oauth/validator"
)

func Provide(appBearer core.AppBearer, dbBearer core.DatabaseBearer, pluginBearer core.PluginBearer) error {
	if appBearer.Config().PluginExists("oauth") == false {
		return nil
	}

	var oauthPluginConfig entity.OauthPlugin

	if err := pluginBearer.CompilePlugin("oauth", &oauthPluginConfig); err != nil {
		return err
	}

	db, _, err := dbBearer.Database(oauthPluginConfig.DatabaseInstance())
	if err != nil {
		return err
	}

	var accessTokenTimeout time.Duration
	var authorizationCodeTimeout time.Duration

	accessTokenTimeout, err = oauthPluginConfig.AccessTokenTimeout()
	if err != nil {
		return err
	}

	authorizationCodeTimeout, err = oauthPluginConfig.AuthorizationCodeTimeout()
	if err != nil {
		return err
	}

	// Model
	oauthApplicationModel := model.OauthApplication(db)
	oauthAccessTokenModel := model.OauthAccessToken(db)
	oauthAccessGrantModel := model.OauthAccessGrant(db)

	// Formatter
	oauthApplicationFormatter := formatter.OauthApplication()
	oauthModelFormatter := formatter.Model(accessTokenTimeout, authorizationCodeTimeout)
	oauthFormatter := formatter.Oauth()

	// Validator
	oauthValidator := validator.Oauth()

	// Service
	applicationManager := service.ApplicationManager(oauthApplicationFormatter, oauthModelFormatter, oauthApplicationModel, oauthValidator)
	authorization := service.Authorization(oauthApplicationModel, oauthAccessTokenModel, oauthAccessGrantModel, oauthModelFormatter, oauthValidator, oauthFormatter)

	appBearer.InjectController(controller.Application().List(applicationManager))
	appBearer.InjectController(controller.Application().One(applicationManager))
	appBearer.InjectController(controller.Application().Create(applicationManager))
	appBearer.InjectController(controller.Authorization().Grant(authorization))
	appBearer.InjectController(controller.Authorization().Revoke(authorization))

	return nil
}
