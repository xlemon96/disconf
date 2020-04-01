package dao

import (
	"fmt"

	"disconf/server/model/bean"
	"disconf/server/util"
)

var G_cluster = newClusterDao()

type clusterDao struct {
}

func newClusterDao() *clusterDao {
	return &clusterDao{}
}

func (d *clusterDao) CreateCluster(cluster *bean.Cluster) error {
	if !G_app.IsAppExist(cluster.AppName) {
		return fmt.Errorf("app is not exist, %s", cluster.AppName)
	}
	if d.IsClusterExist(cluster.AppName, cluster.ClusterName) {
		return fmt.Errorf("cluster has exist, %s", cluster.ClusterName)
	}
	if err := util.G_db.Create(cluster).Error; err != nil {
		return err
	}
	return nil
}

func (d *clusterDao) GetCluster(appName, clusterName string) (*bean.Cluster, error) {
	cluster := &bean.Cluster{}
	err := util.G_db.Where("app_name = ? and cluster_name = ?", appName, clusterName).Find(cluster).Error
	if err != nil {
		return nil, err
	}
	return cluster, nil
}

func (d *clusterDao) IsClusterExist(appName, clusterName string) bool {
	return !util.G_db.Table("cluster").Where("app_name = ? and cluster_name = ?", appName, clusterName).First(&bean.Cluster{}).RecordNotFound()
}
