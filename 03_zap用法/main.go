package main

import (
	"time"

	"go.uber.org/zap"
)

func main() {
	Init("my-service", false)
	defer Sync()

	Log.Info("用户登录成功")
	Log.Debug("处理请求",
		zap.String("username", "张三"),
		zap.Int("age", 18),
		zap.Duration("duration", 150*time.Millisecond),
		zap.String("ip", "192.168.1.1"),
		zap.String("method", "GET"),
	)

}
