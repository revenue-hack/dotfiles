//go:build !feature_test

package logger

import (
	"context"

	"gitlab.kaonavi.jp/ae/kgm/log"
	"gitlab.kaonavi.jp/ae/sardine/internal/ctxt"
	"gitlab.kaonavi.jp/ae/sardine/internal/flags"
)

func Info(ctx context.Context, msg Message, params ...parameter) {
	l.Info(msg, fields(ctx, params...)...)
}

func Error(ctx context.Context, err error, params ...parameter) {
	l.Error(err, fields(ctx, params...)...)
}

func Panic(ctx context.Context, err interface{}, params ...parameter) {
	l.Panic(err, fields(ctx, params...)...)
}

// ログにカスタムパラメータを追加する
func fields(ctx context.Context, params ...parameter) []log.Field {
	// 固定で追加するフィールド数分capを追加で確保する
	fields := make([]log.Field, 0, len(params)+3)
	for _, param := range params {
		fields = append(fields, log.NewField(param.Key, param.Value))
	}

	fields = append(fields, log.NewField("request_id", getRequestId(ctx)))
	fields = append(fields, log.NewField("environment", flags.EnvironmentName()))
	fields = append(fields, log.NewField("version", flags.BuildVersion()))
	return fields
}

func getRequestId(ctx context.Context) string {
	// エラーがあるとログ出力できないので握りつぶす
	id, _ := ctxt.RequestId(ctx)
	return id
}
