package bean

type Release struct {
	BaseModel
	AppName       string
	ClusterName   string
	NamespaceName string
	Version       string
	Format        string
	Value         string
	Description   string
}
