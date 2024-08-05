package cmd

import (
	"fmt"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func LogFormat() {

	cfg := zap.Config{
		Level:       zap.NewAtomicLevelAt(zapcore.InfoLevel), // INFO 레벨 설정
		OutputPaths: []string{"stdout"},                      // 표준 출력으로 출력
		Encoding:    "console",                               // json                         // 표준 출력으로 출력
		EncoderConfig: zapcore.EncoderConfig{
			MessageKey:  "message",
			LevelKey:    "level",
			TimeKey:     "time",
			EncodeLevel: zapcore.LowercaseLevelEncoder, // 로그 레벨을 소문자로 표시
			EncodeTime:  zapcore.ISO8601TimeEncoder,    // ISO 8601 형식의 시간으로 표시
		},
	}

	// 로그 생성
	logger, err := cfg.Build()
	if err != nil {
		fmt.Println("Failed to create logger:", err)
		return
	}
	defer logger.Sync()

	filename := "test"
	logger.Info("Hello Zap!", zap.String("filename", filename))
	logger.Warn("Beware of getting Zapped! (Pun)")
	logger.Error("I'm out of Zap joke!")

}
