package api

import (
	"net/http"
	"time"

	"github.com/labstack/echo"

	"disconf/server/dao"
	"disconf/server/model"
	"disconf/server/model/bean"
	httpmodel "disconf/server/model/http"
	tcpmodel "disconf/server/model/tcp"
	"disconf/server/util"
)

type Release struct {
	Server model.Server
}

func (a *Release) PubcConfig(context echo.Context) error {
	req := &httpmodel.CreateReleaseRequest{}
	if err := context.Bind(req); err != nil {
		return context.String(http.StatusOK, "pub config param is not correct")
	}
	release := &bean.Release{
		BaseModel: bean.BaseModel{
			CreateAt: time.Now(),
			UpdateAt: time.Now(),
			DeleteAt: time.Now(),
		},
		AppName:       req.AppName,
		ClusterName:   req.ClusterName,
		NamespaceName: req.NamespaceName,
		Version:       req.Version,
		Format:        req.Format,
		Value:         req.Value,
		Description:   req.Description,
	}
	if err := dao.G_release.CreateRelease(release); err != nil {
		return context.String(http.StatusOK, err.Error())
	}
	msg := &tcpmodel.NotifyMsg{
		Format:    release.Format,
		Namespace: generateNamespace(release.AppName, release.ClusterName, release.NamespaceName),
		Value:     release.Value,
	}
	util.G_log.Info("[start notify client] [namespace:%s] [format:%s] [value:%s]", msg.Namespace, msg.Format, msg.Value)
	a.Server.NotifyEvent(msg)
	return context.String(http.StatusOK, "pub config success")
}

func generateNamespace(appName, clusterName, namespaceName string) string {
	return appName + "/" + clusterName + "/" + namespaceName
}
