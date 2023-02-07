package env

// TokenKey はトークンのキー情報を表す構造体
type TokenKey struct {
	EncryptKey string `json:"encrypt_key"`
	SigningKey string `json:"signing_key"`
}

// DbConnectSetting はDBの接続情報を表す構造体
//
// @see https://docs.aws.amazon.com/ja_jp/secretsmanager/latest/userguide/reference_secret_json_structure.html
//
// ※engine, dbnameは使用しないので意図的に省いています
type DbConnectSetting struct {
	Host     string `json:"host"`
	Username string `json:"username"`
	Password string `json:"password"`
	Port     string `json:"port"`
}
