//go:build local || feature_test

// ローカル・機能テスト向けの環境変数の取得を行う関数です
// ローカルにSecrets Managerは無いため、実行環境に設定されている環境変数から値を取得します

package env

import (
	"os"
)

func GetString(key string) string {
	return os.Getenv(key)
}

func GetTokenKey() (*TokenKey, error) {
	return &TokenKey{
		EncryptKey: os.Getenv("TOKEN_ENCRYPT_KEY"),
		SigningKey: os.Getenv("TOKEN_SIGNING_KEY"),
	}, nil
}

// GetReadDbConnectSetting は読み込み用のDB接続情報を返却します
func GetReadDbConnectSetting() (*DbConnectSetting, error) {
	return &DbConnectSetting{
		Host:     os.Getenv("DB_READ_HOST_NAME"),
		Username: os.Getenv("DB_USER_NAME"),
		Password: os.Getenv("DB_PASSWORD"),
		Port:     os.Getenv("DB_PORT"),
	}, nil
}

// GetWriteDbConnectSetting は読み込み用のDB接続情報を返却します
func GetWriteDbConnectSetting() (*DbConnectSetting, error) {
	return &DbConnectSetting{
		Host:     os.Getenv("DB_WRITE_HOST_NAME"),
		Username: os.Getenv("DB_USER_NAME"),
		Password: os.Getenv("DB_PASSWORD"),
		Port:     os.Getenv("DB_PORT"),
	}, nil
}
