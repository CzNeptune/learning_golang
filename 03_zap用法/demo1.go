package main

import (
	"time"

	"go.uber.org/zap"
)

func demo1() {
	logger, _ := zap.NewProduction()
	defer logger.Sync() // 确保日志被写入

	logger.Info("用户登录成功",
		zap.String("username", "张三"),
		zap.Int("age", 18),
		zap.Duration("duration", 150*time.Millisecond),
	)

}
