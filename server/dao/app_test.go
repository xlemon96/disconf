package dao

import (
	"fmt"
	"testing"
	"time"

	"disconf/server/model/bean"
	"disconf/server/util"
)

func init() {
	err := util.InitConfig("/Users/jiajianyun/go/src/disconf/server/etc/config.yaml")
	if err != nil {
		panic(err)
	}
	err = util.InitDB()
	if err != nil {
		panic(err)
	}
}

func TestAppDao_CreateApp(t *testing.T) {
	app := &bean.App{
		BaseModel: bean.BaseModel{
			CreateAt:  time.Now(),
			UpdateAt:  time.Now(),
			DeleteAt: time.Now(),
		},
		Name:        "test_ap",
		Description: "test_description",
	}
	err := G_app.CreateApp(app)
	if err != nil {
		panic(err)
	}
}

func TestAppDao_GetAppByName(t *testing.T) {
	//correct
	app, err := G_app.GetAppByName("test_app")
	if err != nil {
		panic(err)
	}
	fmt.Println(app)
}

func TestAppDao_IsAppExist(t *testing.T) {
	appName := "test"
	fmt.Println(G_app.IsAppExist(appName))
}