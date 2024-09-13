package handlers

import (
	log "github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
	"runtime"
)

func InitLogging() {
	log.SetOutput(&lumberjack.Logger{
		Filename:   "app.log", // Tên file log
		MaxSize:    10,        // Kích thước tối đa (MB) trước khi xoay vòng log
		MaxBackups: 3,         // Số file log cũ giữ lại
		MaxAge:     28,        // Số ngày giữ lại file log cũ
		Compress:   true,      // Nén các file log cũ
	})
}

// LogDebug Ghi lại thông tin cần thiết để giúp debug, thường là các bước thực thi của chương trình.
func LogDebug(errMsg string) {
	_, file, line, _ := runtime.Caller(1)
	log.WithFields(log.Fields{
		"file": file,
		"line": line,
	}).Debug(errMsg)
}

// LogInfo Ghi lại các thông tin bình thường về tiến trình của chương trình, không phải lỗi nhưng có giá trị theo dõi.
func LogInfo(errMsg string) {
	_, file, line, _ := runtime.Caller(1)
	log.WithFields(log.Fields{
		"file": file,
		"line": line,
	}).Info(errMsg)
}

// LogErr Ghi lại các lỗi, các vấn đề mà chương trình gặp phải nhưng vẫn có thể tiếp tục chạy.
func LogErr(errMsg string) {
	_, file, line, _ := runtime.Caller(1)
	log.WithFields(log.Fields{
		"file": file,
		"line": line,
	}).Error(errMsg)
}

// LogWarn Ghi lại các cảnh báo, thể hiện rằng có điều gì đó không đúng nhưng chương trình vẫn có thể tiếp tục hoạt động.
func LogWarn(errMsg string) {
	_, file, line, _ := runtime.Caller(1)
	log.WithFields(log.Fields{
		"file": file,
		"line": line,
	}).Warn(errMsg)
}

// LogFatal Ghi lại các lỗi nghiêm trọng khiến chương trình phải dừng lại.
func LogFatal(errMsg string) {
	_, file, line, _ := runtime.Caller(1)
	log.WithFields(log.Fields{
		"file": file,
		"line": line,
	}).Fatal(errMsg)
}
