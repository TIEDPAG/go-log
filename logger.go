// package log 使用github.com/alecthomas/log4go库实现了一个日志工具类
// 初始化时默认输出至控制台中,打印DEBUG级别及所有级别日志
// 可调用To方法设置日志输出方式
package log

import (
	log "github.com/alecthomas/log4go"
)

func init() {
	To("stdout", "DEBUG")
}

var root = make(log.Logger)

// TO 设置日志对象写入方式
func To(target string, levelName string) {
	var writer log.LogWriter = nil

	switch target {
	case "stdout":
		writer = log.NewConsoleLogWriter()
	case "none": //no logging
	default:
		writer = log.NewFileLogWriter(target, true)
	}

	if writer != nil {
		var level = log.DEBUG

		switch levelName {
		case "FINEST":
			level = log.FINEST
		case "FINE":
			level = log.FINE
		case "DEBUG":
			level = log.DEBUG
		case "TRACE":
			level = log.TRACE
		case "INFO":
			level = log.INFO
		case "WARNING":
			level = log.WARNING
		case "ERROR":
			level = log.ERROR
		case "CRITICAL":
			level = log.CRITICAL
		default:
			level = log.DEBUG
		}

		root.AddFilter("log", level, writer)
	}
}

// Logger 可设置日志前缀的日志接口
// 例: 调用AddLogPrefix添加一个"Test"前缀
//     则该实例稍后调用日志输出则会在日志前加上"[Test]"的前缀
//    调用ClearLogPrefixes后则清空之前的日志前缀  该方法不保证线程安全
type Logger interface {
	AddLogPrefix(string)
	ClearLogPrefixes()
	Debug(string, ...interface{})
	Info(string, ...interface{})
	Warn(string, ...interface{}) error
	Error(string, ...interface{}) error
}

// PrefixLogger 可设置前缀日志实际对象
type PrefixLogger struct {
	*log.Logger
	prefix string
}

// NewPrefixLogger 初始化一个包含前缀的日志对象
func NewPrefixLogger(prefixes ...string) Logger {
	l := &PrefixLogger{Logger: &root}

	for _, p := range prefixes {
		l.AddLogPrefix(p)
	}

	return l
}

// AddLogPrefix 添加一个日志前缀
func (pl *PrefixLogger) AddLogPrefix(prefix string) {
	if len(pl.prefix) > 0 {
		pl.prefix += " "
	}

	pl.prefix += "[" + prefix + "]"
}

// ClearLogPrefixes 清空当前日志对象所有前缀
func (pl *PrefixLogger) ClearLogPrefixes() {
	pl.prefix = ""
}

func (pl *PrefixLogger) pfm(arg0 string) interface{} {
	return pl.prefix + " " + arg0
}

// Debug 输出Debug级别日志
func (pl *PrefixLogger) Debug(arg0 string, args ...interface{}) {
	pl.Logger.Debug(pl.pfm(arg0), args...)
}

// Info 输出Info级别日志
func (pl *PrefixLogger) Info(arg0 string, args ...interface{}) {
	pl.Logger.Info(pl.pfm(arg0), args...)
}

// Warn 输出Warn级别日志
func (pl *PrefixLogger) Warn(arg0 string, args ...interface{}) error {
	return pl.Logger.Warn(pl.pfm(arg0), args...)
}

// Error 输出Error级别日志
func (pl *PrefixLogger) Error(arg0 string, args ...interface{}) error {
	return pl.Logger.Error(pl.pfm(arg0), args...)
}

// Debug 输出Debug级别日志
func Debug(arg0 string, args ...interface{}) {
	root.Debug(arg0, args...)
}

// Info 输出Info级别日志
func Info(arg0 string, args ...interface{}) {
	root.Info(arg0, args...)
}

// Warn 输出Warn级别日志
func Warn(arg0 string, args ...interface{}) error {
	return root.Warn(arg0, args...)
}

// Error 输出Error级别日志
func Error(arg0 string, args ...interface{}) error {
	return root.Error(arg0, args...)
}
