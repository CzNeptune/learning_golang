package main

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

/*
生产级配置
*/

var Log *zap.Logger

// Init 初始化日志组件
func Init(serviceName string, debug bool) {
	// 日志切割配置
	lumberJack := &lumberjack.Logger{
		Filename:   "logs/" + serviceName + ".log",
		MaxSize:    100,  // MB
		MaxBackups: 5,    // 保留5个旧文件
		MaxAge:     30,   // 天
		Compress:   true, // 压缩旧文件
	}

	// 编码器配置
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoder := zapcore.NewJSONEncoder(encoderConfig)

	// 多输出：文件 + 控制台
	fileWriter := zapcore.AddSync(lumberJack)
	consoleWriter := zapcore.AddSync(os.Stdout)

	// 日志级别
	level := zapcore.InfoLevel
	if debug {
		level = zapcore.DebugLevel
	}

	// 核心配置
	core := zapcore.NewTee(
		zapcore.NewCore(encoder, fileWriter, level),
		zapcore.NewCore(encoder, consoleWriter, zapcore.DebugLevel),
	)

	// 添加全局字段
	Log = zap.New(core,
		zap.AddCaller(),
		zap.Fields(
			zap.String("service", serviceName),
			zap.String("version", "1.0.0"),
		),
	)
}

// Sync 确保日志刷盘
func Sync() {
	_ = Log.Sync()
}
