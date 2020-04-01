package dao

import (
	"fmt"

	"disconf/server/model/bean"
	"disconf/server/util"
)

var G_app = newAppDao()

type appDao struct {
}

func newAppDao() *appDao {
	return &appDao{}
}

func (d *appDao) CreateApp(app *bean.App) error {
	if d.IsAppExist(app.Name) {
		return fmt.Errorf("app has exist, name: %s", app.Name)
	}
	err := util.G_db.Create(app).Error
	return err
}

func (d *appDao) GetAppByName(name string) (*bean.App, error) {
	app := &bean.App{}
	err := util.G_db.Where("name = ?", name).Find(app).Error
	if err != nil {
		return nil, err
	}
	return app, nil
}

func (d *appDao) IsAppExist(appName string) bool {
	return  !util.G_db.Table("app").Where("name = ?", appName).First(&bean.App{}).RecordNotFound()
}