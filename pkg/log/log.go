package log

import (
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/micjn89757/TeaBlog/internal/conf"
	"github.com/micjn89757/TeaBlog/pkg/util"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

func NewLogger(conf conf.Config) *zap.Logger {

	// 定义两种级别
	highPriority := zap.LevelEnablerFunc(func(l zapcore.Level) bool {
		return l >= zapcore.ErrorLevel
	})


	lowPriority := zap.LevelEnablerFunc(func(l zapcore.Level) bool {
		return l < zapcore.ErrorLevel
	})

	// 根据配置文件进行初始化，不同级别信息输出到不同的文件中
	var core zapcore.Core
	switch conf.Log.Env {
	case "development":
		core = zapcore.NewTee(
			zapcore.NewCore(setDevEncoder(), setLogWriter("test.err.log"), highPriority),
			zapcore.NewCore(setDevEncoder(), setLogWriter("test.log"), lowPriority),
		)
	case "production":
		core = zapcore.NewTee(
			zapcore.NewCore(setProductEncoder(), setLogWriter("test.err.log"), highPriority),
			zapcore.NewCore(setProductEncoder(), setLogWriter("test.log"), lowPriority),
		)
	default:
		panic(fmt.Errorf("invalid log env: %s", conf.Log.Env))
	}

	logger := zap.New(core, zap.AddCaller())
	defer logger.Sync()

	return logger
}


// 日志编码器(development)
func setDevEncoder() zapcore.Encoder {
	encoderConfig := zap.NewDevelopmentEncoderConfig()
	encoderConfig.CallerKey = "caller"
	encoderConfig.NameKey = "logger"
	encoderConfig.LevelKey = "level"
	encoderConfig.MessageKey = "msg"
	encoderConfig.TimeKey = "ts"
	encoderConfig.StacktraceKey = "stacktrace"
	encoderConfig.EncodeTime = zapcore.RFC3339TimeEncoder	// 修改时间编码器
	return zapcore.NewJSONEncoder(encoderConfig)
}

// 日志编码器(production)
func setProductEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder	// 修改时间编码器
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder	// 日志文件中使用大写字母记录日志级别
	return zapcore.NewJSONEncoder(encoderConfig)
}

// TODO: 后续需要转发到Kafka，存到相应的数据库用于分析
// 日志输出, 根据文件大小进行切割
func setLogWriter(filename string) zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename: filepath.Join(util.GetRootPath(), "log", filename), 	// 日志文件位置
		MaxSize: 10,	// 在进行切割之前，日志文件的最大大小(MB)
		MaxAge: 30,		// 保留旧文件的最大天数
		MaxBackups: 5,	// 保留旧文件的最大个数
		Compress: false,	// 是否压缩/归档旧文件
	}

	multiwriter := io.MultiWriter(lumberJackLogger, os.Stdout)
	return zapcore.AddSync(multiwriter)
}


