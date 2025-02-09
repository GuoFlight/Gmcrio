package conf

const (
	Version     = "V0.0.0"
	KeyTraceId  = "TraceId"
	KeyUsername = "username"
)

const (
	defaultHttpLogFilePath = "./logs/access.log"
)

type Role string

const (
	RoleAdmin  Role = "admin"
	RoleWriter Role = "writer"
	RoleViewer Role = "viewer"
)
