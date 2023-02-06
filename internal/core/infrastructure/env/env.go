//go:build !local && !feature_test

// 開発環境、ステージング環境、本番環境向けの環境変数の取得を行う関数です
// Secrets Managerから値を取得します
// @see https://docs.aws.amazon.com/ja_jp/secretsmanager/latest/userguide/retrieving-secrets_cache-go.html

package env

import (
	"encoding/json"

	"github.com/aws/aws-secretsmanager-caching-go/secretcache"
)

var secretCache, _ = secretcache.New()

// GetString は文字列の値を返却します
func GetString(key string) string {
	val, _ := secretCache.GetSecretString(key)
	return val
}

// GetTokenKey はトークンのキー情報を返却します
func GetTokenKey() (*TokenKey, error) {
	// TODO: キーは仮なので正式に決まったら変更
	val, err := secretCache.GetSecretString("SHIELD_TOKEN_KEYS")
	if err != nil {
		return nil, err
	}

	var res *TokenKey
	if err = json.Unmarshal([]byte(val), res); err != nil {
		return nil, err
	}
	return res, nil
}

// GetReadDbConnectSetting は読み込み用のDB接続情報を返却します
func GetReadDbConnectSetting() (*DbConnectSetting, error) {
	// TODO: キーは仮なので正式に決まったら変更
	val, err := secretCache.GetSecretString("DB_READ_CONNECT_SETTING")
	if err != nil {
		return nil, err
	}

	var res *DbConnectSetting
	if err = json.Unmarshal([]byte(val), res); err != nil {
		return nil, err
	}
	return res, nil
}

// GetWriteDbConnectSetting は読み込み用のDB接続情報を返却します
func GetWriteDbConnectSetting() (*DbConnectSetting, error) {
	// TODO: キーは仮なので正式に決まったら変更
	val, err := secretCache.GetSecretString("DB_WRITE_CONNECT_SETTING")
	if err != nil {
		return nil, err
	}

	var res *DbConnectSetting
	if err = json.Unmarshal([]byte(val), res); err != nil {
		return nil, err
	}
	return res, nil
}
