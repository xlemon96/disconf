package api

import (
	"net/http"
	"time"

	"github.com/labstack/echo"

	"disconf/server/dao"
	"disconf/server/model/bean"
	model "disconf/server/model/http"
)

type App struct {
}

func (a *App) CreateApp(context echo.Context) error {
	req := &model.CreateAppRequest{}
	if err := context.Bind(req); err != nil {
		return context.String(http.StatusOK, "create app param is not correct")
	}
	app := &bean.App{
		BaseModel:   bean.BaseModel{
			CreateAt: time.Now(),
			UpdateAt: time.Now(),
			DeleteAt: time.Now(),
		},
		Name:        req.Name,
		Description: req.Description,
	}
	if err := dao.G_app.CreateApp(app); err != nil {
		return context.String(http.StatusOK, err.Error())
	}
	return context.String(http.StatusOK, "create app success")
}
