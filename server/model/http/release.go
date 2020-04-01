package http

type CreateReleaseRequest struct {
	AppName       string `json:"app_name"`
	ClusterName   string `json:"cluster_name"`
	NamespaceName string `json:"namespace_name"`
	Version       string `json:"version"`
	Format        string `json:"format"`
	Value         string `json:"value"`
	Description   string `json:"description"`
}
