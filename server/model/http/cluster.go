package http

type CreateClusterRequest struct {
	AppName     string `json:"app_name"`
	ClusterName string `json:"cluster_name"`
	Description string `json:"description"`
}
