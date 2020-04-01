package dao

import (
	"fmt"
	"time"

	"disconf/server/model/bean"
	"disconf/server/util"
)

var G_release = newReleaseDao()

type releaseDao struct {
}

func newReleaseDao() *releaseDao {
	return &releaseDao{}
}

//发布一条消息
func (d *releaseDao) CreateRelease(release *bean.Release) error {
	_, err := G_namespace.GetNamespace(release.AppName, release.ClusterName, release.NamespaceName)
	if err != nil {
		return err
	}
	_, err = d.QueryRelease(release.AppName, release.ClusterName, release.NamespaceName, release.Version)
	if err == nil {
		return  fmt.Errorf("release version is conflict, %s", release.Version)
	}
	util.G_db.Begin()
	//1、记录发布信息
	if err := util.G_db.Create(release).Error; err != nil {
		util.G_db.Rollback()
		return err
	}
	//2、更改配置信息
	if err := util.G_db.Table("namespace").Where("app_name = ? and cluster_name = ? and namespace_name = ?",
		release.AppName, release.ClusterName, release.NamespaceName).Updates(map[string]interface{}{
		"released":  true,
		"value":     release.Value,
		"format":    release.Format,
		"update_at": time.Now(),
	}).Error; err != nil {
		util.G_db.Rollback()
		return err
	}
	util.G_db.Commit()
	return nil
}

func (d *releaseDao) QueryRelease(appName, clusterName, namespaceName, version string) (*bean.Release, error) {
	release := &bean.Release{}
	db := util.G_db.Where("app_name = ?", appName)
	db = db.Where("cluster_name = ?", clusterName)
	db = db.Where("namespace_name = ?", namespaceName)
	db = db.Where("version = ?", version)
	if err := db.First(release).Error; err != nil {
		return nil, err
	}
	return release, nil
}

func (d *releaseDao) GetLatestRelease(appName, clusterName, namespaceName string) (*bean.Release, error) {
	release := &bean.Release{}
	db := util.G_db.Where("app_name = ?", appName)
	db = db.Where("cluster_name = ?", clusterName)
	db = db.Where("namespace_name = ?", namespaceName)
	if err := db.Last(release).Error; err != nil {
		return nil, err
	}
	return release, nil
}
