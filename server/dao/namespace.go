package dao

import (
	"fmt"

	"disconf/server/model/bean"
	"disconf/server/util"
)

var G_namespace = newNameSpaceDao()

type namespaceDao struct {
}

func newNameSpaceDao() *namespaceDao {
	return &namespaceDao{}
}

func (d *namespaceDao) CreateNamespae(namespace *bean.Namespace) error {
	if !G_cluster.IsClusterExist(namespace.AppName, namespace.ClusterName) {
		return fmt.Errorf("cluster is not exist, %s", namespace.ClusterName)
	}
	if d.IsNameSpaceExit(namespace.AppName, namespace.ClusterName, namespace.NamespaceName) {
		return fmt.Errorf("namespace has exist, %s", namespace.NamespaceName)
	}
	err := util.G_db.Create(namespace).Error
	if err != nil {
		return err
	}
	return nil
}

func (d *namespaceDao) GetNamespace(appName, clusterName, namespaceName string) (*bean.Namespace, error) {
	namespace := &bean.Namespace{}
	db := util.G_db.Where("app_name = ?", appName)
	db = db.Where("cluster_name = ?", clusterName)
	db = db.Where("namespace_name = ?", namespaceName)
	err := db.First(namespace).Error
	if err != nil {
		return nil, err
	}
	return namespace, nil
}

func (d *namespaceDao) IsNameSpaceExit(appName, clusterName, namespaceName string) bool {
	return !util.G_db.Table("namespace").Where("app_name = ? and cluster_name = ? and namespace_name = ?",
		appName, clusterName, namespaceName).First(&bean.Namespace{}).RecordNotFound()
}
