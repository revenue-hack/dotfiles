package logger

import (
	"io"

	"gitlab.kaonavi.jp/ae/kgm/log"
)

type (
	message   = log.Message
	parameter = log.Field
)

var (
	l *log.Logger
)

// New はロガーを初期化します
func New(w io.Writer) {
	l = log.New(w, log.Option{})
}

// Message はログ出力用のメッセージを生成します
func Message(format string, args ...interface{}) message {
	return message{
		Format: format,
		Args:   args,
	}
}

// MakeParameter はログ出力用のパラメータを生成します
func MakeParameter(key string, value interface{}) parameter {
	return parameter{
		Key:   key,
		Value: value,
	}
}
