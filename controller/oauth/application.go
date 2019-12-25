package oauth

import (
	app "github.com/codefluence-x/altair/controller/oauth/application"
	"github.com/codefluence-x/altair/core"
)

type application struct{}

func Application() core.OauthApplicationDispatcher {
	return application{}
}

func (a application) List(applicationManager core.ApplicationManager) core.Controller {
	return app.List(applicationManager)
}
