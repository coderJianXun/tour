package setting

import "github.com/spf13/viper"

type Setting struct {
	vp *viper.Viper
}

// 初始化项目配置，设定配置文件名称，配置文件类型以及路径
func NewSetting() (*Setting, error) {
	vp := viper.New()
	vp.SetConfigName("config")
	vp.AddConfigPath("configs/")
	vp.SetConfigType("yaml")
	err := vp.ReadInConfig()
	if err != nil {
		return nil, err
	}

	return &Setting{vp}, nil
}
