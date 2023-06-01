package tools

import (
	"github.com/zouxingyuks/tools/config"
	"github.com/zouxingyuks/tools/dao"
)

// InitTools 由于各个功能模块之间存在功能依赖关系,因此需要手动进行初始化
func InitTools() {
	config.LoadDefaultConfig()
	config.InitLog()
	config.ParseConfig()
	dao.InitDao()
}
