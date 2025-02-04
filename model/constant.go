package model

// log-level
var (
	LogLevel      = "log-level"
	LogLevelInfo  = "info"
	LogLevelError = "error"
	LogLevelDebug = "debug"
	LogLevelFatal = "fatal"
	LogLevelWarn  = "warn"
)

// package name

var (
	ApiPackage         = "api"
	StorePackage       = "store"
	ModelPackage       = "model"
	ControllersPackage = "controllers"
	UtilsPackage       = "utils"
	MainPackage        = "main"
)

// function name

var (
	Controllers = "controllers"
	Store       = "store"
	Cmd         = "cmd"
	Utils       = "utils"
	Api         = "api"
)

var (
	NewServer  = "new-server"
	NewStore   = "new-store"
	CreateUser = "create-user"
	GetUsers = "get-users"
	GetUserByID = "get-user-by-id"
	UpdateUser = "update-user"
	DeleteUser = "delete-user"
)
