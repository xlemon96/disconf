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

func TestNamespaceDao_CreateNamespae(t *testing.T) {
	namespace := &bean.Namespace{
		BaseModel: bean.BaseModel{
			CreateAt: time.Now(),
			UpdateAt: time.Now(),
			DeleteAt: time.Now(),
		},
		AppName:       "test_app",
		ClusterName:   "test_cluster",
		NamespaceName: "test_namespace",
		Format:        "json",
		Value:         "{test:data}",
		Released:      0,
		Description:   "",
	}
	if err := G_namespace.CreateNamespae(namespace); err != nil {
		panic(err)
	}
}

func TestNamespaceDao_GetNamespace(t *testing.T) {
	appName := "test_app"
	clusterName := "test_cluster"
	namespaceName := "test_namespace"
	namespace, err := G_namespace.GetNamespace(appName, clusterName, namespaceName)
	if err != nil {
		panic(err)
	}
	fmt.Println(namespace)
}

func TestNamespaceDao_IsNameSpaceExit(t *testing.T) {
	appName := "test_apppp"
	clusterName := "test_cluster"
	namespaceName := "test_namespace"
	fmt.Println(G_namespace.IsNameSpaceExit(appName, clusterName, namespaceName))
}
