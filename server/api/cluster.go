package api

import (
	"net/http"
	"time"

	"github.com/labstack/echo"

	"disconf/server/dao"
	"disconf/server/model/bean"
	model "disconf/server/model/http"
)

type Cluster struct {
}

func (a *Cluster) CreateCluster(context echo.Context) error {
	req := &model.CreateClusterRequest{}
	if err := context.Bind(req); err != nil {
		return context.String(http.StatusOK, "create cluster param is not correct")
	}
	cluster := &bean.Cluster{
		BaseModel: bean.BaseModel{
			CreateAt: time.Now(),
			UpdateAt: time.Now(),
			DeleteAt: time.Now(),
		},
		AppName:     req.AppName,
		ClusterName: req.ClusterName,
		Description: req.Description,
	}
	if err := dao.G_cluster.CreateCluster(cluster); err != nil {
		return context.String(http.StatusOK, err.Error())
	}
	return context.String(http.StatusOK, "create cluster success")
}
