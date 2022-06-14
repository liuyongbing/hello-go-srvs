package initialize

import "go.uber.org/zap"

/*
InitLogger
初始化 Logger
*/
func InitLogger() {
	logger, _ := zap.NewProduction()
	zap.ReplaceGlobals(logger)
}
