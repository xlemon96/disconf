package bean

type Namespace struct {
	BaseModel
	AppName       string
	ClusterName   string
	NamespaceName string
	Format        string
	Value         string
	Released      int8
	Description   string
}
