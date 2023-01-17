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
	fmt.Println(err)
}

func Panic(_ context.Context, err interface{}, _ ...parameter) {
	panic(err)
}
