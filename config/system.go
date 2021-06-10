package config

// System 基础配置
type System struct { // 环境值
	Name    string `mapstructure:"name" json:"name" yaml:"name"`          // 项目名称 	// 项目模式
	Port    int    `mapstructure:"port" json:"port" yaml:"port"`          // 项目使用端口
	Version string `mapstructure:"version" json:"version" yaml:"version"` // 项目版本
}
