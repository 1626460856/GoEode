package initialize

import (
	"dianshang/app/api/global"
	"github.com/spf13/viper"
)

func SetupViper() {
	viper.SetConfigType("yaml")                     //指定了配置文件的类型为 YAML 格式。这告诉 viper 库在读取配置文件时应该将其解析为 YAML 格式的数据。
	viper.SetConfigName("config")                   //这行代码指定了配置文件的名称为 "config"。这表明您的应用程序预期加载名为 "config.yaml" 的配置文件
	viper.SetConfigFile("app/manifest/config.yaml") //这行代码指定了配置文件的路径为 "app/manifest/config.yaml"。这告诉 viper 库去加载位于指定路径的配置文件
	//一旦 viper 加载了配置文件，您就可以使用它来访问配置信息并填充到相应的结构体中，以便在应用程序中使用
	err := viper.ReadInConfig()
	if err != nil {
		panic("viper read config failed," + err.Error())

	}
	err = viper.Unmarshal(&global.Config)
	if err != nil {
		panic("viper Unmarshal config failed," + err.Error())
	}

}
