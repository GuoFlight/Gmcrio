package conf

var (
	Version     = "V0.0.0"
	TraceIdName = "TraceId"
)

type Role string

const (
	RoleAdmin  Role = "admin"
	RoleWriter Role = "writer"
	RoleViewer Role = "viewer"
)
