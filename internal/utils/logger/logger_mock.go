//go:build feature_test

package logger

import (
	"context"
	"fmt"
)

func Info(context.Context, Message, ...parameter) {
	// Infoログは捨てる
}

func Error(_ context.Context, err error, _ ...parameter) {
	fmt.Printf("\n    エラーが発生しました: %#v\n\n", err)
}

func Panic(_ context.Context, err interface{}, _ ...parameter) {
	panic(fmt.Sprintf("    Panicが発生しました: %#v\n", err))
}
