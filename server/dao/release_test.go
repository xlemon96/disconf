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

func TestReleaseDao_CreateRelease(t *testing.T) {
	release := &bean.Release{
		BaseModel: bean.BaseModel{
			CreateAt: time.Now(),
			UpdateAt: time.Now(),
			DeleteAt: time.Now(),
		},
		AppName:       "test_app",
		ClusterName:   "test_cluster",
		NamespaceName: "test_namespace_v2",
		Format:        "ini",
		Value:         "{test:data_v2}",
		Version:       "v4",
		Description:   "",
	}
	err := G_release.CreateRelease(release)
	if err != nil {
		panic(err)
	}
}

func TestReleaseDao_QueryRelease(t *testing.T) {
	appName := "test_app"
	clusterName := "test_cluster"
	namespaceName := "test_namespace"
	version := "v1"
	release, err := G_release.QueryRelease(appName, clusterName, namespaceName, version)
	if err != nil {
		panic(err)
	}
	fmt.Println(release)
}

func TestReleaseDao_GetLatestRelease(t *testing.T) {
	appName := "test_app"
	clusterName := "test_cluster"
	namespaceName := "test_namespace"
	release, err := G_release.GetLatestRelease(appName, clusterName, namespaceName)
	if err != nil {
		panic(err)
	}
	fmt.Println(release)
}