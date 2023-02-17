//go:build feature_test

package logger

import (
	"context"
	"fmt"
)

func Info(context.Context, Message, ...Parameter) {
	// Infoログは捨てる
}

func Error(_ context.Context, err error, _ ...Parameter) {
	fmt.Printf("\n    エラーが発生しました: %#v\n\n", err)
}

func Panic(_ context.Context, err interface{}, _ ...Parameter) {
	panic(fmt.Sprintf("    Panicが発生しました: %#v\n", err))
}
