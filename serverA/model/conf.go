package model

// 定义配置struct, 属性与配置文件相同
type DBConfig struct {
	Name    string `value:"${name}"`
	Url     string `value:"${url}"`
	Type    string `value:"${type}"`
	ShowSql bool   `value:"${showSql}"`
}

type AppConfig struct {
	AppName     string `value:"${cfg.appName}"`
	ReleaseMode bool   `value:"${cfg.releaseMode}"`
	Port        string `value:"${cfg.port}"`
}
