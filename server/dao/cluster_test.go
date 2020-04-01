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

func TestClusterDao_CreateCluster(t *testing.T) {
	cluster := &bean.Cluster{
		BaseModel: bean.BaseModel{
			CreateAt: time.Now(),
			UpdateAt: time.Now(),
			DeleteAt: time.Now(),
		},
		AppName:     "test_app",
		ClusterName: "test_cluster",
		Description: "test",
	}
	if err := G_cluster.CreateCluster(cluster); err != nil {
		panic(err)
	}
}

func TestClusterDao_ClusterExist(t *testing.T) {
	clusterName := "test_cluster"
	appName := "test_app"
	cluster, err := G_cluster.GetCluster(appName, clusterName)
	if err != nil {
		panic(err)
	}
	fmt.Println(cluster)
}

func TestClusterDao_IsClusterExist(t *testing.T) {
	appName := "test_app"
	clusterName := "test_cluster"
	fmt.Println(G_cluster.IsClusterExist(appName, clusterName))
}