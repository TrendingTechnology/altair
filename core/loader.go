package core

import (
	"database/sql"
	"time"
)

type AppLoader interface {
	Compile(configPath string) (AppConfig, error)
}

type DatabaseLoader interface {
	Compile(configPath string) (map[string]DatabaseConfig, error)
}

type PluginLoader interface {
	Compile(pluginPath string) (PluginBearer, error)
}

type PluginBearer interface {
	ConfigExists(pluginName string) bool
	CompilePlugin(pluginName string, injectedStruct interface{}) error
	ForEach(callbackFunc func(pluginName string) error)
	Length() int
}

type DatabaseConfig interface {
	Driver() string
	DBMigrationSource() string
	DBHost() string
	DBPort() (int, error)
	DBUsername() string
	DBPassword() string
	DBDatabase() string
	DBConnectionMaxLifetime() (time.Duration, error)
	DBMaxIddleConn() (int, error)
	DBMaxOpenConn() (int, error)
	Dump() string
}

type DatabaseBearer interface {
	Database(dbName string) (*sql.DB, DatabaseConfig, error)
}

type AppConfig interface {
	Port() int
	BasicAuthUsername() string
	BasicAuthPassword() string
	ProxyHost() string
	PluginExists(pluginName string) bool
	Plugins() []string
	Dump() string
}
